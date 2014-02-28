package config

import (
	"testing"
)

type SampleCfg struct {
	Name string `json:"name"`
}

type Entry struct {
	Field string `json:"field"`
}

type SampleMap struct {
	Records map[string]Entry `json:"map"`
}

func TestConfig(t *testing.T) {
	Init("./config_test.json")

	if Get().file != "./config_test.json" {
		t.Fatalf("Expected ./test.json, got: %v\n", Get().file)
	}

	v := &SampleCfg{}

	LoadInto(v)

	if v.Name != "xxx" {
		t.Fatalf("Expected name = xxx, got: %v", v.Name)
	}

	m := &SampleMap{}

	LoadInto(m)

	if m.Records["one"].Field != "whatever" || m.Records["two"].Field != "onemore" {
		t.Fatalf("Expected fields 'whatever' and 'onemore', got: %v\n", m.Records)
	}
}
