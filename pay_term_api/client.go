package pay_term_api

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	Client struct {
		c      *http.Client
		addr   string
		apiKey []byte
	}

	Response struct {
		Code  int         `json:"code"`
		Descr string      `json:"descr"`
		Data  interface{} `json:"data"`
	}

	EmptyResponse struct {
	}
)

func NewClient(addr, key string) *Client {
	return &Client{
		c: &http.Client{
			Timeout: 15 * time.Second,
			Transport: &http.Transport{
				IdleConnTimeout:   30 * time.Second,
				DisableKeepAlives: false,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				DisableCompression: true,
			},
		},
		addr:   addr,
		apiKey: []byte(key),
	}
}

func (cl *Client) Post(reqName string, values url.Values, obj_resp interface{}) error {
	url := "https://" + cl.addr + "/pay_term_api/1.0/" + reqName

	body := []byte(values.Encode())

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Signature", fmt.Sprintf("%x", md5.Sum(append(body[:], cl.apiKey[:]...))))

	resp, err := cl.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var r = Response{Data: obj_resp}
	err = decoder.Decode(&r)
	if err != nil {
		return err
	}

	if r.Code != 0 {
		return errors.New(fmt.Sprintf("Code: %d. Descr: %s", r.Code, r.Descr))
	}

	return nil
}
