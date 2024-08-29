package iso11649

import (
	"fmt"
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
	mod97, _ := strconv.Atoi(preResult)
	checksum := 98 - mod97%97
	// Pad the checksum to ensure it's at least 2 digits
	return fmt.Sprintf("%02d", checksum)
}

// generateRfReference generates the RF reference from the input string.
func GenerateRfReference(input string) string {
	normalizedRef := strings.ToUpper(strings.TrimSpace(input))
	checksum := calculateRfChecksum(normalizedRef)
	rfReference := "RF" + checksum + normalizedRef
	return strings.TrimSpace(rfReference)
}
