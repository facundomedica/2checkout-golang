package twocheckout

type SaleDetail struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Sale            struct {
		Comments []interface{} `json:"comments"`
		Customer struct {
			Address1       string      `json:"address_1"`
			Address2       string      `json:"address_2"`
			AddressID      string      `json:"address_id"`
			CardholderName string      `json:"cardholder_name"`
			City           string      `json:"city"`
			CountryCode    string      `json:"country_code"`
			CountryCode2   string      `json:"country_code_2"`
			CountryName    string      `json:"country_name"`
			CustomerID     string      `json:"customer_id"`
			EmailAddress   string      `json:"email_address"`
			FirstName      string      `json:"first_name"`
			Lang           string      `json:"lang"`
			LastName       string      `json:"last_name"`
			MiddleInitial  interface{} `json:"middle_initial"`
			PayMethod      struct {
				Avs            interface{} `json:"avs"`
				Cvv            interface{} `json:"cvv"`
				FirstSixDigits string      `json:"first_six_digits"`
				LastTwoDigits  string      `json:"last_two_digits"`
				Method         string      `json:"method"`
			} `json:"pay_method"`
			Phone      string      `json:"phone"`
			PhoneExt   string      `json:"phone_ext"`
			PostalCode string      `json:"postal_code"`
			Prefix     interface{} `json:"prefix"`
			State      string      `json:"state"`
			State2     interface{} `json:"state_2"`
		} `json:"customer"`
		DatePlaced string `json:"date_placed"`
		DetailIP   struct {
			Address     string `json:"address"`
			AreaCode    int    `json:"area_code"`
			City        string `json:"city"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			Region      string `json:"region"`
			Zip         string `json:"zip"`
		} `json:"detail_ip"`
		Invoices []struct {
			ChangeID           string      `json:"change_id"`
			CustomerName       string      `json:"customer_name"`
			CustomerSnapshotID interface{} `json:"customer_snapshot_id"`
			CustomerTotal      string      `json:"customer_total"`
			DatePlaced         string      `json:"date_placed"`
			DateShipped        interface{} `json:"date_shipped"`
			DateVendorPaid     interface{} `json:"date_vendor_paid"`
			Fees2Co            string      `json:"fees_2co"`
			InvoiceID          string      `json:"invoice_id"`
			Lineitems          []struct {
				AffiliateVendorID interface{} `json:"affiliate_vendor_id"`
				Billing           struct {
					Amount          string      `json:"amount"`
					BillMethod      string      `json:"bill_method"`
					BillingID       string      `json:"billing_id"`
					CustomerAmount  string      `json:"customer_amount"`
					CustomerID      string      `json:"customer_id"`
					DateDeposited   interface{} `json:"date_deposited"`
					DateEnd         string      `json:"date_end"`
					DateFail        string      `json:"date_fail"`
					DateNext        interface{} `json:"date_next"`
					DatePending     interface{} `json:"date_pending"`
					DateStart       string      `json:"date_start"`
					LineitemID      string      `json:"lineitem_id"`
					RecurringStatus interface{} `json:"recurring_status"`
					Status          string      `json:"status"`
					UsdAmount       string      `json:"usd_amount"`
					VendorAmount    string      `json:"vendor_amount"`
				} `json:"billing"`
				Commission                  interface{}   `json:"commission"`
				CommissionAffiliateVendorID interface{}   `json:"commission_affiliate_vendor_id"`
				CommissionFlatRate          interface{}   `json:"commission_flat_rate"`
				CommissionPercentage        interface{}   `json:"commission_percentage"`
				CommissionType              interface{}   `json:"commission_type"`
				CommissionUsdAmount         interface{}   `json:"commission_usd_amount"`
				CustomerAmount              string        `json:"customer_amount"`
				FlatRate                    interface{}   `json:"flat_rate"`
				Installment                 interface{}   `json:"installment"`
				InvoiceID                   string        `json:"invoice_id"`
				LcAffiliateVendorID         interface{}   `json:"lc_affiliate_vendor_id"`
				LcUsdAmount                 interface{}   `json:"lc_usd_amount"`
				LineitemID                  string        `json:"lineitem_id"`
				LinkedID                    interface{}   `json:"linked_id"`
				Options                     []interface{} `json:"options"`
				Percentage                  interface{}   `json:"percentage"`
				ProductDescription          string        `json:"product_description"`
				ProductDuration             string        `json:"product_duration"`
				ProductHandling             string        `json:"product_handling"`
				ProductID                   string        `json:"product_id"`
				ProductIsCart               string        `json:"product_is_cart"`
				ProductName                 string        `json:"product_name"`
				ProductPrice                string        `json:"product_price"`
				ProductRecurrence           string        `json:"product_recurrence"`
				ProductStartupFee           interface{}   `json:"product_startup_fee"`
				ProductTangible             string        `json:"product_tangible"`
				SaleID                      string        `json:"sale_id"`
				Status                      string        `json:"status"`
				Type                        string        `json:"type"`
				UsdAmount                   string        `json:"usd_amount"`
				UsdCommission               interface{}   `json:"usd_commission"`
				VendorAmount                string        `json:"vendor_amount"`
				VendorProductID             string        `json:"vendor_product_id"`
			} `json:"lineitems"`
			NewCustomerQuoteID string      `json:"new_customer_quote_id"`
			NewVendorQuoteID   string      `json:"new_vendor_quote_id"`
			PayoutStatus       string      `json:"payout_status"`
			Recurring          string      `json:"recurring"`
			Referrer           interface{} `json:"referrer"`
			SaleID             string      `json:"sale_id"`
			Shipping           interface{} `json:"shipping"`
			Status             string      `json:"status"`
			Timestamp          string      `json:"timestamp"`
			UsdTotal           string      `json:"usd_total"`
			VendorID           string      `json:"vendor_id"`
			VendorOrderID      interface{} `json:"vendor_order_id"`
			VendorTotal        string      `json:"vendor_total"`
		} `json:"invoices"`
		IPAddress        string      `json:"ip_address"`
		IPCountry        string      `json:"ip_country"`
		RecurringDecline interface{} `json:"recurring_decline"`
		SaleID           string      `json:"sale_id"`
	} `json:"sale"`
	Warnings []struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		Parameter string `json:"parameter"`
	} `json:"warnings"`
}
