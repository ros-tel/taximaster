package common_api

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type (
	Client struct {
		c       *http.Client
		addr    string
		apiKey  []byte
		user_id *int
	}

	Response struct {
		Code  int         `json:"code"`
		Descr string      `json:"descr"`
		Data  interface{} `json:"data"`
	}

	EmptyResponse struct {
	}
)

func NewClient(addr, key string, id *int) *Client {
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
		addr:    addr,
		apiKey:  []byte(key),
		user_id: id,
	}
}

func (cl *Client) invoke(e errorMap, req *http.Request, obj_resp interface{}) error {
	if cl.user_id != nil {
		req.Header.Add("X-User-Id", strconv.Itoa(*cl.user_id))
	}

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
		return errorByCode(e, r.Code, r.Descr)
	}

	return nil
}

func (cl *Client) Get(reqName string, e errorMap, values url.Values, obj_resp interface{}) error {
	url := "https://" + cl.addr + "/common_api/1.0/" + reqName
	var request string
	if values != nil {
		request = values.Encode()
	}
	if request != "" {
		url += "?" + request
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Signature", fmt.Sprintf("%x", md5.Sum(append([]byte(request)[:], cl.apiKey[:]...))))

	return cl.invoke(e, req, obj_resp)
}

func (cl *Client) Post(reqName string, e errorMap, values url.Values, obj_resp interface{}) error {
	url := "https://" + cl.addr + "/common_api/1.0/" + reqName

	body := []byte(values.Encode())

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Signature", fmt.Sprintf("%x", md5.Sum(append(body[:], cl.apiKey[:]...))))

	return cl.invoke(e, req, obj_resp)
}

func (cl *Client) PostJson(reqName string, e errorMap, obj_req, obj_resp interface{}) error {
	url := "https://" + cl.addr + "/common_api/1.0/" + reqName

	body, err := json.Marshal(obj_req)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Signature", fmt.Sprintf("%x", md5.Sum(append(body[:], cl.apiKey[:]...))))

	return cl.invoke(e, req, obj_resp)
}

func errorByCode(e errorMap, code int, descr string) error {
	var (
		ok  bool
		err error
	)

	switch code {
	case 1:
		return ErrUnknownError
	case 2:
		return ErrUnknownApiType
	case 3:
		return ErrApiDisabledInSettings
	case 4:
		return ErrSecretKeyDoesNotMatch
	case 5:
		return ErrUnsupportedApiVersion
	case 6:
		return ErrUnknownRequestName
	case 7:
		return ErrInvalidRequestType
	case 8:
		return ErrMissingParameter
	case 9:
		return ErrIncorrectParameter
	case 10:
		return ErrInternalRequestProcessing
	case 13:
		return ErrUserCommonAPINotFound
	case 14:
		return ErrRequestNotAvailableToCommonAPIUser
	}

	err, ok = e[code]
	if ok {
		return err
	}

	return fmt.Errorf("common_api: unknown Code: %d. Descr: %s", code, descr)
}
