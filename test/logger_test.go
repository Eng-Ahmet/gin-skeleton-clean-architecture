package test

import (
	"bytes"
	"hwai-api/pkg/logger"
	"testing"
)

func TestGetInfoLogger(t *testing.T) {
	// Simulating writing to memory instead of an actual file
	var buf bytes.Buffer
	log := logger.GetInfoLogger()
	log.SetOutput(&buf)

	// Write a message to log
	log.Println("Test info message")

	// Check if the message is present in the buffer
	if !bytes.Contains(buf.Bytes(), []byte("Test info message")) {
		t.Errorf("Expected 'Test info message' to be logged, but got: %s", buf.String())
	}
}

func TestGetWarningLogger(t *testing.T) {
	// Simulating writing to memory instead of an actual file
	var buf bytes.Buffer
	log := logger.GetWarningLogger()
	log.SetOutput(&buf)

	// Write a message to log
	log.Println("Test warning message")

	// Check if the message is present in the buffer
	if !bytes.Contains(buf.Bytes(), []byte("Test warning message")) {
		t.Errorf("Expected 'Test warning message' to be logged, but got: %s", buf.String())
	}
}

func TestGetErrorLogger(t *testing.T) {
	// Simulating writing to memory instead of an actual file
	var buf bytes.Buffer
	log := logger.GetErrorLogger()
	log.SetOutput(&buf)

	// Write a message to log
	log.Println("Test error message")

	// Check if the message is present in the buffer
	if !bytes.Contains(buf.Bytes(), []byte("Test error message")) {
		t.Errorf("Expected 'Test error message' to be logged, but got: %s", buf.String())
	}
}
