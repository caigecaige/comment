package handler

import (
	"context"
	"math/big"
	"time"

	"birthday/constant"
	"birthday/form"
	"birthday/token"
	"caige/componnet"

	"github.com/Unknwon/goconfig"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/martini-contrib/render"
)

//获取合约的hash地址
func getContractAddress(config *goconfig.ConfigFile) (common.Address, error) {
	addHex := config.MustValue("geth", "contract")
	return common.HexToAddress(addHex), nil
}

//保存数据
func Set(config *goconfig.ConfigFile, resJson componnet.ZasResponseJson, r render.Render, form form.RequestForm, client *ethclient.Client, opts *bind.TransactOpts) {
	if form.Name == "" || form.Message == "" {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", "name跟message不能为空"))
		return
	}
	contractAdd, getAddErr := getContractAddress(config)
	if getAddErr != nil {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", "contract address not found"))
		return
	}
	contractHandler, err := token.NewComment(contractAdd, client)
	if err != nil {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", err.Error()))
		return
	}
	ts := bind.TransactOpts{From: opts.From, Signer: opts.Signer}
	res, err := contractHandler.Set(&ts, []byte(form.Name), []byte(form.Message), big.NewInt(time.Now().Unix()))
	r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_SUCC, res, "succ"))
	return
}

//获取单个数据
func Get(config *goconfig.ConfigFile, r render.Render, resJson componnet.ZasResponseJson, form form.RequestForm, client *ethclient.Client, opts *bind.TransactOpts) {
	indexId := form.IndexId
	contractAdd, getAddErr := getContractAddress(config)
	if getAddErr != nil {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", "contract address not found"))
		return
	}
	contractHandler, err := token.NewComment(contractAdd, client)
	if err != nil {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", err.Error()))
		return
	}
	cts := bind.CallOpts{From: opts.From, Pending: false}
	res, err := contractHandler.Get(&cts, big.NewInt(indexId))
	if err != nil {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", err.Error()))
		return
	}
	resp := map[string]string{"name": string(res.Name), "message": string(res.Message)}
	r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, resp, "succ"))
	return
}

//部署合约
func Deploy(r render.Render, resJson componnet.ZasResponseJson, form form.RequestForm, client *ethclient.Client, opts *bind.TransactOpts) {
	address, tx, _, err := token.DeployComment(opts, client)
	if err == nil {
		configData := make(map[string]string)
		configData["address"] = address.Hex()
		configData["ts"] = tx.Hash().String()
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_SUCC, configData, "deploy succ!"))
		return
	} else {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", "deploy fail! "+err.Error()))
		return
	}
}

//节点信息
func Node(r render.Render, resJson componnet.ZasResponseJson, form form.RequestForm, client *ethclient.Client, opts *bind.TransactOpts) {
	resp := make(map[string]interface{})
	resp["ip"] = "120.78.144.20"
	networkId, _ := client.NetworkID(context.TODO())
	resp["networkId"] = networkId.String()
	header, _ := client.HeaderByNumber(context.Background(), nil)
	resp["blockNumber"] = header.Number.String()
	resp["pendingTransaction"], _ = client.PendingTransactionCount(context.TODO())
	resp["blockGasLimit"] = header.GasLimit
	resp["headerTime"] = time.Unix(header.Time.Int64(), 0).Format(constant.TimeLayout)
	r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_SUCC, resp, "succ!"))
	return
}

//取出列表数据
func List(config *goconfig.ConfigFile, r render.Render, resJson componnet.ZasResponseJson, form form.RequestForm, client *ethclient.Client, opts *bind.TransactOpts) {
	if form.Limit == 0 {
		form.Limit = 10
	}
	data := make([]map[string]interface{}, 0)
	//处理合约
	contractAdd, getAddErr := getContractAddress(config)
	if getAddErr != nil {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", "contract address not found"))
		return
	}
	contractHandler, err := token.NewComment(contractAdd, client)
	if err != nil {
		r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, "", err.Error()))
		return
	}
	cts := bind.CallOpts{From: opts.From, Pending: false}
	total, _ := contractHandler.Length(&cts)
	start := int(total.Int64() - 1)
	for i := start; i > (start - form.Limit); i-- {
		contractSetData, err := contractHandler.Get(&cts, big.NewInt(int64(i)))
		if err != nil {
			continue
		}
		addTime := time.Unix(contractSetData.Time.Int64(), 0).Format(constant.TimeLayout)
		name := "同步中..."
		if string(contractSetData.Name) != "" {
			name = string(contractSetData.Name)
		}
		message := string(contractSetData.Message)
		item := map[string]interface{}{"id": i, "name": name, "message": message, "add_time": addTime}
		data = append(data, item)
	}
	resp := make(map[string]interface{})
	resp["list"] = data
	block := make(map[string]interface{})
	networkId, _ := client.NetworkID(context.TODO())
	block["network_id"] = networkId.String()
	header, _ := client.HeaderByNumber(context.Background(), nil)
	block["max"] = header.Number.String()
	block["pending"], _ = client.PendingTransactionCount(context.TODO())
	resp["block"] = block
	r.JSON(constant.HTTP_STATUS_200, resJson.Render(constant.JSON_CODE_ERROR, resp, "succ"))
	return
}
