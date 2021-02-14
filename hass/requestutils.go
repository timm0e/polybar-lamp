package hass

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func makeBodyBuffer(data interface{}) *bytes.Buffer {
	body, err2 := json.Marshal(data)
	if err2 != nil {
		panic(err2)
	}
	bodybuffer := bytes.NewBuffer(body)
	return bodybuffer
}


func addBearerTokenHeader(request *http.Request, apiKey string) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
}
