// app
package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"

	"birthday/constant"
	"birthday/job"
	"caige/componnet"
	"caige/database"
	"caige/schedule-service"

	"github.com/Unknwon/goconfig"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-martini/martini"
	"github.com/go-xorm/core"
	"github.com/martini-contrib/render"
)

type appServe struct {
	Name            string
	Engine          *martini.ClassicMartini
	configFile      string
	ConfigHandler   *goconfig.ConfigFile
	LogHandler      *componnet.CLog
	Version         string
	VersionNote     string
	MysqlEngine     *CDatabase.Mysql
	MongoEngine     *CDatabase.Mongo
	ScheduleService *schedule_service.ScheduleService
	RunStartTime    time.Time
	RunEndTime      time.Time
	Error           error
}

func NewApp() *appServe {
	app := new(appServe)
	app.Name = "birthday-dapp"
	return app
}

func (app *appServe) Init(configFile string) {
	app.Engine = martini.Classic()
	martini.Env = constant.Prod
	app.configFile = configFile
	app.Version = "1.0"
	app.VersionNote = "release:2018-1-5"
	app.checRuntime()
}

//检查runtime
func (app *appServe) checRuntime() {
	app.ConfigHandler, app.Error = goconfig.LoadConfigFile(app.configFile)
	if app.Error != nil {
		app.End()
	}
}

//app服务安装组件
func (app *appServe) Setup() {
	app.LogHandler = new(componnet.CLog)
	app.LogHandler.Init(app.ConfigHandler.MustValue("core", "log_path", "logs"))
	app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "系统启动...配置文件["+app.configFile+"]")
	app.Engine.Logger(app.LogHandler.Handlers[componnet.LOG_TYPE_API])
	initRouter(app.Engine)
	app.setupDb()
	app.registerService()
	app.registerMiddleware()
	go app.receiveSignal()
	app.setupCronJob()
	app.ConnectGethRpc()
	app.CreateGethAuth()
}

//连接db
func (app *appServe) setupDb() {
	mysqlSection, secErr := app.ConfigHandler.GetSection("mysql")
	if len(mysqlSection) > 0 && secErr == nil {
		host := mysqlSection["host"]
		port := mysqlSection["port"]
		user := mysqlSection["user"]
		password := mysqlSection["password"]
		database := mysqlSection["database"]
		chartset := mysqlSection["charset"]
		tbPrefix := mysqlSection["table_prefix"]
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "连接Mysql...")
		app.MysqlEngine = CDatabase.NewMysql(host, port, user, password, database, chartset, tbPrefix)
		connectErr := app.MysqlEngine.TryConnect()
		if connectErr != nil {
			app.Error = connectErr
			app.End()
		}
		//注入mysql实例
		CDatabase.RegisterMysqlEngine(app.MysqlEngine.Engine, app.LogHandler.SqlHandler, core.LOG_DEBUG)
	}
	mongoSection, mSecErr := app.ConfigHandler.GetSection("mongo")
	if len(mongoSection) > 0 && mSecErr == nil {
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "连接Mongodb...")
		mongoHost := mongoSection["host"]
		mongoPort := mongoSection["port"]
		if mongoSection["user"] != "" && mongoSection["password"] != "" {
			app.MongoEngine = CDatabase.NewDefaultMongoWithAuth(mongoHost, mongoPort, mongoSection["user"], mongoSection["password"], mongoSection["dbname"])
		} else {
			app.MongoEngine = CDatabase.NewDefaultMongo(mongoHost, mongoPort)
		}
		mongoConnErr := app.MongoEngine.TryConnect()
		if mongoConnErr != nil {
			app.Error = mongoConnErr
			app.End()
		}
		mgLogger := log.New(app.LogHandler.SqlHandler, "[mongo] ", log.Ldate|log.Ltime|log.Llongfile)
		CDatabase.RegisterMongoEngine(app.MongoEngine.Engine.DB(mongoSection["dbname"]), mgLogger, true)
	}
}

//计划任务
func (app *appServe) setupCronJob() {
	app.ScheduleService = schedule_service.NewScheduleService(app.LogHandler.Handlers[componnet.LOG_TYPE_CRON])
	app.ScheduleService.AddMultiJobs(job.ScheduleJobList())
	go app.ScheduleService.Run()
}

//注册服务
func (app *appServe) registerService() {
	app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "注册系统服务...")
	//注册返回的json结构
	var resJson componnet.ZasResponseJson
	resJson.Logger = app.LogHandler.Handlers[componnet.LOG_TYPE_API]
	app.Engine.Map(resJson)
	//注册日志
	app.Engine.Map(app.LogHandler)
	//注入全局配置文件
	app.Engine.Map(app.ConfigHandler)
}

//注册中间件
func (app *appServe) registerMiddleware() {
	app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "注册中间件...")
	app.Engine.Use(render.Renderer(render.Options{Charset: "UTF-8", IndentJSON: false}))
	staticDir := app.ConfigHandler.MustValue("core", "static_directory")
	if staticDir != "" {
		app.Engine.Use(martini.Static(staticDir))
	}
}

//监听系统输入信号
func (app *appServe) receiveSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	signalContent := <-c
	if signalContent == os.Interrupt || signalContent == os.Kill {
		app.RunEndTime = time.Now()
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "System Interrupt")
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "Remove Pid File...")
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "----  Server Info  ----")
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "Version   :"+app.Version)
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "Start Time:"+app.RunStartTime.Format(constant.TimeLayout))
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "End   Time:"+app.RunEndTime.Format(constant.TimeLayout))
		duration := app.RunEndTime.Sub(app.RunStartTime)
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "Run   Time:"+duration.String())
		os.Exit(0)
	}
}

//执行部署的内容
func (app *appServe) Deployment() {
	job.DeploymentFuncs(app.LogHandler.Handlers[componnet.LOG_TYPE_SERVICE])
}

//连接geth
func (app *appServe) ConnectGethRpc() {
	client, dialErr := rpc.Dial("http://localhost:8545")
	if dialErr != nil {
		panic(dialErr)
	}
	app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "连接geth host...")
	conn := ethclient.NewClient(client)
	network, _ := conn.NetworkID(context.TODO())
	app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "geth host:"+network.String())
	app.Engine.Map(conn)
}

//生成geth认证账户
func (app *appServe) CreateGethAuth() {
	keyStore := app.ConfigHandler.MustValue("geth", "keystore")
	if keyStore == "" {
		panic("not found key store file")
	}
	keyStoreContent, readErr := ioutil.ReadFile(keyStore)
	if readErr != nil {
		panic("read key store file error:" + readErr.Error())
	}
	key := string(keyStoreContent)
	pw := app.ConfigHandler.MustValue("geth", "pass")
	auth, err := bind.NewTransactor(strings.NewReader(key), pw)
	if err != nil {
		panic(err)
	}
	app.Engine.Map(auth)
}

//app运行
func (app *appServe) Run() {
	app.RunStartTime = time.Now()
	host := app.ConfigHandler.MustValue("core", "host")
	port := app.ConfigHandler.MustValue("core", "port")
	app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "启动完成...listening["+host+":"+port+"]")
	app.Engine.RunOnAddr(host + ":" + port)
}

//结束app
func (app *appServe) End() {
	if app.LogHandler != nil {
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, "系统发生错误:")
		app.LogHandler.Println(componnet.LOG_TYPE_SERVICE, app.Error)
	} else {
		fmt.Println("错误:" + app.Error.Error())
	}
	os.Exit(1)
}
