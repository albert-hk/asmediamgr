package parser

import (
	"asmediamgr/pkg/dirinfo"
	"testing"
)

type MockParser struct{}

func (p *MockParser) IsDefaultEnable() bool {
	return true
}

func (p *MockParser) Init(cfgPath string) (priority float32, err error) {
	return 0, nil
}

func (p *MockParser) ParseV2(entry *dirinfo.Entry) (ok bool, err error) {
	return true, nil
}

func TestNewParserMgr(t *testing.T) {
	RegisterParser("parser1", &MockParser{})
	RegisterParser("parser2", &MockParser{})
	RegisterParser("parser3", &MockParser{})
	opts := &ParserMgrOpts{
		Logger:         nil,
		ConfigDir:      "config",
		EnableParsers:  []string{"parser1", "parser2"},
		DisableParsers: []string{"parser3", "parser2"},
	}
	pm, err := NewParserMgr(opts)
	if err != nil {
		t.Fatalf("NewParserMgr() error = %v", err)
	}
	if pm == nil {
		t.Fatalf("NewParserMgr() pm is nil")
	}
	if len(pm.parsers) != 1 {
		t.Fatalf("NewParserMgr() pm.parsers length = %v", len(pm.parsers))
	}
	if pm.parsers[0].parser == nil {
		t.Fatalf("NewParserMgr() pm.parsers[0] is nil")
	}
}
