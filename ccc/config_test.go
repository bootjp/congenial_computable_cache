package ccc

import (
	"fmt"
	"testing"
)

func Test_NewConfig(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		hasError bool
	}{
		{"valid test", "./test/config_valid.yml", false},
		{"invalid test", "./test/config_invalid.yml", true},
	}

	for _, test := range tests {
		_, err := NewConfig(test.path)
		if (err != nil) != test.hasError {
			fmt.Println(err)
			t.Errorf("%s want %t given %v", test.name, test.hasError, (err != nil) != test.hasError)
		}

	}

}
