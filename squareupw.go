/*
Package squareupw implements a wrapper for squareup.com API for Go (Golang).
*/
package squareupw

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

const (
	BaseURL    = "https://connect.squareup.com"
	MethodGet  = "GET"
	MethodPost = "POST"
	MethodPut  = "PUT"
)

type API struct {
	token string
}

//Error represents Error response.
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

//NewAPI returns new API instance.
func NewAPI(token string) *API {
	return &API{
		token: token,
	}
}

//Send http request.
func (a API) Send(method, url string, reqData []byte) (httpResp *http.Response, body []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqData))
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer "+a.token)
	req.Header.Add("Accept", "application/json")
	if method == MethodPost || method == MethodPut {
		req.Header.Add("Content-Type", "application/json")
	}

	client := http.DefaultClient

	httpResp, err = client.Do(req)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	body, err = ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	if httpResp.StatusCode >= 400 {
		errorResp := Error{}
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return
		}
		err = errors.New(errorResp.Message)
	}

	return
}

//ExtractURLFromLinkHeader extracts url from Link header.
func ExtractURLFromLinkHeader(header []string) (url string, err error) {
	link := header[0]
	pattern := "^<(.+)>;rel='next'$"
	re := regexp.MustCompile(pattern)
	result := re.FindStringSubmatch(link)

	i := 1
	if len(result)-1 >= i {
		url := result[i]
		return url, nil
	}

	err = errors.New("Invalid value for Link header")
	return
}

//GetQueryStringByStruct take struct and build query string.
func GetQueryStringByStruct(s interface{}, tagName string, queryEscape bool) (queryString string, err error) {
	queryStringSlice := []string{}
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		tagVal := tag.Get(tagName)
		if len(tagVal) < 1 {
			continue
		}

		if valueField.Kind() != reflect.String {
			err = errors.New("Values for query string must have string type.")
			return
		}

		if len(valueField.String()) < 1 {
			continue
		}

		v := valueField.String()
		if queryEscape {
			v = url.QueryEscape(v)
		}

		queryStringSlice = append(queryStringSlice, tagVal+"="+v)
	}
	queryString = strings.Join(queryStringSlice, "&")
	return
}
