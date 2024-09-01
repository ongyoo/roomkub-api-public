package rabbitmq

// Exchange
const (
	ExchangePaymentCMS             = "payment.cms"
	ExchangePaymentCMSRetry        = "payment.cms.retry"
	ExchangePaymentOnboarding      = "payment.onboarding"
	ExchangePaymentSettlement      = "payment.settlement"
	ExchangePaymentSettlementRetry = "payment.settlement.retry"
	ExchangePaymentRefund          = "payment.refund"
	ExchangePaymentRefundRetry     = "payment.refund.retry"
)

// Queue
const (
	QueuePaymentCMSUnsuspendPayoutMerchant = "payment.cms.unsuspend.payout.merchant.event"
	QueuePaymentCMSUnsuspendPayoutShop     = "payment.cms.unsuspend.payout.shop.event"
	QueuePaymentSettlement                 = "payment.settlement.event"
	QueuePaymentRefund                     = "payment.refund.event"
)

// Routing Key
const (
	RoutingKeyUnsuspendPayoutMerchant  = "unsuspend-payout-merchant"
	RoutingKeyUnsuspendPayoutShop      = "unsuspend-payout-shop"
	RoutingKeySuspendPayoutMerchant    = "suspend-payout-merchant"
	RoutingKeySuspendPayoutShop        = "suspend-payout-shop"
	RoutingKeyUnsuspendPaymentMerchant = "unsuspend-payment-merchant"
	RoutingKeyUnsuspendPaymentShop     = "unsuspend-payment-shop"
	RoutingKeySuspendPaymentMerchant   = "suspend-payment-merchant"
	RoutingKeySuspendPaymentShop       = "suspend-payment-shop"
	RoutingKeyApproveForm              = "approve-form"
	RoutingKeyRejectForm               = "reject-form"
	RoutingKeyRetry                    = "*"
	RoutingKeySettlement               = "transferable"
	RoutingKeyRefund                   = "refund"
	RoutingKeyUnbanForm                = "unban-form"
	RoutingKeyBanForm                  = "ban-form"
	RoutingKeyApproveFormChange        = "approve-form-change"
	RoutingKeyRejectFormChange         = "reject-form-change"
	RoutingKeyResetFormChange          = "reset-form"
	RoutingKeyRequestChange            = "request-change"
)
