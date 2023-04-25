package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Data 数据结构定义
type Data struct {
	DataHash string //json:"dataHash"
}
type UpdateData struct {
	DataHash    string //json:"dataHash"
	NewDataHash string //json:"newDataHash"
}

// Path 文件路径结构定义
type Path struct {
	FIlePath string //json:"filePath"
}

// Share 分享码结构定义
type Share struct {
	OwnerId   string //json:"ownerId"
	ShareCode string //json:"shareCode"
}

type SimpleChaincode struct {
}

// Init 初始化智能合约
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke 执行智能合约的操作
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "storeDataHash" {
		return t.storeDataHash(stub, args)
	} else if function == "queryDataHash" {
		return t.queryDataHash(stub, args)
	} else if function == "updateDataHash" {
		return t.updateDataHash(stub, args)
	} else if function == "queryFilePath" {
		return t.queryFilePath(stub, args)
	} else if function == "storeFIlePath" {
		return t.storeFilePath(stub, args)
	} else if function == "queryShareCode" {
		return t.queryShareCode(stub, args)
	} else if function == "storeShareCode" {
		return t.storeShareCode(stub, args)
	}
	return shim.Error("Invalid function name.")
}

func (t *SimpleChaincode) storeFilePath(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	id := args[0]
	filePath := args[1]
	data := Path{FIlePath: filePath}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 将数据哈希值写入链上
	err = stub.PutState("P"+id, dataBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent(args[2], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) queryFilePath(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	filePath, err := stub.GetState("P" + args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(filePath)
}

func (t *SimpleChaincode) queryDataHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// 查询数据哈希值
	dataBytes, err := stub.GetState("H" + args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	//if dataBytes == nil {
	//   return shim.Error("No data with key: " + args[0])
	//}

	return shim.Success(dataBytes)
}

func (t *SimpleChaincode) updateDataHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	updateData := UpdateData{}
	err := json.Unmarshal([]byte(args[0]), &updateData)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 查询数据哈希值
	dataBytes, err := stub.GetState(updateData.DataHash)
	if err != nil {
		return shim.Error(err.Error())
	}

	if dataBytes == nil {
		return shim.Error("No data with key: " + updateData.DataHash)
	}

	data := Data{}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 更新数据哈希值
	data.DataHash = updateData.NewDataHash
	dataBytes, err = json.Marshal(data)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(data.DataHash, dataBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// 将数据哈希值存储到链上
func (t *SimpleChaincode) storeDataHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	id := args[0]
	dataHash := args[1]
	data := Data{DataHash: dataHash}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 将数据哈希值写入链上
	err = stub.PutState("H"+id, dataBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent(args[2], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *SimpleChaincode) queryShareCode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	shareCode, err := stub.GetState("S" + args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(shareCode)
}

func (t *SimpleChaincode) storeShareCode(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	id := args[0]
	fileName := args[1]
	ownerId := args[2]

	timestamp, err := stub.GetTxTimestamp()
	timestampBytes, err := json.Marshal(timestamp)
	filenameBytes := []byte(fileName)
	randomNumber, err := generateRandomNumber(16)

	data := append(timestampBytes, filenameBytes...)
	data = append(data, randomNumber...)

	hashBytes := sha256.Sum256(data)
	hashString := hex.EncodeToString(hashBytes[:])
	shareCode := hashString[:8]

	shareData := Share{ShareCode: shareCode, OwnerId: ownerId}

	shareDataBytes, err := json.Marshal(shareData)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 将数据哈希值写入链上
	err = stub.PutState("S"+id, shareDataBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent(args[3], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// 生成指定长度的随机数
func generateRandomNumber(length int) (string, error) {
	var randomNumber []byte
	randomNumber, err := crypto.GetRandomBytes(length)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomNumber), nil
}

// 启动智能合约
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
