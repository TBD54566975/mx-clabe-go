package clabe

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// CLABE (Clave Bancaria Estandarizada, Spanish for "standardized banking cipher" or "standardized bank code")
// is a banking standard for account numbers in Mexico. This standard is a requirement for the sending
// and receiving of domestic inter-bank electronic funds transfer since June 1, 2004
type CLABE struct {
	Value         string
	AccountNumber string
	Bank          Bank
	PlazaCode     string
	CheckDigit    string
}

var (
	CLABERegex = regexp.MustCompile(`^\d{18}$`)
)

// Decode decodes a [CLABE] and returns the encoded bank, plaza code, account number, and check digit
//
// [CLABE]: https://en.wikipedia.org/wiki/CLABE#Structure
func Decode(input string) (CLABE, error) {
	if !CLABERegex.MatchString(input) {
		return CLABE{}, errors.New("invalid CLABE")
	}

	bankCode := input[0:3]
	bank, ok := GetBank(bankCode)
	if !ok {
		return CLABE{}, fmt.Errorf("invalid bank code %s", bankCode)
	}

	checksum, err := Checksum(input)
	if err != nil {
		return CLABE{}, fmt.Errorf("invalid checksum: %w", err)
	}

	checkDigit := input[17:]
	if checksum != checkDigit {
		return CLABE{}, fmt.Errorf("checksum mismatch: expected %s. got %s", checksum, checkDigit)
	}

	return CLABE{
		Bank:          bank,
		PlazaCode:     input[3:6],
		AccountNumber: input[6:17],
		CheckDigit:    checkDigit,
	}, nil
}

// Checksum computes a checksum using an algorithm specific to CLABE. Checksum should match CLABE's last digit
// Algorithm details in spreadsheet [here].
//
// [here]: https://www.frbservices.org/binaries/content/assets/crsocms/financial-services/ach/clabe-check-digit-calc.xls
func Checksum(input string) (string, error) {
	// The nth digit has a weight factor of weights[n % 3]).
	weights := []int{3, 7, 1}

	// Multiply each of first 17 digits by its weight factor, then take only the "ones" digit.
	var productOnes []int
	for index, char := range input[:17] {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return "", fmt.Errorf("failed to convert character to integer: %w", err)
		}

		productOnes = append(productOnes, (digit*weights[index%3])%10)
	}

	// Sum all the "ones" digits.
	sum := 0
	for _, ones := range productOnes {
		sum += ones
	}

	// Get "ones" digit from the sum of productOnes, then subtract from 10.
	checksum := 10 - (sum % 10)

	// If the result is 10, we return 0, as per CLABE standard.
	if checksum == 10 {
		return "", nil
	}

	return strconv.Itoa(checksum), nil
}

// Implementing the encoding.TextMarshaler interface
func (c CLABE) MarshalText() ([]byte, error) {
	return []byte(c.Value), nil
}

// Implementing the encoding.TextUnmarshaler interface
func (c *CLABE) UnmarshalText(text []byte) error {
	c.Value = string(text)
	// Assuming Decode function updates the rest of the fields based on Value
	decoded, err := Decode(c.Value)
	if err != nil {
		return err
	}
	*c = decoded
	return nil
}
