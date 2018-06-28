package smscru

import (
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

func (c *SmscRu) Send(message sachet.Message) (err error) {
	URL := "https://smsc.ru/sys/send.php"

	form := url.Values{
		"login":  {c.Login},
		"psw":    {c.Password},
		"phones": {strings.Join(message.To, ";")},
		"mes":    {message.Text},
		"fmt":    {"3"},
	}

	query := url.Values{}

	err = util.SimpleSend(form, query, URL)
	if err != nil {
		return
	}
	return
}
