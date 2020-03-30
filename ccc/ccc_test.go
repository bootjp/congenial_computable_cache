package ccc

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	v, err := NewConfig("./test/config_valid.yml")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name     string
		config   *Config
		want     bool
		hasError bool
	}{
		{
			"valid test",
			v,
			true,
			false,
		},
	}
	for _, test := range tests {
		res, err := NewCCC().LoadConfig(test.config)
		if res != test.want {
			t.Errorf("%s want %t given %v", test.name, test.want, res)
		}
		if (err != nil) != test.hasError {
			t.Errorf("%s hasError %t given %v", test.name, test.hasError, err)
		}
	}
}
