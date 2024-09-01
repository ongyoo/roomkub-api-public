package PaymentTransaction

import (
	"time"

	mongo "github.com/ongyoo/roomkub-api/pkg/database"
	"github.com/ongyoo/roomkub-api/pkg/document"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionStatus string

const (
	TransactionStatusPaid              TransactionStatus = "PAID"     // จ่ายแล้ว
	TransactionStatusWaitingForApprove TransactionStatus = "WFA"      // Waiting For Approve รอตรวจสอบ
	TransactionStatusApproved          TransactionStatus = "APPROVED" // Approved //อนุมัติ
	TransactionStatusReject            TransactionStatus = "REJECT"   // Reject ปฏิเสธการรับเงิน
)

type TransactionType string

const (
	// Income
	TransactionTypeIncome TransactionType = "INCOME" // Income รายรับ
	// Expenses
	TransactionTypeExpenses TransactionType = "EXPENSES" // Expenses รายจ่าย
	TransactionTypePayRoll  TransactionType = "PAYROLL"  // Payroll เงินเดือน
	TransactionTypeWithdraw TransactionType = "WITHDRAW" // Withdraw เบิก
)

type TransactionWithdrawType string

const (
	// Withdraw เบิก
	TransactionWithdrawTypeRepair        TransactionWithdrawType = "REPAIR"        // ซ่อมแซม
	TransactionWithdrawTypeTechnicianFee TransactionWithdrawType = "TECHNICIAN"    // ค่าช่าง
	TransactionWithdrawTypeEquipmentFee  TransactionWithdrawType = "EQUIPMENT_FEE" // ค่าอุปกรณ์ หรือ อะไหล่ วัสดุ
	TransactionWithdrawTypeOther         TransactionWithdrawType = "OTHER"         // ค่าอื่นๆ
)

type Transaction struct {
	ID              primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	InvoiceID       primitive.ObjectID  `json:"invoice_id" bson:"invoiceId"`
	PaymentMethodID primitive.ObjectID  `json:"payment_method" bson:"paymentMethod"`
	Documents       []document.Document `json:"documents" bson:"documents"` // ข้อมูลรูปต่างๆ
	Status          TransactionStatus   `json:"status" bson:"status"`
	Type            TransactionType     `json:"type" bson:"type"`
	Remark          mongo.Encrypted     `json:"remark" bson:"remark"`
	Ref1            mongo.Encrypted     `json:"ref1" bson:"ref1"`                     // อ้างอิง 1
	Ref2            mongo.Encrypted     `json:"ref2" bson:"ref2"`                     // อ้างอิง 2
	IsCash          bool                `json:"is_cash" bson:"isCash"`                // รับเงินสด
	IsAutomatic     bool                `json:"is_automatic" bson:"isAutomatic"`      // ระบบเป็นคนจัดการ
	AssignUserID    primitive.ObjectID  `json:"assign_user_id" bson:"assign_user_id"` // กำหนดผู้รับ รับเงินสด / รับเงินเดือน / เบิก
	CreateUserID    primitive.ObjectID  `json:"create_user_id" bson:"create_user_id"` // ผู้สร้าง
	CreateAt        time.Time           `json:"create_at" bson:"createAt"`
	UpdateAt        time.Time           `json:"update_at" bson:"updateAt"`
}
