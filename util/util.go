package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/intelfx/sachet"
)

const RequestTimeout = time.Second * 10

func SimpleSend(form url.Values, query url.Values, URL string) (sachet_resp sachet.Response, err error) {
	var req *http.Request
	var resp *http.Response

	headers := http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"User-Agent":   {"Sachet v1.0"},
	}

	req, err = http.NewRequest("POST", URL, strings.NewReader(form.Encode()))
	if err != nil {
		return
	}
	req.URL.RawQuery = query.Encode()
	for k, vs := range headers {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	client := &http.Client{}
	client.Timeout = RequestTimeout
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	sachet_resp = sachet.Response{
		Status: resp.StatusCode,
		Body:   body,
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Failed to send SMS: HTTP %d \"%s\"", resp.StatusCode, string(body))
	}
	return
}
