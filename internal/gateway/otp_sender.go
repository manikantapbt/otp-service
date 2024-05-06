package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// OTPSender is an interface for sending OTPs
type OTPSender interface {
	SendOTP(phone string, otp uint32) error
}

func NewTwilioOTPSender(client *twilio.RestClient, from string, accountSid string, authToken string) *TwilioClient {
	return &TwilioClient{
		client:     client,
		from:       from,
		accountSid: accountSid,
		authToken:  authToken,
	}
}

type TwilioClient struct {
	client     *twilio.RestClient
	from       string
	accountSid string
	authToken  string
}

// SendOTP sends an OTP to the given phone number by invoking twilio api
func (t *TwilioClient) SendOTP(phone string, otp uint32) error {
	accountSid := "twilio_account_sid"
	authToken := "twilio_auth_token"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(t.from)
	params.SetBody(fmt.Sprintf("Your registration otp is : %d", otp))

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return errors.New("unable to send sms, please try after some time")
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}
