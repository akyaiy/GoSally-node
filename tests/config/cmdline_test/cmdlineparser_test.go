package cmdline_test

import (
	"GoSally/internal/config"
	"GoSally/internal/config/parsers"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestParseArgs(t *testing.T) {
	args := []string{"/path/to/prog", "--listen=0.0.0.0"}
	var p config.Parser = &parser.Parser{}
	err := p.ParseArgs(args)
	if err != nil {
		t.Errorf("ParseArgs failed: %v", err)
	}
	want := map[string]string{"execName": "/path/to/prog", "listen": "0.0.0.0"}

	if !reflect.DeepEqual(want, p.ProgramConfig()) {
		t.Fatalf("Expected %+v, got %+v", want, p.ProgramConfig())
	}
}
