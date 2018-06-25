//合约的部署工具
//author:caige
package deploy

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"comment/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	host := "http://localhost:8545"
	conn := connectGethRpc(host)
	keystoreFile := "/Users/caige/geth/init/data/keystore/UTC--2018-06-23T05-02-34.518642288Z--256916bba6f3c2b6519226d083b7b2fbc9eef646"
	auth := createGethAuth(keystoreFile, "caige")
	address, ts, _, err := deployContract(conn, auth)
	if err != nil {
		logger("部署失败:" + err.Error())
		os.Exit(1)
	}
	logger("部署完成!")
	logger("合约地址:" + address.Hex())
	logger("交易Hash:" + ts.Hash().Hex())
}

//连接geth
func connectGethRpc(rpcHost string) *ethclient.Client {
	logger("连接geth host:" + rpcHost)
	client, dialErr := rpc.Dial(rpcHost)
	if dialErr != nil {
		panic(dialErr)
		logger("连接失败:" + dialErr.Error())
		os.Exit(1)
	}
	conn := ethclient.NewClient(client)
	network, getNetworkIdErr := conn.NetworkID(context.TODO())
	if getNetworkIdErr != nil {
		logger("获取network id 失败:" + getNetworkIdErr.Error())
		os.Exit(1)
	} else {
		logger("获取network id:" + network.String())
	}
	return conn
}

//生成geth认证账户
func createGethAuth(keystoreFile, pass string) *bind.TransactOpts {
	keyStoreContent, readErr := ioutil.ReadFile(keystoreFile)
	if readErr != nil {
		logger("读取keystore 文件失败:" + readErr.Error())
		os.Exit(1)
	}
	key := string(keyStoreContent)
	auth, err := bind.NewTransactor(strings.NewReader(key), pass)
	if err != nil {
		logger("创建transactor失败:" + err.Error())
		os.Exit(1)
	}
	return auth
}

//部署合约
func deployContract(conn *ethclient.Client, auth *bind.TransactOpts) (common.Address, *types.Transaction, *token.Comment, error) {
	return token.DeployComment(auth, conn)
}

//日志
func logger(v interface{}) {
	fmt.Println(v)
}
