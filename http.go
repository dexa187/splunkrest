package splunkrest

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
 * HTTP helper methods
 */

func httpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}

//HTTPGet Makes an GET http call to splunk
func (conn SplunkConnection) HTTPGet(url string, data *url.Values) (string, error) {
	return conn.httpCall(url, "GET", data)
}

//HTTPPost Makes an POST http call to splunk
func (conn SplunkConnection) HTTPPost(url string, data *url.Values) (string, error) {
	return conn.httpCall(url, "POST", data)
}

func (conn SplunkConnection) httpCall(url string, method string, data *url.Values) (string, error) {
	client := httpClient()

	var payload io.Reader
	if data != nil {
		payload = bytes.NewBufferString(data.Encode())
	}

	request, err := http.NewRequest(method, url, payload)
	conn.addAuthHeader(request)
	if data != nil {
		request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	}
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return string(body), nil
}

func (conn SplunkConnection) addAuthHeader(request *http.Request) {
	if conn.SessionKey != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Splunk %s", conn.SessionKey))
	} else {
		request.SetBasicAuth(conn.Username, conn.Password)
	}
}
