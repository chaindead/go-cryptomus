# Golang Wrapper for cryptomus payment API

## Install

```bash
go get github.com/chaindead/go-cryptomus
```

## Check signature

```go
package main

import (
	"bytes"
	"net/http"
	
	gocryptomus "github.com/chaindead/go-cryptomus"
)

func main(){
	httpClient := http.Client{}
	merchant := "replace with your merchant id"
	paymentAPIKey := "replace with your payment API key"
	payoutAPIKey := "replace with your payout API key"
	client := gocryptomus.New(&httpClient, merchant, paymentAPIKey, payoutAPIKey)

	// c.Request().Body - handler logic
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request().Body)
	body := buf.Bytes()

	isSigValid := client.VerifySignature(body)
	if !isSigValid{
		return
    }
}
```

