package notification

import (
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioConfig struct {
	AccountSid  string
	Auth        string
	AlertPhone  string
	PersPhone   string
	TwilioPhone string
	AlertMsg    string
}

type TwilioService struct {
	client *twilio.RestClient
	params *twilioApi.CreateMessageParams
	msg    string
}

func NewTwilioConfig() *TwilioConfig {
	return &TwilioConfig{
		AccountSid:  os.Getenv("ACCOUNT_SID"),
		Auth:        os.Getenv("AUTH_TOKEN"),
		AlertPhone:  os.Getenv("ALERT_PHONE_NUMBER"),
		PersPhone:   os.Getenv("PERSONAL_PHONE"),
		TwilioPhone: os.Getenv("TWILIO_PHONE"),
		AlertMsg:    os.Getenv("ALERT_MSG"),
	}
}

func NewTwilioService(cfg *TwilioConfig) *TwilioService {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.AccountSid,
		Password: cfg.Auth,
	})

	phoneTo := cfg.AlertPhone
	log.Println("PHONE NUM TO: ", phoneTo)
	if phoneTo == "0" {
		phoneTo = cfg.PersPhone
	log.Println("PHONE NUM TO 2: ", phoneTo)

	}

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneTo)
	params.SetFrom(cfg.TwilioPhone)
	msg := cfg.AlertMsg

	return &TwilioService{
		client: client,
		params: params,
		msg:    msg,
	}
}

func (ts *TwilioService) Send(cityName string) error {

	ts.params.SetBody(fmt.Sprintf("%v %v", ts.msg, cityName))
	_, err := ts.client.Api.CreateMessage(ts.params)
	if err != nil {
		return fmt.Errorf("Error sending SMS message: " + err.Error())
	}

	log.Printf("ALERT SENT! Bad weather in %v!", cityName)
	return nil
}
