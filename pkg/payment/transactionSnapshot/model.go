package transactionSnapshot

import (
	"time"

	PaymentTransaction "github.com/ongyoo/roomkub-api/pkg/payment/transaction"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionSnapshot struct {
	ID          primitive.ObjectID             `json:"id" bson:"_id,omitempty"`
	Transaction PaymentTransaction.Transaction `json:"transaction" bson:"transaction"`
	CreateAt    time.Time                      `json:"create_at" bson:"createAt"`
}
