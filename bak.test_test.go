package main

import (
	"strings"
	"testing"
)

func TestMyCode(t *testing.T) {
	body := []byte("{ \"nextEPSDate\":\"2024-01-01\" }")
	message := formatResp(body, "AAPL")

	control_message := "AAPL\t --> \t2024-01-01\n"

	if strings.Compare(message, control_message) != 0 {
		t.Log("Expected:\t AAPL\t --> \t2024-01-01")
		t.Logf("Got:\t %v", message)
		t.Fatalf("incorrect formatting on formatResp func")
	}
}
