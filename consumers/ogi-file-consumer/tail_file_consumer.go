package main

/*
* it's a consumer plug-in for Ogi

* a tail-file that consumes new lines added to a watched file & passes them onwards

 */

import (
	"bufio"
	"os"

	"github.com/gol-gol/golenv"
	ulid "github.com/oklog/ulid/v2"

	"github.com/OpenChaos/ogi/logger"
	ogitransformer "github.com/OpenChaos/ogi/transformer"
)

var (
	FileToConsume = golenv.OverrideIfEnv("OGI_FILE_TO_CONSUME", "/tmp/ogi-consumed")
)

func Consume() {
	logger.Infof("tail-file consumption at: %s\n", FileToConsume)
	consumeFile()
}

func consumeFile() {
	fileHandle, err := os.Open(FileToConsume)
	if err != nil {
		logger.Fatalln(err)
	}
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		transform(fileScanner.Text())
	}
}

func transform(lyne string) {
	msgid := ulid.Make().String()
	respB, err := ogitransformer.Transform(msgid, []byte(lyne))
	if err != nil {
		logger.Infoln("[ERROR]", err)
	} else if len(respB) > 0 {
		logger.Infoln("[RESPONSE]", string(respB))
	}
}
