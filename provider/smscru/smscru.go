package smscru

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/intelfx/sachet"
	"github.com/intelfx/sachet/util"
)

type Config struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
}

type SmscRu struct {
	Config
}

func NewSmscRu(config Config) *SmscRu {
	SmscRu := &SmscRu{config}
	return SmscRu
}

func (c *SmscRu) Send(message sachet.Message) (resp sachet.Response, err error) {
	URL := "https://smsc.ru/sys/send.php"

	form := url.Values{
		"login":  {c.Login},
		"psw":    {c.Password},
		"phones": {strings.Join(message.To, ";")},
		"mes":    {message.Text},
		"fmt":    {"3"},
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

	_, resp_is_error := resp_obj.(map[string]interface{})["error"]
	if resp_is_error {
		err = fmt.Errorf("Backend reported error")
	}

	resp.Body = resp_obj
	return
}
