package main

import (
	"encoding/json"
	"io/ioutil"

	logger "github.com/OpenChaos/ogi/logger"

	"github.com/gol-gol/golenv"
	"github.com/gol-gol/golhttpclient"
)

/*
* to enable SSL Verify in HTTP Requests, set env var OGI_HTTP_PRODUCER_SKIP_SSL to "true"

// golhttpclient.HTTPRequest reference for Transformer creators
type HTTPRequest struct {
	Protocol    string
	Method      string
	Url         string
	GetParams   map[string]string
	HTTPHeaders map[string]string
	Body        *bytes.Buffer
}
*/

func Close() {
	return
}

func Produce(msgid string, msg []byte) {
	golhttpclient.SkipSSLVerify = golenv.OverrideIfEnvBool("OGI_HTTP_PRODUCER_SKIP_SSL", true)

	request := golhttpclient.HTTPRequest{}

	if err := json.Unmarshal(msg, &request); err != nil {
		logger.Errorf("Delivery failed for msgid[%s]' : %v\n", msgid, err)
	}

	response, err := request.Response()
	if err != nil {
		logger.Errorf("Delivery failed for msgid[%s] : %v\n", msgid, err)
	}

	resBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	logger.Infof("Delivery success for msgid[%s]\nStatus: %s\nBody: %v", msgid, response.Status, resBody)
	return
}
