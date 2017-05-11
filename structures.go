package mercadopago

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	LiveMode     bool   `json:"live_mode"`
	UserId       int64  `json:"user_id"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_id"`
	Scope        string `json:"scope"`
}

type Phone struct {
	AreaCode  string `json:"area_code"`
	Number    string `json:"number"`
	Extension string `json:"extension"`
}

type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Payer struct {
	Id             int64          `json:"id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Phone          Phone          `json:"phone"`
	Identification Identification `json:"identification"`
	Email          string         `json:"email"`
	Nickname       string         `json:"nickname"`
}

type Collector struct {
	Id        int64  `json:"id"`
	FirtsName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     Phone  `json:"phone"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
}

type Collection struct {
	Id                  int64                  `json:"id"`
	SiteId              string                 `json:"site_id"`
	DateCreated         string                 `json:"date_created"`
	DateApproved        string                 `json:"date_approved"`
	MoneyReleaseDate    string                 `json:"money_release_date"`
	LastModified        string                 `json:"last_modified"`
	Payer               Payer                  `json:"payer"`
	OrderId             string                 `json:"order_id"`
	ExternalReference   string                 `json:"external_reference"`
	MerchantOrderId     string                 `json:"merchant_order_id"`
	Reason              string                 `json:"reason"`
	CurrencyId          string                 `json:"currency_id"`
	TransactionAmount   float64                `json:"transaction_amount"`
	NetReceivedAmount   float64                `json:"net_received_amount"`
	TotalPaidAmount     float64                `json:"total_paid_amount"`
	ShippingCost        float64                `json:"shipping_cost"`
	CouponAmount        float64                `json:"coupon_amount"`
	CouponFee           float64                `json:"coupon_fee"`
	FinanceFee          float64                `json:"finance_fee"`
	DiscountFee         float64                `json:"discount_fee"`
	CouponId            string                 `json:"coupon_id"`
	Status              string                 `json:"status"`
	StatusDetails       string                 `json:"status_details"`
	IssuerId            string                 `json:"issuer_id"`
	InstallmentAmount   float64                `json:"installment_amount"`
	DeferredPeriod      string                 `json:"deffered_period"`
	PaymentType         string                 `json:"payment_type"`
	PaymentMethodId     string                 `json:"payment_method_id"`
	Marketplace         string                 `json:"marketplace"`
	OperationType       string                 `json:"operation_type"`
	MarketplaceFee      float64                `json:"marketplace_fee"`
	DeductionSchema     string                 `json:"deduction_schema"`
	Refunds             []interface{}          `json:"refunds"`
	AmountRefunded      float64                `json:"amount_refunded"`
	LastModifiedByAdmin string                 `json:"last_modified_by_admin"`
	ApiVersion          string                 `json:"api_version"`
	ConceptId           string                 `json:"concept_id"`
	ConceptAmount       float64                `json:"concept_amount"`
	InternalMetadata    map[string]interface{} `json:"internal_metadata"`
	Collector           Collector              `json:"collector"`
}

type Payment struct {
	Collection Collection `json:"collection"`
}
