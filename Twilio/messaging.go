package Twilio

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Connection(number string, message string) {
    accountSid := "ACb2a09b5cd18b66f3b635a8a84cf7e0d8"
    authToken := "f2b1e0a63c1128fa9e9a05f3775dca83"

    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: accountSid,
        Password: authToken,
    })

    params := &openapi.CreateMessageParams{}
    params.SetTo(number)
    params.SetFrom("+19706967906")
    params.SetBody(message)

    resp, err := client.Api.CreateMessage(params)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        response, _ := json.Marshal(*resp)
        log.Println("Response: " + string(response))
    }
}