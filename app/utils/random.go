package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateTraceID() string {
	timeNow := time.Now().Format("20060102150405.000")
	traceId := strings.Replace(timeNow, ".", "", 1) + CreateRandomHex(5)
	return traceId
}

func CreateRandomHex(n int) string {
	const letterBytes = "ABCDEF0123456789"
	const (
		letterIdxBits = 4
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func CreateCustomUUID(data string) (customUUID string) {
	customUUID = strings.ReplaceAll(uuid.NewString(), "-", "")
	customUUID = data + customUUID
	return strings.ToUpper(customUUID)
}
