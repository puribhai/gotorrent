package bencode

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Value
		wantErr bool
	}{
		// Cast standard slices to your custom 'List' type
		{"list_basic", "l4:spam4:eggse", List{"spam", "eggs"}, false},

		// Ensure nested lists are also cast
		{"list_nested", "ll5:helloee", List{List{"hello"}}, false},

		// Cast maps to your custom 'Dict' type
		{"dict_basic", "d3:cow3:moo4:spam4:eggse", Dict{"cow": "moo", "spam": "eggs"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := bufio.NewReader(strings.NewReader(tt.input))
			got, err := Decode(r)

			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
