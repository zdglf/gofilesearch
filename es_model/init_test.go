package es_model

import "testing"

func TestInitEsClient(t *testing.T) {
	if _, err := InitEsClient(); err != nil {
		t.Error(err.Error())
	}
}
