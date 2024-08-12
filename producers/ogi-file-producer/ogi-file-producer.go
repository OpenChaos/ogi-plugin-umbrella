package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gol-gol/golenv"
)

type OgiFile struct {
	f *os.File
}

var (
	OgiFilepath = golenv.OverrideIfEnv("OGI_PRODUCER_FILEPATH", "/tmp/ogi-producer-file")

	fd *OgiFile
)

func init() {
	fd = &OgiFile{f: nil}
}

func (ogifile *OgiFile) Close() {
	if err := ogifile.f.Close(); err != nil {
		log.Fatal(err)
	}
	return
}

func (ogifile *OgiFile) Produce(msgid string, lyn []byte) ([]byte, error) {
	var err error
	ogifile.f, err = os.OpenFile(OgiFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintln(ogifile.f, string(lyn)); err != nil {
		return []byte{}, err
	}
	ogifile.f.Sync()
	return []byte{}, nil
}

func Close() {
	fd.Close()
}

func Produce(msgid string, lyn []byte) ([]byte, error) {
	return fd.Produce(msgid, lyn)
}
