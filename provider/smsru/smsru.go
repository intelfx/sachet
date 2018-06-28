package smsru

import (
	"net/url"
	"strings"

	"github.com/intelfx/sachet"
	"github.com/intelfx/sachet/util"
)

type Config struct {
	APIId string `yaml:"api_id"`
}

type SmsRu struct {
	Config
}

func NewSmsRu(config Config) *SmsRu {
	SmsRu := &SmsRu{config}
	return SmsRu
}

func (c *SmsRu) Send(message sachet.Message) (err error) {
	URL := "https://sms.ru/sms/send"

	form := url.Values{
		"to":     {strings.Join(message.To, ",")},
		"msg":    {message.Text},
		"api_id": {c.APIId},
		"json":   {"1"},
	}

	query := url.Values{}

	err = util.SimpleSend(form, query, URL)
	if err != nil {
		return
	}
	return
}
