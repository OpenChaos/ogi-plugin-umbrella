package main

/*
* it's a consumer plug-in for Ogi

* a TCP based service availability checker, checks loaded from OGI_TCP_CHECK_YAML
* sampl yaml-cfg
---
checks:
  - localhost:8080
  - example.com:443
interval_seconds: 15
backoff: 1
max_backoff: 20

*/

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/gol-gol/golenv"
	ulid "github.com/oklog/ulid/v2"
	"gopkg.in/yaml.v3"

	"github.com/OpenChaos/ogi/logger"
	ogitransformer "github.com/OpenChaos/ogi/transformer"
)

type TcpCheckCfg struct {
	Checks          []string `yaml:"checks"`
	IntervalSeconds uint32   `yaml:"interval_seconds"`
	BackoffTicker   uint32   `yaml:"backoff"`
	MaximumBackoff  uint32   `yaml:"max_backoff"`
}

type Notif struct {
	Context string
	Message string
}

var (
	TcpCheckYAML = golenv.OverrideIfEnv("OGI_TCP_CHECK_YAML", "/tmp/ogi-tcp-check.yaml")
)

func Consume() {
	logger.Infof("tcp-check consumption as per: %s\n", TcpCheckYAML)

	var err error
	var check_cfg TcpCheckCfg
	var yamlB []byte

	if yamlB, err = os.ReadFile(TcpCheckYAML); err != nil {
		logger.Fatalf("tcp-check cfg read failed: %s\n", err.Error())
	}
	if err = yaml.Unmarshal(yamlB, &check_cfg); err != nil {
		logger.Fatalf("tcp-check cfg parse failed: %s\n", err.Error())
	}

	backoffTicker := check_cfg.BackoffTicker
	for {
		for _, conn := range check_cfg.Checks {
			if tcpOK(conn) {
				backoffTicker = check_cfg.BackoffTicker
				continue
			}
			notif := Notif{
				Context: conn,
				Message: fmt.Sprintf("TCP Connect call failed: %s", conn),
			}
			msgB, errJson := json.Marshal(&notif)
			if errJson != nil {
				logger.Errorf("JSON bytes failed for tcp-check failure: %s\n", err.Error())
			}
			transform(msgB)
			backoffTicker = backoff(backoffTicker, check_cfg.MaximumBackoff)
		}
		waitForSeconds := time.Duration(check_cfg.IntervalSeconds * backoffTicker)
		time.Sleep(time.Second * waitForSeconds)
	}
}

func tcpOK(address string) bool {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		logger.Infoln("[error] TCP check failed for resolution")
		return false
	}
	conn, errConn := net.DialTCP("tcp", nil, tcpAddr)
	if errConn != nil {
		logger.Infoln("[error] TCP check failed for connection")
		return false
	}
	conn.Close()
	return true
}

func transform(msgB []byte) {
	msgid := ulid.Make().String()
	respB, err := ogitransformer.Transform(msgid, msgB)
	if err != nil {
		logger.Infoln("[ERROR]", err)
	} else if len(respB) > 0 {
		logger.Infoln("[RESPONSE]", string(respB))
	}
}

func backoff(current, maximumBackoff uint32) uint32 {
	current += 1
	if current <= maximumBackoff {
		return current
	}
	return maximumBackoff
}
