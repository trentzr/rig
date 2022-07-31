package logger_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trentzr/rig/logger"
)

func TestScreeningJSON(t *testing.T) {

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	log := logger.New(
		logger.WithFormat(logger.FormatJSON),
		logger.WithLevel(logger.LevelDebug),
		logger.WithOutputs(buf),
		logger.WithTimestampEnabled(false),
	)

	log.WithField("request_body", `{"str":"str"}`).Debug("test message")

	msgStruct := struct {
		Level       string `json:"level"`
		RequestBody string `json:"request_body"`
		Message     string `json:"message"`
	}{}

	err := json.Unmarshal(buf.Bytes(), &msgStruct)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, msgStruct.Level, "debug")
	assert.Equal(t, msgStruct.RequestBody, `{"str":"str"}`)
	assert.Equal(t, msgStruct.Message, "test message")

	if !strings.Contains(buf.String(), `{\"str\":\"str\"}`) {
		t.Fatal("screening for json is failed")
	}
}
