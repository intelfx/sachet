package smsru

import (
	"encoding/json"
	"fmt"
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

func (c *SmsRu) Send(message sachet.Message) (resp sachet.Response, err error) {
	URL := "https://sms.ru/sms/send"

	form := url.Values{
		"to":     {strings.Join(message.To, ",")},
		"msg":    {message.Text},
		"api_id": {c.APIId},
		"json":   {"1"},
	}

	query := url.Values{}

	resp, err = util.SimpleSend(form, query, URL)
	if err != nil {
		return
	}

	var resp_obj interface{}
	err = json.Unmarshal(resp.Body.([]byte), &resp_obj)
	if err != nil {
		return
	}

	resp_status := resp_obj.(map[string]interface{})["status"].(string)
	if resp_status != "OK" {
		err = fmt.Errorf("Backend reported error")
	}

	resp.Body = resp_obj
	return
}
