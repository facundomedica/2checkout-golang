package mercadopago

import (
	"encoding/json"
)

const apiBaseUrl = "https://api.mercadopago.com"

// MpClient
//
// MercadoPago client
type MpClient struct {
	clientId     string
	clientSecret string
	accessToken  string
	rest         *RestClient
	sandbox      bool
}

// Init
//
// Configure main data for MercadoPago
func (mp *MpClient) Init(args ...string) {
	if len(args) < 1 || len(args) > 2 {
		panic("Invalid arguments. Use CLIENT_ID and CLIENT SECRET, or ACCESS_TOKEN")
	}

	if len(args) == 2 {
		mp.clientId = args[0]
		mp.clientSecret = args[1]
	}

	if len(args) == 1 {
		mp.clientSecret = args[0]
	}

	mp.rest = new(RestClient)
	mp.sandbox = false
}

// SandboxMode
//
// Change sandbox mode
func (mp *MpClient) SandboxMode(mode bool) bool {
	mp.sandbox = mode
	return mp.sandbox
}

// GetAccessToken
//
// Get access credentials
func (mp *MpClient) GetAccessToken() (string, error) {
	if mp.accessToken != "" {
		return mp.accessToken, nil
	}

	appClientValues := RequestData{
		"grant_type":    "client_credentials",
		"client_id":     mp.clientId,
		"client_secret": mp.clientSecret,
	}

	url := apiBaseUrl + "/oauth/token"
	accessData, err := mp.rest.Post(url, appClientValues, nil)
	if err != nil {
		return "", err
	}

	var token AuthToken
	err = json.Unmarshal(accessData, &token)

	return token.AccessToken, err
}

// GetPayment
//
// Get payment infoirmation by id
func (mp *MpClient) GetPayment(id string) (Collection, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return Collection{}, err
	}

	uriPrefix := ""
	if mp.sandbox {
		uriPrefix = "/sandbox"
	}

	url := apiBaseUrl + uriPrefix + "/collections/notifications/" + id + "?access_token=" + accessToken

	response, errReq := mp.rest.Get(url, nil, nil)
	if errReq != nil {
		return Collection{}, errReq
	}

	var payment Collection
	err = json.Unmarshal(response, &payment)

	return payment, err
}

// GetAuthorizedPayment
//
// Get authorized paymnet
func (mp *MpClient) GetAuthorizedPayment(id string) (string, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return "", err
	}

	url := apiBaseUrl + "/authorized_payments/" + id + "?access_token=" + accessToken
	response, errReq := mp.rest.Get(url, nil, nil)

	return string(response), errReq
}

// Post
//
// generic post request to MP API
func (mp *MpClient) CreatePayment(data RequestData) ([]byte, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := apiBaseUrl + "/v1/payments?access_token=" + accessToken
	return mp.rest.Post(url, data, nil)
}
