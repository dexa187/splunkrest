package splunkrest

import (
	"fmt"
	"net/url"
)

//InstallApp installs app from file path
func (conn SplunkConnection) InstallApp(path string, update bool) (string, error) {
	data := make(url.Values)
	data.Add("name", path)

	updateApp := "false"
	if update == true {
		updateApp = "true"
	}

	data.Add("update", updateApp)
	response, err := conn.HTTPPost(fmt.Sprintf("%s/services/apps/appinstall/", conn.BaseURL), &data)
	return response, err
}
