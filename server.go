package main

import (
	"flag"

	"comment/core"
)

func main() {
	//全局异常
	//defer func() {
	//	if globalError := recover(); globalError != nil {
	//		fmt.Println(globalError)
	//		os.Exit(1)
	//	}
	//}()
	configFile := parseArgs()
	app := core.NewApp()
	app.Init(configFile)
	app.Setup()
	app.Deployment()
	app.Run()
}

func parseArgs() string {
	var configFile string
	flag.StringVar(&configFile, "c", "", "请输入配置文件路径")
	flag.Parse()
	return configFile
}
