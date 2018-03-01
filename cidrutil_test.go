package cidrutil

import (
	"fmt"
	"reflect"
	"testing"
)

var valid = []struct {
    cidr  string
    network *Network
}{
	{"192.168.0.1/32", &Network{
		CIDR: "192.168.0.1/32",
		Ones: 32,
		Bits: 32,
		Network: "",
		Hosts: []string{
			"192.168.0.1",
		},
		Broadcast: "",
	}},
	{"192.168.0.0/31", &Network{
		CIDR: "192.168.0.0/31",
		Ones: 31,
		Bits: 32,
		Network: "",
		Hosts: []string{
			"192.168.0.0",
			"192.168.0.1",
		},
		Broadcast: "",
	}},
	{"192.168.0.0/30", &Network{
		CIDR: "192.168.0.0/30",
		Ones: 30,
		Bits: 32,
		Network: "192.168.0.0",
		Hosts: []string{
			"192.168.0.1",
			"192.168.0.2",
		},
		Broadcast: "192.168.0.3",
	}},
}

var invalid = []struct {
    cidr  string
	err error
}{
	{"192.168.0.0/33", fmt.Errorf("invalid CIDR address: %s", "192.168.0.0/33")},
}

func TestParse (t *testing.T) {
	p := NewParser()

	t.Run("valid data", func(t *testing.T ) {
		for _, data := range valid {
			actual, _ := p.Parse(data.cidr)
			expected := data.network

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("\nactual:\t\t%v\nexpected:\t%v", actual, expected)
			}
		}
	})

	t.Run("invalid data", func(t *testing.T ) {
		for _, data := range invalid {
			_, actual := p.Parse(data.cidr)
			expected := data.err

			if actual == expected {
				t.Errorf("got: %v\nexpected: %v", actual, expected)
			}
		}
	})
}
