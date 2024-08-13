package main

import (
	"encoding/json"
	"errors"

	"github.com/OpenChaos/ogi/logger"
)

type YFinResponse struct {
	Chart struct {
		Results []struct {
			Meta struct {
				Symbol          string `json:"symbol"`
				LongName        string `json:"longName"`
				Range           string `json:"range"`
				DataGranularity string `json:"dataGranularity"`
			} `json:"meta"`
			Indicators struct {
				Quote []struct {
					Close []float32 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

type Response struct {
	Symbol          string
	LastClose       float32
	FiftyTwoWeekLow float32
	At52Low         bool
}

func Close() {
	return
}

func Produce(msgid string, msg []byte) ([]byte, error) {
	var yfinData YFinResponse
	if decErr := json.Unmarshal(msg, &yfinData); decErr != nil {
		return []byte{}, decErr
	}
	if yfinData.Chart.Results[0].Meta.Range != "1y" {
		return []byte{}, errors.New("YFinance API call should use range=1y")
	} else if yfinData.Chart.Results[0].Meta.DataGranularity != "1d" {
		return []byte{}, errors.New("YFinance API call should use interval=1d")
	}

	response := Response{Symbol: yfinData.Chart.Results[0].Meta.Symbol}
	stock := yfinData.Chart.Results[0].Meta.LongName
	total_close := len(yfinData.Chart.Results[0].Indicators.Quote[0].Close)
	response.LastClose = yfinData.Chart.Results[0].Indicators.Quote[0].Close[total_close-1]
	response.FiftyTwoWeekLow = response.LastClose
	for _, day_close := range yfinData.Chart.Results[0].Indicators.Quote[0].Close {
		if day_close < response.FiftyTwoWeekLow {
			response.FiftyTwoWeekLow = day_close
		}
	}
	logger.Infof("%s Closed at: %f (52 week low: %f)",
		response.Symbol, response.LastClose, response.FiftyTwoWeekLow)
	if response.FiftyTwoWeekLow >= response.LastClose {
		response.At52Low = true
		logger.Infof("msgid[%s]: %s at 52WeekLow.", msgid, stock)
	} else {
		logger.Infof("msgid[%s]: %s not at 52WeekLow.", msgid, stock)
	}
	return json.Marshal(&response)
}
