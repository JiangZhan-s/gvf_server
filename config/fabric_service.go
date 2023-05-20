package config

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"strconv"
)

func (t *ServiceSetup) StoreShareCode(fileId, fileName, ownerId string) (string, error) {
	eventID := "eventStoreShareCode"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "storeShareCode", Args: [][]byte{[]byte(fileId), []byte(fileName), []byte(ownerId), []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) QueryShareCode(fileId string) (string, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryShareCode", Args: [][]byte{[]byte(fileId)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) StoreDataHash(name, num string) (string, error) {

	eventID := "eventstoreDataHash"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "storeDataHash", Args: [][]byte{[]byte(name), []byte(num), []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) QueryDataHash(name string) (string, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryDataHash", Args: [][]byte{[]byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) QueryLedger() (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryLedger", Args: [][]byte{}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) LogAction(userID uint, action string, details string) (string, error) {

	userId := strconv.Itoa(int(userID))

	eventID := "eventlogAction"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "logAction", Args: [][]byte{[]byte(userId), []byte(action),
		[]byte(details), []byte(eventID)}}

	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil

}

func (t *ServiceSetup) QueryLogs() (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryLogs", Args: [][]byte{}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}

func (t *ServiceSetup) QueryUserLogs(userID string) (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryUserLogs", Args: [][]byte{[]byte(userID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return "", err
	}

	return string(respone.Payload), nil
}
