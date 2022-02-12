package xanhttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type HttpInterface interface {
	PostRawData(m PostRawDataCaller) (result []byte, err error)
	PostFormData(m PostFormDataCaller) (result []byte, err error)
}

const (
	defaultTimeout = time.Second * 80
)

func defaultTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

type CustomHTTP struct{}

type CallerOptions struct {
	Url       string
	Method    string
	Data      []byte
	Headers   map[string]string
	TLSConfig *tls.Config
}

func AddXanHTTPInterface() CustomHTTP {
	return CustomHTTP{}
}

type globalCaller struct {
	URL     string
	Query   map[string]string
	Body    interface{}
	Header  map[string]string
	Options struct {
		Timeout         time.Duration
		TLSClientConfig *tls.Config
	}
}

type PostRawDataCaller struct {
	globalCaller
}

func fillingURLQuery(requestURL *url.URL, mapQuery map[string]string) {
	if mapQuery != nil {
		query := requestURL.Query()
		for key, value := range mapQuery {
			query.Add(key, value)
		}
		requestURL.RawQuery = query.Encode()
	}
}

func fillingHeader(request *http.Request, mapHeader map[string]string) {
	if mapHeader != nil {
		for key, value := range mapHeader {
			request.Header.Add(key, value)
		}
	}
}

func (opt CustomHTTP) PostRawData(m PostRawDataCaller) (result []byte, err error) {
	// evaluate the options
	if m.Options.TLSClientConfig == nil {
		m.Options.TLSClientConfig = defaultTLSConfig()
	}

	if m.Options.Timeout <= 0 {
		m.Options.Timeout = defaultTimeout
	}

	// declare the variable
	var requestBody *bytes.Buffer
	var requestURL *url.URL
	var responses *http.Response
	var rbj []byte // shortness as request body json

	// model validation
	if m.URL == "" {
		return nil, fmt.Errorf("url is required")
	}

	requestURL, err = url.Parse(m.URL)
	if err != nil {
		return nil, fmt.Errorf("url is invalid: %s", err.Error())
	}

	// filling url query
	fillingURLQuery(requestURL, m.Query)
	if m.Body != nil {
		rbj, err = json.Marshal(m.Body)
		if err != nil {
			return nil, fmt.Errorf("json marshal error: %s", err.Error())
		}

		requestBody = bytes.NewBuffer(rbj)
	}

	client := &http.Client{
		Timeout: m.Options.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: m.Options.TLSClientConfig,
		},
	}
	request, err := http.NewRequest(http.MethodPost, requestURL.String(), requestBody)
	if err != nil {
		return nil, fmt.Errorf("http new request error: %s", err.Error())
	}

	// add http request header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", strconv.Itoa(len(requestBody.Bytes())))
	request.Header.Set("User-Agent", "golang-http-client")
	request.Header.Set("Cache-Control", "no-cache")
	fillingHeader(request, m.Header)
	responses, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	defer responses.Body.Close()

	// translate
	result, err = ioutil.ReadAll(responses.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %s", err.Error())
	}

	// add your validation responses code here
	if responses.StatusCode > 299 {
		return result, fmt.Errorf("server request rejected because: %s", responses.Status)
	}

	return result, nil
}

type PostFormDataCaller struct {
	globalCaller
}

func (opt CustomHTTP) PostFormData(m PostFormDataCaller) (result []byte, err error) {
	// evaluate the options
	if m.Options.TLSClientConfig == nil {
		m.Options.TLSClientConfig = defaultTLSConfig()
	}

	if m.Options.Timeout <= 0 {
		m.Options.Timeout = defaultTimeout
	}

	// variable declaration
	var payload *bytes.Buffer
	var param url.Values
	var responses *http.Response
	var request *http.Request

	// parsing and checking url data
	taggedURL, err := url.Parse(m.URL)
	if err != nil {
		err = fmt.Errorf("url parsing error: %s", err.Error())
		return nil, err
	}

	fillingURLQuery(taggedURL, m.Query)
	if m.Body != nil { // prevent null pointer
		drj, _ := json.Marshal(m.Body)
		var dr map[string]interface{}
		err = json.Unmarshal(drj, &dr)
		if err != nil {
			return nil, err
		}
		for k, v := range dr {
			param.Add(k, fmt.Sprintf("%s", v))
		}
		payload = bytes.NewBufferString(param.Encode())
	}

	client := &http.Client{
		Timeout: m.Options.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: m.Options.TLSClientConfig,
		},
	}
	request, err = http.NewRequest(http.MethodPost, taggedURL.String(), payload)
	if err != nil {
		return nil, err
	}

	// add required http request header
	request.Header.Set("Content-Type", "multipart/form-data")
	request.Header.Set("Content-Length", strconv.Itoa(len(payload.Bytes())))
	request.Header.Set("User-Agent", "golang-http-client")
	request.Header.Set("Cache-Control", "no-cache")
	fillingHeader(request, m.Header)
	responses, err = client.Do(request)
	if err != nil {
		return nil, err
	}

	// converting
	result, err = ioutil.ReadAll(responses.Body)
	if err != nil {
		err = fmt.Errorf("response body read error: %s", err.Error())
		return nil, err
	}

	// add your validation responses code here
	if responses.StatusCode > 299 {
		return result, fmt.Errorf("server request rejected because: %s", responses.Status)
	}

	return result, nil
}
