package braintree

const (
	// Transaction Status Codes
	TxAuthorizationExpired   string = "authorization_expired"
	TxAuthorizing            string = "authorizing"
	TxAuthorized             string = "authorized"
	TxGatewayRejected        string = "gateway_rejected"
	TxFailed                 string = "failed"
	TxProcessorDeclined      string = "processor_declined"
	TxSettled                string = "settled"
	TxSettling               string = "settling"
	TxSubmittedForSettlement string = "submitted_for_settlement"
	TxVoided                 string = "voided"
	TxUnrecognized           string = "unrecognized"

	// Transaction Escrow Status
	TxEscrowHoldPending    string = "hold_pending"
	TxEscrowHeld           string = "held"
	TxEscrowReleasePending string = "release_pending"
	TxEscrowReleased       string = "released"
	TxEscrowRefunded       string = "refunded"

	// Transaction Types
	TxSale   string = "sale"
	TxCredit string = "credit"

	// Transaction Created Using
	TxFullInformation string = "full_information"
	TxToken           string = "token"

	// Transaction Sources
	TxAPI          string = "api"
	TxControlPanel string = "control_panel"
	TxRecurring    string = "recurring"

	// Gateway Rejection Reason
	TxAVS       string = "avs"
	TxAVSAndCVV string = "avs_and_cvv"
	TxCVV       string = "cvv"
	TxDuplicate string = "duplicate"
	TxFraud     string = "fraud"
)

type Transaction struct {
	XMLName            string              `xml:"transaction"`
	Id                 string              `xml:"id,omitempty"`
	CustomerID         string              `xml:"customer-id,omitempty"`
	Status             string              `xml:"status,omitempty"`
	EscrowStatus       string              `xml:"escrow-status,omitempty"`
	Type               string              `xml:"type,omitempty"`
	Amount             float64             `xml:"amount"`
	OrderId            string              `xml:"order-id,omitempty"`
	PaymentMethodToken string              `xml:"payment-method-token,omitempty"`
	MerchantAccountId  string              `xml:"merchant-account-id,omitempty"`
	PlanId             string              `xml:"plan-id,omitempty"`
	CreditCard         *CreditCard         `xml:"credit-card,omitempty"`
	Customer           *Customer           `xml:"customer,omitempty"`
	BillingAddress     *Address            `xml:"billing,omitempty"`
	ShippingAddress    *Address            `xml:"shipping,omitempty"`
	Options            *TransactionOptions `xml:"options,omitempty"`
	ServiceFeeAmount   float64             `xml:"service-fee-amount,attr,omitempty"`
	CreatedAt          string              `xml:"created-at,omitempty"`
	UpdatedAt          string              `xml:"updated-at,omitempty"`
	CurrencyISOCode    string              `xml:"currency-iso-code,omitempty"`
	Authorization      string              `xml:"processor-authorization-code,omitempty"`
}

// TODO: not all transaction fields are implemented yet, here are the missing fields (add on demand)
//
// <transaction>
//   <currency-iso-code>USD</currency-iso-code>
//   <refund-id nil="true"></refund-id>
//   <refund-ids type="array"/>
//   <refunded-transaction-id nil="true"></refunded-transaction-id>
//   <settlement-batch-id>2013-10-08_49grybq7pbtsnvsr</settlement-batch-id>
//   <custom-fields>
//   </custom-fields>
//   <avs-error-response-code nil="true"></avs-error-response-code>
//   <avs-postal-code-response-code>I</avs-postal-code-response-code>
//   <avs-street-address-response-code>I</avs-street-address-response-code>
//   <cvv-response-code>I</cvv-response-code>
//   <gateway-rejection-reason nil="true"></gateway-rejection-reason>
//   <processor-authorization-code>YCSBWR</processor-authorization-code>
//   <processor-response-code>1000</processor-response-code>
//   <processor-response-text>Approved</processor-response-text>
//   <voice-referral-number nil="true"></voice-referral-number>
//   <purchase-order-number nil="true"></purchase-order-number>
//   <tax-amount nil="true"></tax-amount>
//   <tax-exempt type="boolean">false</tax-exempt>
//   <status-history type="array">
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>authorized</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-07T17:26:14Z</timestamp>
//       <status>submitted_for_settlement</status>
//       <amount>7.00</amount>
//       <user>eaigner</user>
//       <transaction-source>Recurring</transaction-source>
//     </status-event>
//     <status-event>
//       <timestamp type="datetime">2013-10-08T07:06:38Z</timestamp>
//       <status>settled</status>
//       <amount>7.00</amount>
//       <user nil="true"></user>
//       <transaction-source></transaction-source>
//     </status-event>
//   </status-history>
//   <plan-id>bronze</plan-id>
//   <subscription-id>jqsydb</subscription-id>
//   <subscription>
//     <billing-period-end-date type="date">2013-11-06</billing-period-end-date>
//     <billing-period-start-date type="date">2013-10-07</billing-period-start-date>
//   </subscription>
//   <add-ons type="array"/>
//   <discounts type="array"/>
//   <descriptor>
//     <name nil="true"></name>
//     <phone nil="true"></phone>
//   </descriptor>
//   <recurring type="boolean">true</recurring>
//   <channel nil="true"></channel>
//   <disbursement-details>
//     <disbursement-date type="date">2013-10-08</disbursement-date>
//     <settlement-amount>7.00</settlement-amount>
//     <settlement-currency-iso-code>USD</settlement-currency-iso-code>
//     <settlement-currency-exchange-rate>1</settlement-currency-exchange-rate>
//     <funds-held type="boolean">false</funds-held>
//   </disbursement-details>
// </transaction>

type Transactions struct {
	Transaction []*Transaction `xml:"transaction"`
}

type TransactionOptions struct {
	SubmitForSettlement              bool `xml:"submit-for-settlement,omitempty"`
	StoreInVault                     bool `xml:"store-in-vault,omitempty"`
	AddBillingAddressToPaymentMethod bool `xml:"add-billing-address-to-payment-method,omitempty"`
	StoreShippingAddressInVault      bool `xml:"store-shipping-address-in-vault,omitempty"`
	HoldInEscrow                     bool `xml:"hold-in-escrow,omitempty"`
}

type TransactionSearchResult struct {
	XMLName           string         `xml:"credit-card-transactions"`
	CurrentPageNumber string         `xml:"current-page-number"` // int
	PageSize          string         `xml:"page-size"`           // int
	TotalItems        string         `xml:"total-items"`         // int
	Transactions      []*Transaction `xml:"transaction"`
}
