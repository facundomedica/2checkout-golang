package twocheckout

import (
	"encoding/json"
	"log"
)

const sandboxAPIBaseURL = "sandbox.2checkout.com/"
const apiBaseURL = "www.2checkout.com/"

// Client is the 2Checkout client object
type Client struct {
	APIUser     string
	APIPassword string
	VendorID    string
	rest        *RestClient
	sandbox     bool
}

// New instatiates a 2Checkout client
func New(APIUser string, APIPassword string, VendorID string) *Client {
	tcc := Client{}
	if APIUser == "" || APIPassword == "" {
		return nil
	}
	tcc.APIUser = APIUser
	tcc.APIPassword = APIPassword
	tcc.VendorID = VendorID
	tcc.rest = new(RestClient)
	tcc.sandbox = false
	return &tcc
}

// SandboxMode changes sandbox mode
func (tcc *Client) SandboxMode(mode bool) bool {
	tcc.sandbox = mode
	return tcc.sandbox
}

func (tcc *Client) getBasicAuthURL() string {
	url := apiBaseURL
	if tcc.sandbox == true {
		url = sandboxAPIBaseURL
	}
	return "https://" + tcc.APIUser + ":" + tcc.APIPassword + "@" + url
}

func (tcc *Client) DetailSale(saleId string, invoiceId string) (SaleDetail, error) {
	url := tcc.getBasicAuthURL() + "api/sales/detail_sale"
	data := ""

	if saleId != "" {
		data = "sale_id=" + saleId
	}

	if invoiceId != "" {
		data = "invoice_id=" + invoiceId
	}

	response, errReq := tcc.rest.PostWithStringData(url, data, nil)
	log.Println(string(response))
	if errReq != nil {
		return SaleDetail{}, errReq
	}
	var saleDetail SaleDetail
	err := json.Unmarshal(response, &saleDetail)
	return saleDetail, err
}
