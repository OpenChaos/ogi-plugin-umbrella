package main

import (
	"fmt"
	"io"

	"github.com/gol-gol/golenv"
	"github.com/gol-gol/golhttpclient"

	logger "github.com/OpenChaos/ogi/logger"
	ogiproducer "github.com/OpenChaos/ogi/producer"
)

var (
	SkipSSLVerify   = golenv.OverrideIfEnvBool("OGI_HTTP_REQUEST_SKIP_SSL", true)
	YFinAPIRange    = golenv.OverrideIfEnv("OGI_YFIN_RANGE", "max")
	YFinAPIInterval = golenv.OverrideIfEnv("OGI_YFIN_INTERVAL", "1d")
)

func Transform(msgid string, msg []byte) ([]byte, error) {
	stock := string(msg)
	request := golhttpclient.Request{
		Url: fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", stock),
		Params: map[string]string{
			"interval": YFinAPIInterval,
			"range":    YFinAPIRange,
		},
		Method:        "GET",
		SkipSSLVerify: SkipSSLVerify,
	}

	response, err := request.Fetch()
	if err != nil || response.StatusCode >= 400 {
		logger.Errorf("Delivery failed for msgid[%s] : %v\n", msgid, err)
		return []byte{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	response.Body.Close()
	logger.Infof("YFinance fetch success for msgid[%s]: %s", msgid, stock)
	return ogiproducer.Produce(msgid, resBody)
}
