package addresses_test

import (
	"testing"
	"tests/addresses"
)

type Scenario struct {
	inputAddressType    string
	expectedAddressType string
}

func TestAddressType(t *testing.T) {
	scenarios := []Scenario{
		{"Rua British", "Rua"},
		{"Avenida Indian", "Avenida"},
		{"Praça Ocean", "Invalid type"},
		{"Rodovia Territory", "Rodovia"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Invalid type"},
	}

	for _, scenario := range scenarios {
		addressType := addresses.AddressType(scenario.inputAddressType)

		if addressType != scenario.expectedAddressType {
			t.Errorf("Address type is different from expected! Expected %v, got %v",
				scenario.expectedAddressType,
				addressType,
			)
		}
	}
}
