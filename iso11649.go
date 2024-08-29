package iso11649

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// charToDigitMap maps alphanumeric characters to their corresponding digits.
var charToDigitMap = map[rune]string{
	'A': "10", 'B': "11", 'C': "12", 'D': "13", 'E': "14", 'F': "15",
	'G': "16", 'H': "17", 'I': "18", 'J': "19", 'K': "20", 'L': "21",
	'M': "22", 'N': "23", 'O': "24", 'P': "25", 'Q': "26", 'R': "27",
	'S': "28", 'T': "29", 'U': "30", 'V': "31", 'W': "32", 'X': "33",
	'Y': "34", 'Z': "35",
}

// replaceChars replaces alphanumeric characters in the input string with corresponding numeric values.
func replaceChars(input string) string {
	var sb strings.Builder
	for _, char := range input {
		if unicode.IsDigit(char) {
			sb.WriteRune(char)
		} else if unicode.IsLetter(char) {
			sb.WriteString(charToDigitMap[unicode.ToUpper(char)])
		}
	}
	return sb.String()
}

// calculateRfChecksum calculates the RF checksum for a given reference string.
func calculateRfChecksum(ref string) string {
	preResult := ref + "RF00"
	preResult = replaceChars(preResult)
	val := new(big.Int)
	val, _ = val.SetString(preResult, 10)
	mod97 := big.NewInt(int64(97))
	result := new(big.Int)
	result = result.Mod(val, mod97)
	res, _ := strconv.Atoi(result.String())
	checksum := 98 - res
	// Pad the checksum to ensure it's at least 2 digits
	return fmt.Sprintf("%02d", checksum)

}

// GenerateReference generates the RF reference from the input string.
func GenerateReference(input string) (string, error) {
	if !isAlphanum(input) {
		return "", errors.New("input is not alphanumeric")
	}
	normalizedRef := strings.ToUpper(strings.TrimSpace(input))
	checksum := calculateRfChecksum(normalizedRef)
	rfReference := "RF" + checksum + normalizedRef
	return strings.TrimSpace(rfReference), nil
}

func isAlphanum(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
}
