package receipt

import (
	"time"
)

type AppleReceipt struct {
	Status             int                    `json:"status"`
	Receipt            ApplicationReceipt     `json:"receipt"`
	LatestReceiptInfo  []InAppPurchaseReceipt `json:"lastest_receipt_info"`
	PendingRenewalInfo string                 `json:"pending_renewal_info"`
	isRetryable        int                    `json:"is_retryable"`
}

type ApplicationReceipt struct {
	BundleIdentifier           string
	AppVersion                 string
	OpaqueValue                string
	InAppPurchaseReceipt       []InAppPurchaseReceipt
	OriginalApplicationVersion string
	ReceiptCreationDate        time.Time
	ReceiptExpirationDate      time.Time
}

type InAppPurchaseReceipt struct {
	Quantity                            int
	ProductIdentifier                   string
	TransactionIdentifier               string
	OriginalTransactionIdentifier       string
	PurchaseDate                        time.Time
	OriginalPurchaseDate                time.Time
	SubscriptionExpirationDate          time.Time
	SubscriptionExpirationIntent        int
	SubscriptionRetryFlag               int
	SubsxcriptionTrialPeriod            time.Time
	SubscriptionIntroductoryPricePeriod time.Time
	CancellationDate                    time.Time
	CancellationReason                  string
	AppItemID                           string
	ExternalVersionIdentifier           string
	WebOrderLineItemID                  string
	SubscriptionAutoRenewStatus         int
	SubscriptionAutoRenewPreference     string
	SubscriptionPriceConsentStatus      int
}
