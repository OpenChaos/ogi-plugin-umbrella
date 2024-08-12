package main

/*
* it is a consumer plug-in for Ogi

* a simple HTTP Service which consumes Body at '/consume' and passes it to configured Transformer

* exposes health check endpoint at '/ping' AND prometheus metrics at '/metrics' routes

* has Basic Auth for consume enabled by default, can be disabled by env var CONSUMER_API_BASICAUTH_ENABLED=false

* basic auth username and password can be configured using env var CONSUMER_API_BASICAUTH_USERNAME & CONSUMER_API_BASICAUTH_PASSWORD
 */

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gol-gol/golenv"
	ulid "github.com/oklog/ulid/v2"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/pseidemann/finish"

	"github.com/OpenChaos/ogi/logger"
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
	http.Handle("/metrics", promhttp.Handler())
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

	if basicAuthEnabled {
		logger.Infoln("Basic Auth is enabled. Could be configured with env CONSUMER_API_BASICAUTH_USERNAME/CONSUMER_API_BASICAUTH_PASSWORD.")
	} else {
		logger.Infoln("Basic Auth is disabled. Could be enabled with env CONSUMER_API_BASICAUTH_ENABLED")
	}
	logger.Infof("Listening at: %s\n", listenAt)
	logger.Infoln("a11y: GET /ping")
	logger.Infoln("o11y: GET /metrics")
	logger.Infoln("work: POST /consume")

	fin.Wait()
}

func BadRequest(w http.ResponseWriter, msg string) {
	errResponse := `{"status": "error", "error": "Bad Request"}`
	if msg != "" {
		errResponse = msg
	}
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, errResponse, http.StatusMethodNotAllowed)
}

func ping(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong!"))
	default:
		BadRequest(w, "")
	}
}

func consumeBody(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if basicAuthEnabled && !basicAuthHeaders(w, req) {
			BadRequest(w, "")
			return
		}
	default:
		BadRequest(w, "")
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Errorf("Error reading body: %v", err)
		BadRequest(w, "")
		return
	}

	msgid := ulid.Make().String()
	respB, err := ogitransformer.Transform(msgid, body)
	if err != nil {
		errStr := strings.ReplaceAll(err.Error(), `"`, "")
		BadRequest(w, `{"error": "`+errStr+`"}`)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(respB)
	}
}

func basicAuthHeaders(w http.ResponseWriter, req *http.Request) bool {
	username, password, ok := req.BasicAuth()
	if !ok {
		logger.Errorln("failed to get request basic auth details")
		return false
	} else if basicAuthUsername != username || basicAuthPassword != password {
		logger.Errorln("unauthorized request")
		return false
	}
	return true
}
