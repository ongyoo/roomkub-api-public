package paymentMethod

import (
	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod struct {
	ID                        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	BusinessID                primitive.ObjectID `json:"business_id" bson:"businessId"`
	Name                      mongo.Encrypted    `json:"name" bson:"name"`
	NameIdentityID            string             `json:"name_identity_id" bson:"nameIdentityId"`
	Description               mongo.Encrypted    `json:"description" bson:"description"`
	BankAccountNo             mongo.Encrypted    `json:"bank_account_no" bson:"bankAccountNo"`
	BankAccountName           mongo.Encrypted    `json:"bank_account_name" bson:"bankAccountName"`
	BankAccountNameIdentityID string             `json:"bank_account_name_identity_id" bson:"bankAccountNameIdentityId"`
	IsActive                  bool               `json:"is_active" bson:"isActive"`
}
