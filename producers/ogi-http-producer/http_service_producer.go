package main

import (
	"encoding/json"
	"io"

	logger "github.com/OpenChaos/ogi/logger"

	"github.com/gol-gol/golenv"
	"github.com/gol-gol/golhttpclient"
)

/*
* to enable SSL Verify in HTTP Requests, set env var OGI_HTTP_PRODUCER_SKIP_SSL to "true"

// golhttpclient.Request reference for Transformer creators
type Request struct {
	Protocol    string
	Method      string
	Url         string
	Path        string
	Params      map[string]string
	Headers		map[string]string
	Body        *bytes.Buffer
	SkipSSLVerify	bool
}
*/

var (
	SkipSSLVerify = golenv.OverrideIfEnvBool("OGI_HTTP_PRODUCER_SKIP_SSL", true)
)

func Close() {
	return
}

func Produce(msgid string, msg []byte) ([]byte, error) {
	request := golhttpclient.Request{
		SkipSSLVerify: SkipSSLVerify,
	}

	if err := json.Unmarshal(msg, &request); err != nil {
		logger.Errorf("Delivery failed for msgid[%s]' : %v\n", msgid, err)
	}

	response, err := request.Fetch()
	if err != nil {
		logger.Errorf("Delivery failed for msgid[%s] : %v\n", msgid, err)
	}

	resBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	logger.Infof("Delivery success for msgid[%s]; HTTP %s", msgid, response.Status)
	return resBody, err
}
