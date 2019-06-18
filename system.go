package splunkrest

import (
	"fmt"
)

// RestartServer restarts splunk
func (conn SplunkConnection) RestartServer() (string, error) {
	response, err := conn.HTTPPost(fmt.Sprintf("%s/services/server/control/restart", conn.BaseURL), nil)
	return response, err
}
