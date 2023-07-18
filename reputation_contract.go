package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type ReputationContract struct {
}

func (rc *ReputationContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (rc *ReputationContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取调用的函数和参数
	function, args := stub.GetFunctionAndParameters()

	if function == "increaseReputation" {
		// 增加信誉度
		return rc.increaseReputation(stub, args)
	} else if function == "decreaseReputation" {
		// 减少信誉度
		return rc.decreaseReputation(stub, args)
	} else if function == "getReputation" {
		// 获取信誉度
		return rc.getReputation(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (rc *ReputationContract) increaseReputation(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// 检查参数个数
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	// 获取用户或服务提供者的唯一标识符
	entityID := args[0]

	// 根据标识符从区块链中获取当前信誉度
	reputationBytes, _ := stub.GetState(entityID)
	var reputation int
	if reputationBytes == nil {
		reputation = 0
	} else {
		reputation = int(reputationBytes[0])
	}

	// 增加信誉度并保存到区块链中
	reputation++
	stub.PutState(entityID, []byte{byte(reputation)})

	return shim.Success([]byte(fmt.Sprintf("Reputation increased to %d.", reputation)))
}

func (rc *ReputationContract) decreaseReputation(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// 检查参数个数
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	// 获取用户或服务提供者的唯一标识符
	entityID := args[0]

	// 根据标识符从区块链中获取当前信誉度
	reputationBytes, _ := stub.GetState(entityID)
	var reputation int
	if reputationBytes == nil {
		reputation = 0
	} else {
		reputation = int(reputationBytes[0])
	}

	// 减少信誉度并保存到区块链中
	reputation--
	stub.PutState(entityID, []byte{byte(reputation)})

	return shim.Success([]byte(fmt.Sprintf("Reputation decreased to %d.", reputation)))
}

func (rc *ReputationContract) getReputation(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// 检查参数个数
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	// 获取用户或服务提供者的唯一标识符
	entityID := args[0]

	// 根据标识符从区块链中获取当前信誉度
	reputationBytes, _ := stub.GetState(entityID)
	var reputation int
	if reputationBytes == nil {
		reputation = 0
	} else {
		reputation = int(reputationBytes[0])
	}

	return shim.Success([]byte(fmt.Sprintf("Reputation: %d.", reputation)))
}

func main() {
	err := shim.Start(new(ReputationContract))
	if err != nil {
		fmt.Printf("Error starting Reputation smart contract: %s", err)
	}
}
