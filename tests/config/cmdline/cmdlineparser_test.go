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
	args := []string{"/path/to/prog", "--a=1", "--b=2"}
	var p config.Parser = &parser.Parser{}
	p.ParseArgs(args)
	want := map[string]string{"execName": "/path/to/prog", "a": "1", "b": "2"}

	if !reflect.DeepEqual(want, p.ProgramConfig()) {
		t.Fatalf("Expected %+v, got %+v", want, p.ProgramConfig())
	}
}
