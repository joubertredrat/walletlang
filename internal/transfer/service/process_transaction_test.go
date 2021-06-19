package service_test

import (
	"testing"
)

func TestProcessTransactionHandleProcess(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Proccess transaction with success",
		},
		{
			name: "Proccess transaction with transaction not found",
		},
		{
			name: "Proccess transaction with payer not found",
		},
		{
			name: "Proccess transaction with payee not found",
		},
		{
			name: "Proccess transaction with insufficient payer funds",
		},
		{
			name: "Proccess transaction with error on wallet movement in payer",
		},
		{
			name: "Proccess transaction with error on wallet movement in payee",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
