package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	//"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
)

func main() {
	// 创建SDK实例
	sdk, err := fabsdk.New(config.FromFile("./connection.yaml"))
	if err != nil {
		fmt.Printf("Failed to create SDK: %v\n", err)
		return
	}
	defer sdk.Close()

	// 创建通道客户端
	chClient, err := channel.New(sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1")))
	if err != nil {
		fmt.Printf("Failed to create channel client: %v\n", err)
		return
	}

	// 调用智能合约方法
	response, err := chClient.Execute(channel.Request{
		ChaincodeID: "reputation-contract",
		Fcn:         "increaseReputation",
		Args:        [][]byte{[]byte("user1")},
	})
	if err != nil {
		fmt.Printf("Failed to execute transaction: %v\n", err)
		return
	}

	fmt.Println(string(response.Payload))
}
