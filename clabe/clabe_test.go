package clabe_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/tbd54566975/mx-clabe-go/clabe"
)

func TestDecode_Valid(t *testing.T) {
	vectors := []struct {
		input    string
		expected clabe.CLABE
	}{
		{
			input: "140180009000015533",
			expected: clabe.CLABE{
				Bank:          clabe.BankMap["140"],
				PlazaCode:     "180",
				AccountNumber: "00900001553",
				CheckDigit:    "3",
			},
		},
		{
			input: "002010077777777771",
			expected: clabe.CLABE{
				Bank:          clabe.BankMap["002"],
				PlazaCode:     "010",
				AccountNumber: "07777777777",
				CheckDigit:    "1",
			},
		},
		{
			input: "014027000005555558",
			expected: clabe.CLABE{
				Bank:          clabe.BankMap["014"],
				PlazaCode:     "027",
				AccountNumber: "00000555555",
				CheckDigit:    "8",
			},
		},
	}

	for _, tt := range vectors {
		t.Run(tt.input, func(t *testing.T) {
			actual, err := clabe.Decode(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDecode_Invalid(t *testing.T) {
	vectors := []string{
		"",
		"7invalid_clabe7",
		"THIS_IS_18_LETTERS",
		"999180009000015533",
		"002010077777777779",
	}

	for _, tt := range vectors {
		t.Run(tt, func(t *testing.T) {
			cl, err := clabe.Decode(tt)
			assert.Error(t, err)
			assert.Zero(t, cl)
		})
	}
}
