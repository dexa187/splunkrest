package splunkrest

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// SplunkConnection connection to splunk
type SplunkConnection struct {
	Username, Password, BaseURL, SessionKey string
}

// SessionKey represents the JSON object returned from the Splunk authentication REST call

// Login connects to the Splunk server and retrieves a session key
func (conn SplunkConnection) Login() error {

	data := make(url.Values)
	data.Add("username", conn.Username)
	data.Add("password", conn.Password)
	data.Add("output_mode", "json")
	response, err := conn.HTTPPost(fmt.Sprintf("%s/services/auth/login", conn.BaseURL), &data)

	if err != nil {
		return err
	}

	type SessionKey struct {
		Value string `json:"sessionKey"`
	}

	bytes := []byte(response)
	var key SessionKey
	err = json.Unmarshal(bytes, &key)
	conn.SessionKey = key.Value
	return err
}
