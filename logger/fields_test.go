package logger_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/trentzr/rig/logger"

	"github.com/stretchr/testify/assert"
)

func TestField(t *testing.T) {

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	log := logger.New(
		logger.WithFormat(logger.FormatJSON),
		logger.WithLevel(logger.LevelDebug),
		logger.WithOutputs(buf),
		logger.WithTimestampEnabled(false),
	)

	log.WithField("string", "Abc").
		WithField("int", 123).
		WithField("bool", true).
		Debug("test message")

	msgStruct := struct {
		Level   string `json:"level"`
		Str     string `json:"string"`
		Int     int    `json:"int"`
		Bool    bool   `json:"bool"`
		Message string `json:"message"`
	}{}

	err := json.Unmarshal(buf.Bytes(), &msgStruct)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, msgStruct.Level, "debug")
	assert.Equal(t, msgStruct.Str, "Abc")
	assert.Equal(t, msgStruct.Int, 123)
	assert.Equal(t, msgStruct.Bool, true)
	assert.Equal(t, msgStruct.Message, "test message")
}

func TestFields(t *testing.T) {

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	log := logger.New(
		logger.WithFormat(logger.FormatJSON),
		logger.WithLevel(logger.LevelDebug),
		logger.WithOutputs(buf),
		logger.WithTimestampEnabled(false),
	)

	log.WithFields(logger.Fields{
		"string": "Abc",
		"int":    123,
		"bool":   true,
	}).Debug("test message")

	msgStruct := struct {
		Level   string `json:"level"`
		Str     string `json:"string"`
		Int     int    `json:"int"`
		Bool    bool   `json:"bool"`
		Message string `json:"message"`
	}{}

	err := json.Unmarshal(buf.Bytes(), &msgStruct)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, msgStruct.Level, "debug")
	assert.Equal(t, msgStruct.Str, "Abc")
	assert.Equal(t, msgStruct.Int, 123)
	assert.Equal(t, msgStruct.Bool, true)
	assert.Equal(t, msgStruct.Message, "test message")
}

func TestErrorField(t *testing.T) {

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	log := logger.New(
		logger.WithFormat(logger.FormatJSON),
		logger.WithLevel(logger.LevelDebug),
		logger.WithOutputs(buf),
		logger.WithTimestampEnabled(false),
	)

	log.WithError(errors.New("test error")).Debug("test message")

	msgStruct := struct {
		Level   string `json:"level"`
		Error   string `json:"error"`
		Message string `json:"message"`
	}{}

	err := json.Unmarshal(buf.Bytes(), &msgStruct)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, msgStruct.Level, "debug")
	assert.Equal(t, msgStruct.Error, "test error")
	assert.Equal(t, msgStruct.Message, "test message")
}
