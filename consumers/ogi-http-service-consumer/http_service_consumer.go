package main

/*
* it is a consumer plug-in for Ogi

* a simple HTTP Service which consumes Body at '/consume' and passes it to configured Transformer

* exposes health check endpoint at '/ping' AND prometheus metrics at '/metrics' routes

* has Basic Auth for consume enabled by default, can be disabled by env var CONSUMER_API_BASICAUTH_ENABLED=false

* basic auth username and password can be configured using env var CONSUMER_API_BASICAUTH_USERNAME & CONSUMER_API_BASICAUTH_PASSWORD
 */

import (
	"fmt"
	"io/ioutil"
	"net/http"

	golenv "github.com/abhishekkr/gol/golenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pseidemann/finish"

	logger "github.com/OpenChaos/ogi/logger"
	ogitransformer "github.com/OpenChaos/ogi/transformer"
)

var (
	listenAt          = golenv.OverrideIfEnv("CONSUMER_API_LISTENAT", ":8080")
	basicAuthEnabled  = golenv.OverrideIfEnvBool("CONSUMER_API_BASICAUTH_ENABLED", true)
	basicAuthUsername = golenv.OverrideIfEnv("CONSUMER_API_BASICAUTH_USERNAME", "changeit")
	basicAuthPassword = golenv.OverrideIfEnv("CONSUMER_API_BASICAUTH_PASSWORD", "changeit")
)

/*
Consume function will be called by Ogi's plugin manager when this plugin gets used.
*/
func Consume() {
	logger.Infof("listening at: %s\n", listenAt)

	http.Handle("^/metrics$", promhttp.Handler())
	http.HandleFunc("/ping", ping)

	http.HandleFunc("/consume", consumeBody)

	svr := &http.Server{Addr: listenAt}
	fin := finish.New()
	fin.Add(svr)
	go func() {
		err := svr.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()
	fin.Wait()
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `pong!\n`)
}

func consumeBody(w http.ResponseWriter, req *http.Request) {
	if basicAuthEnabled {
		if !basicAuthHeaders(w, req) {
			return
		}
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	} else {
		ogitransformer.Transform(body)
	}
	fmt.Fprintf(w, `{"status": "queued", "message": "request body has been passed to transformer"}`)
}

func basicAuthHeaders(w http.ResponseWriter, req *http.Request) bool {
	username, password, ok := req.BasicAuth()
	if !ok {
		logger.Errorln("failed to get request basic auth details")
		http.Error(w, "can't read basic auth details", http.StatusBadRequest)
		return false
	} else if basicAuthUsername != username || basicAuthPassword != password {
		logger.Errorln("unauthorized request")
		http.Error(w, "can't read valid basic auth details", http.StatusForbidden)
		return false
	}
	return true
}
