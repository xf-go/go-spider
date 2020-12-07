package main

import (
	"bytes"
	"os"
)

func writeFileLine(filename string, date []byte) error {
	var buf bytes.Buffer
	buf.Write(date)
	buf.Write([]byte("\r\n"))
	err := writeFile(filename, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func writeFile(filename string, date []byte) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(date)
	if err != nil {
		return err
	}

	return nil
}
