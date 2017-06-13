# MercadoPago GO

Basic library for mercadopago (Forked from http://github.com/danteay/mercadopago)


## Install

```bash
go get github.com/facundomedica/mercadopago
```

## Usage

### Importing

```go
import mp "github.com/facundomedica/mercadopago"
```

### Configuring

#### ClientId and ClientSecret

```go
clientId := "12345678987"
clientSecret := "iugbkjSfFewASndfvñjn1234"

client := new(mp.MpClient)
client.Init(clientId, clientSecret)
```

#### AccessToken

```go
accessToken := "APP_PP_iugbkjSfFewASndfvñjn1234nlkefa__T_TT__adfhbrjnkfkl"

client := new(mp.MpClient)
client.Init(accessToken)
```

### Create order

```go
paymentData := mp.RequestData{
  "payer": mp.RequestData{
    "type":       "customer",
    "email":      "dante@testuser.com",
    "first_name": "Dante Aligeri",
    "last_name":  "",
  },
  "transaction_amount": 20,
  "description":        "Service expres - Service regular",
  "payment_method_id":  "oxxo",
}

response, err := client.CreatePayment(paymentData)

fmt.Println(err)
fmt.Println(string(response))
```