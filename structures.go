package mercadopago

type AuthToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	LiveMode     bool   `json:"live_mode"`
	UserID       int64  `json:"user_id"`
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
	ID             string         `json:"id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Phone          Phone          `json:"phone"`
	Identification Identification `json:"identification"`
	Email          string         `json:"email"`
	Nickname       string         `json:"nickname"`
}

type Collector struct {
	ID        int64  `json:"id"`
	FirtsName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     Phone  `json:"phone"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
}

type Collection struct {
	ID                  int64                  `json:"id"`
	SiteID              string                 `json:"site_id"`
	DateCreated         string                 `json:"date_created"`
	DateApproved        string                 `json:"date_approved"`
	MoneyReleaseDate    string                 `json:"money_release_date"`
	LastModified        string                 `json:"last_modified"`
	Payer               Payer                  `json:"payer"`
	OrderID             string                 `json:"order_id"`
	ExternalReference   string                 `json:"external_reference"`
	MerchantOrderID     string                 `json:"merchant_order_id"`
	Reason              string                 `json:"reason"`
	CurrencyID          string                 `json:"currency_id"`
	TransactionAmount   float64                `json:"transaction_amount"`
	NetReceivedAmount   float64                `json:"net_received_amount"`
	TotalPaidAmount     float64                `json:"total_paid_amount"`
	ShippingCost        float64                `json:"shipping_cost"`
	CouponAmount        float64                `json:"coupon_amount"`
	CouponFee           float64                `json:"coupon_fee"`
	FinanceFee          float64                `json:"finance_fee"`
	DiscountFee         float64                `json:"discount_fee"`
	CouponID            string                 `json:"coupon_id"`
	Status              string                 `json:"status"`
	StatusDetails       string                 `json:"status_details"`
	IssuerID            string                 `json:"issuer_id"`
	InstallmentAmount   float64                `json:"installment_amount"`
	DeferredPeriod      string                 `json:"deffered_period"`
	PaymentType         string                 `json:"payment_type"`
	PaymentMethodID     string                 `json:"payment_method_id"`
	Marketplace         string                 `json:"marketplace"`
	OperationType       string                 `json:"operation_type"`
	MarketplaceFee      float64                `json:"marketplace_fee"`
	DeductionSchema     string                 `json:"deduction_schema"`
	Refunds             []interface{}          `json:"refunds"`
	AmountRefunded      float64                `json:"amount_refunded"`
	LastModifiedByAdmin string                 `json:"last_modified_by_admin"`
	APIVersion          string                 `json:"api_version"`
	ConceptID           string                 `json:"concept_id"`
	ConceptAmount       float64                `json:"concept_amount"`
	InternalMetadata    map[string]interface{} `json:"internal_metadata"`
	Collector           Collector              `json:"collector"`
}

type Payment struct {
	Collection Collection `json:"collection"`
}

type Item struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	CurrencyID  string  `json:"currency_id"`
	PictureURL  string  `json:"picture_url"`
	Description string  `json:"description"`
	CategoryID  string  `json:"category_id"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float32 `json:"unit_price"`
}

type PreferencePayer struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Phone       struct {
		AreaCode string `json:"area_code"`
		Number   string `json:"number"`
	} `json:"phone"`
	Identification struct {
		Type   string `json:"type"`
		Number string `json:"number"`
	} `json:"identification"`
	Address struct {
		StreetName   string `json:"street_name"`
		StreetNumber int    `json:"street_number"`
		ZipCode      string `json:"zip_code"`
	} `json:"address"`
}

type PaymentMethodID struct {
	ID string `json:"id"`
}

type PaymentMethods struct {
	ExcludedPaymentMethods []PaymentMethodID `json:"excluded_payment_methods"`
	ExcludedPaymentTypes   []PaymentMethodID `json:"excluded_payment_types"`
	Installments           int               `json:"installments"`
	DefaultPaymentMethodID interface{}       `json:"default_payment_method_id"`
	DefaultInstallments    interface{}       `json:"default_installments"`
}

type Preference struct {
	Items    []Item          `json:"items"`
	Payer    PreferencePayer `json:"payer"`
	BackUrls struct {
		Success string `json:"success"`
		Failure string `json:"failure"`
		Pending string `json:"pending"`
	} `json:"back_urls"`
	AutoReturn     string         `json:"auto_return"`
	PaymentMethods PaymentMethods `json:"payment_methods"`
	Shipments      struct {
		ReceiverAddress struct {
			ZipCode      string `json:"zip_code"`
			StreetNumber int    `json:"street_number"`
			StreetName   string `json:"street_name"`
			Floor        int    `json:"floor"`
			Apartment    string `json:"apartment"`
		} `json:"receiver_address"`
	} `json:"shipments"`
	NotificationURL    string `json:"notification_url"`
	ExternalReference  string `json:"external_reference"`
	Expires            bool   `json:"expires"`
	ExpirationDateFrom string `json:"expiration_date_from"`
	ExpirationDateTo   string `json:"expiration_date_to"`
}
