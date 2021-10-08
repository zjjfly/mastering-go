package ch2

import "testing"

func TestCallEmbedC(t *testing.T) {
	callEmbedC()
}

func TestCallSeparateC(t *testing.T) {
	callSeparateCCode()
}
