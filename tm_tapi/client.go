package tm_tapi

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/xml"
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
		Code  int    `xml:"code"`
		Descr string `xml:"descr"`
		Data  any    `xml:"data"`
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

func (cl *Client) Get(reqName string, values url.Values, obj_resp any) error {
	url := "https://" + cl.addr + "/tm_tapi/1.0/" + reqName

	var request string
	if values != nil {
		request = values.Encode()
	}

	request += fmt.Sprintf("&signature=%x", md5.Sum(append([]byte(request)[:], cl.apiKey[:]...)))
	url += "?" + request

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := cl.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	var r = Response{Data: obj_resp}
	err = decoder.Decode(&r)
	if err != nil {
		return err
	}

	if r.Code != 0 {
		return errors.New(r.Descr)
	}

	return nil
}

func (cl *Client) Post(reqName string, values url.Values, obj_resp any) error {
	url := "https://" + cl.addr + "/tm_tapi/1.0/" + reqName

	body := []byte(values.Encode())

	signature := fmt.Sprintf("&signature=%x", md5.Sum(append(body[:], cl.apiKey[:]...)))
	body = append(body[:], signature[:]...)

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := cl.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	var r = Response{Data: obj_resp}
	err = decoder.Decode(&r)
	if err != nil {
		return err
	}

	if r.Code != 0 {
		return errors.New(r.Descr)
	}

	return nil
}
