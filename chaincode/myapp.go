package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"time"
)

// Data 数据结构定义
type Data struct {
	DataHash   string //json:"dataHash"
	MerkleRoot string
}

type MerkleNode struct {
	LeftChild  *MerkleNode
	RightChild *MerkleNode
	Data       string
	Hash       string
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

type Log struct {
	Timestamp time.Time // 日志时间戳
	UserID    string    // 用户ID
	Action    string    // 操作类型
	Details   string    // 操作详情
	// 其他日志属性
}

type LedgerData struct {
	Key      string      `json:"key"`
	DataType string      `json:"dataType"`
	Data     interface{} `json:"data"`
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
	} else if function == "queryLedger" {
		return t.queryLedger(stub)
	} else if function == "logAction" {
		return t.logAction(stub, args)
	} else if function == "queryLogs" {
		return t.queryLogs(stub)
	} else if function == "queryUserLogs" {
		return t.queryUserLogs(stub, args)
	} else if function == "verifyFile" {
		return t.verifyFile(stub, args)
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

func computeHash(data string) string {
	hashBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashBytes[:])
}

func computeMerkleRoot(data []string) string {
	// 构建 Merkle 树
	leaves := make([]*MerkleNode, len(data))
	for i, c := range data {
		leaves[i] = &MerkleNode{
			Data: c,
			Hash: computeHash(c),
		}
	}

	for len(leaves) > 1 {
		var nextLevel []*MerkleNode
		for i := 0; i < len(leaves); i += 2 {
			node := &MerkleNode{
				LeftChild:  leaves[i],
				RightChild: leaves[i+1],
				Hash:       computeHash(leaves[i].Hash + leaves[i+1].Hash),
			}
			nextLevel = append(nextLevel, node)
		}
		leaves = nextLevel
	}

	return leaves[0].Hash
}

// 将数据哈希值存储到链上
func (t *SimpleChaincode) storeDataHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	id := args[0]
	dataHash := args[1]
	fileContent := []string{args[2]}
	// 计算文件的Merkle树根哈希
	rootHash := computeMerkleRoot(fileContent)
	data := Data{
		DataHash:   dataHash,
		MerkleRoot: rootHash,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return shim.Error(err.Error())
	}

	// 将数据哈希值写入链上
	err = stub.PutState("H"+id, dataBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent(args[3], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *SimpleChaincode) verifyFile(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证文件完整性
	// args[0]: 文件ID

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	fileID := args[0]
	fileContent := []string{args[1]}
	fileKey := "H" + fileID

	// 查询文件数据
	fileBytes, _ := stub.GetState(fileKey)
	if fileBytes == nil {
		return shim.Error("File not found.")
	}
	var fileData Data
	// 查询文件的Merkle树根哈希
	err := json.Unmarshal(fileBytes, &fileData)
	if err != nil {
		return shim.Error("Root hash not found.")
	}

	rootHash := fileData.MerkleRoot

	// 计算文件数据的Merkle树根哈希
	computedRootHash := computeMerkleRoot(fileContent)

	// 验证Merkle树根哈希是否匹配
	if computedRootHash != rootHash {
		return shim.Error("File integrity check failed.")
	}

	return shim.Success([]byte("File integrity check passed."))
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

	return shim.Success(shareDataBytes)
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

func (t *SimpleChaincode) queryLedger(stub shim.ChaincodeStubInterface) pb.Response {
	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to query ledger: %s", err.Error()))
	}
	defer resultsIterator.Close()

	var ledgerData []LedgerData
	for resultsIterator.HasNext() {
		result, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to retrieve ledger data: %s", err.Error()))
		}

		// 根据实际情况解析不同类型的数据
		// 这里假设数据类型保存在result.Key的前缀中，例如"H"代表数据类型为Data，"P"代表数据类型为Path
		dataType := result.Key[:1]
		switch dataType {
		case "H":
			data := Data{}
			err = json.Unmarshal(result.Value, &data)
			if err != nil {
				return shim.Error(fmt.Sprintf("Failed to unmarshal Data: %s", err.Error()))
			}
			ledgerData = append(ledgerData, LedgerData{Key: result.Key, DataType: "Data", Data: data})
		case "P":
			data := Path{}
			err = json.Unmarshal(result.Value, &data)
			if err != nil {
				return shim.Error(fmt.Sprintf("Failed to unmarshal Path: %s", err.Error()))
			}
			ledgerData = append(ledgerData, LedgerData{Key: result.Key, DataType: "Path", Data: data})
		case "S":
			data := Share{}
			err = json.Unmarshal(result.Value, &data)
			if err != nil {
				return shim.Error(fmt.Sprintf("Failed to unmarshal Path: %s", err.Error()))
			}
			ledgerData = append(ledgerData, LedgerData{Key: result.Key, DataType: "Share", Data: data})
		}
	}

	ledgerBytes, err := json.Marshal(ledgerData)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to marshal ledger data: %s", err.Error()))
	}

	return shim.Success(ledgerBytes)
}

func (t *SimpleChaincode) logAction(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	userID := args[0]
	action := args[1]
	details := args[2]
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error(err.Error())
	}
	txTime := time.Unix(timestamp.Seconds, int64(timestamp.Nanos))
	log := Log{
		Timestamp: txTime,
		UserID:    userID,
		Action:    action,
		Details:   details,
	}

	logBytes, err := json.Marshal(log)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState("LOG-"+strconv.FormatInt(log.Timestamp.UnixNano(), 10), logBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent(args[3], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) queryLogs(stub shim.ChaincodeStubInterface) pb.Response {
	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var logs []Log

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if queryResponse.Key[:1] == "L" {
			var log Log
			err = json.Unmarshal(queryResponse.Value, &log)
			if err != nil {
				return shim.Error(err.Error())
			}

			logs = append(logs, log)
		}

	}

	logsBytes, err := json.Marshal(logs)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(logsBytes)
}

func (t *SimpleChaincode) queryUserLogs(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	userID := args[0]
	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var logs []Log

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if queryResponse.Key[:1] == "L" {
			_, compositeKeyParts, err := stub.SplitCompositeKey(queryResponse.Key)
			if err != nil {
				return shim.Error(err.Error())
			}

			if compositeKeyParts[1] == userID {
				var log Log
				err = json.Unmarshal(queryResponse.Value, &log)
				if err != nil {
					return shim.Error(err.Error())
				}

				logs = append(logs, log)
			}
		}

	}

	logsBytes, err := json.Marshal(logs)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(logsBytes)
}

// 启动智能合约
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
