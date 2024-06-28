package utils

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func SetEncoding(outputEncoding string, dataInBytes []byte) (data string) {
	switch strings.ToUpper(outputEncoding) {
	case "HEX":
		data = hex.EncodeToString(dataInBytes)
	case "BASE64":
		data = base64.StdEncoding.EncodeToString(dataInBytes)
	}
	return data
}

func SetDecoding(inputEncoding, data string) (dataInBytes []byte, err error) {
	switch strings.ToUpper(inputEncoding) {
	case "HEX":
		dataInBytes, err = hex.DecodeString(data)
		if err != nil {
			return nil, err
		}
	case "BASE64":
		dataInBytes, err = base64.StdEncoding.DecodeString(data)
		if err != nil {
			return nil, err
		}
	}
	return dataInBytes, err
}
