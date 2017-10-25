package mercadopago

import (
	"encoding/json"
)

const apiBaseURL = "https://api.mercadopago.com"

// MpClient is the MercadoPago client object
type MpClient struct {
	clientID     string
	clientSecret string
	accessToken  string
	rest         *RestClient
	sandbox      bool
}

// Init configures main data for MercadoPago
func (mp *MpClient) Init(args ...string) {
	if len(args) < 1 || len(args) > 2 {
		panic("Invalid arguments. Use CLIENT_ID and CLIENT SECRET, or ACCESS_TOKEN")
	}

	if len(args) == 2 {
		mp.clientID = args[0]
		mp.clientSecret = args[1]
	}

	if len(args) == 1 {
		mp.accessToken = args[0]
	}

	mp.rest = new(RestClient)
	mp.sandbox = false
}

// SandboxMode changes sandbox mode
func (mp *MpClient) SandboxMode(mode bool) bool {
	mp.sandbox = mode
	return mp.sandbox
}

// GetAccessToken gets access credentials
func (mp *MpClient) GetAccessToken() (string, error) {
	if mp.accessToken != "" {
		return mp.accessToken, nil
	}

	appClientValues := RequestData{
		"grant_type":    "client_credentials",
		"client_id":     mp.clientID,
		"client_secret": mp.clientSecret,
	}

	url := apiBaseURL + "/oauth/token"
	accessData, err := mp.rest.Post(url, appClientValues, nil)
	if err != nil {
		return "", err
	}

	var token AuthToken
	err = json.Unmarshal(accessData, &token)

	return token.AccessToken, err
}

// GetPayment gets payment information by id
func (mp *MpClient) GetPayment(id string) (Collection, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return Collection{}, err
	}

	uriPrefix := ""
	if mp.sandbox {
		uriPrefix = "/sandbox"
	}

	url := apiBaseURL + uriPrefix + "v1/payments/" + id + "?access_token=" + accessToken

	response, errReq := mp.rest.Get(url, nil, nil)
	if errReq != nil {
		return Collection{}, errReq
	}

	var payment Collection
	err = json.Unmarshal(response, &payment)

	return payment, err
}

// GetAuthorizedPayment gets authorized paymnet
func (mp *MpClient) GetAuthorizedPayment(id string) (string, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return "", err
	}

	url := apiBaseURL + "/authorized_payments/" + id + "?access_token=" + accessToken
	response, errReq := mp.rest.Get(url, nil, nil)

	return string(response), errReq
}

// CreatePayment creates a payment using data
func (mp *MpClient) CreatePayment(data RequestData) ([]byte, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return nil, err
	}

	url := apiBaseURL + "/v1/payments?access_token=" + accessToken
	return mp.rest.Post(url, data, nil)
}

//CreatePreference creates a new preference to make a payment
func (mp *MpClient) CreatePreference(preference Preference) ([]byte, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return nil, err
	}

	payloadByte, err := json.Marshal(preference)
	if err != nil {
		return nil, err
	}

	uriPrefix := ""
	if mp.sandbox {
		uriPrefix = "/sandbox"
	}

	url := apiBaseURL + uriPrefix + "/checkout/preferences?access_token=" + accessToken
	return mp.rest.PostWithStringData(url, string(payloadByte), nil)
}

//GetPaymentMethods gets the available payment methods
func (mp *MpClient) GetPaymentMethods() ([]byte, error) {
	accessToken, err := mp.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uriPrefix := ""
	if mp.sandbox {
		uriPrefix = "/sandbox"
	}

	url := apiBaseURL + uriPrefix + "/v1/payment_methods?access_token=" + accessToken
	return mp.rest.Get(url, nil, nil)
}
