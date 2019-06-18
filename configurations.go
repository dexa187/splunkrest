package splunkrest

import (
	"encoding/json"
	"fmt"
	"net/url"
)

//Stanza in splunk config file
type Stanza struct {
	Name    string                 `json:"name"`
	ID      string                 `json:"id"`
	Content map[string]interface{} `json:"content"`
}

//Conf splunk configuration file
type Conf struct {
	Entry []Stanza `json:"entry"`
}

//GetConf gets configuration file from splunk
func (conn SplunkConnection) GetConf(app string, file string) (map[string]map[string]interface{}, error) {
	data := make(url.Values)
	data.Add("output_mode", "json")
	response, err := conn.HTTPGet(fmt.Sprintf("%s/servicesNS/%s/%s/configs/conf-%s?output_mode=json", conn.BaseURL, conn.Username, app, file), &data)
	var conf Conf
	err = json.Unmarshal([]byte(response), &conf)
	confMap := make(map[string]map[string]interface{})
	for _, entry := range conf.Entry {
		confMap[entry.Name] = entry.Content
	}
	return confMap, err
}
