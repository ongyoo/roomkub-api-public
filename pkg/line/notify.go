package line

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
)

type NotifyType string

const (
	// CART
	NotifyTypeCart NotifyType = "CART"
	// Stock
	NotifyTypeUpdateStock   NotifyType = "UPDATE_STOCK"
	NotifyTypeStockDepleted NotifyType = "DEPLETED"
	NotifyTypeStockLow      NotifyType = "STOCK_LOW"

	// User
	NotifyTypeUserLogin NotifyType = "USER_LOGIN"
	// Test
	NotifyTypeTest NotifyType = "TEST"
)

type LineNotifyMessage struct {
	Message   string
	Thumbnail string
	Type      NotifyType
}

func SendMessageLineNotify(ctx context.Context, message LineNotifyMessage) {
	messageFormat := generateMessage(message)
	userBy := ""
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err == nil {
		userBy = "ทำรายการโดย คุณ "
		userBy += userClaims.Payload.FirstName + " " + userClaims.Payload.LastName + " ( " + userClaims.Payload.NickName + " )"
	}
	messageFormat += " " + userBy
	err = sendLineNotify(messageFormat, message.Thumbnail)
	if err != nil {
		fmt.Println(err)
	}
}

func generateMessage(message LineNotifyMessage) string {
	if message.Type == NotifyTypeCart {
		return message.Message
	}

	if message.Type == NotifyTypeStockDepleted {
		return message.Message + " [สินค้าหมด กรุณาเติมสินค้า]"
	}

	if message.Type == NotifyTypeStockLow {
		return message.Message + " ชิ้น [สินค้าใกล้หมด กรุณาเติมสินค้า]"
	}

	// User
	if message.Type == NotifyTypeUserLogin {
		return "[มีการล็อคอินเข้าใช้ระบบโดย] " + message.Message
	}

	// TEST
	if message.Type == NotifyTypeTest {
		return "[Test] ทดสอบ Notify " + message.Message
	}
	return message.Message
}

func sendLineNotify(message, imageUrl string) error {
	resty := resty.New()
	urlEndPoint := "https://notify-api.line.me/api/notify"
	// Enabling debug
	resty.SetDebug(true)

	// Sets `Content-Length` header automatically
	resty.SetContentLength(true)

	// resty will set content-type to `application/x-www-form-urlencoded`
	// automatically
	resp, err := resty.R().
		SetFormData(map[string]string{
			"message":        message,
			"imageThumbnail": imageUrl,
			"imageFullsize":  imageUrl,
		}).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("").
		Post(urlEndPoint)
	fmt.Println(err, resp)
	if err != nil {
		return err
	}
	return nil
}
