package generator

import (
	"fmt"
	"time"
)

func GeneratorOrderID(count int64, orderType string) string {
	id := orderType[0:1]
	orderID := fmt.Sprintf("%s-%06d", id, count+1)
	return orderID
}

func GeneratorTransactionID(count int64) string {
	//T-{time.Now().Unix()}-{count+1}
	transactionID := fmt.Sprintf("T-%d-%06d", time.Now().Unix(), count+1)
	return transactionID
}
