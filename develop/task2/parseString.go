package parsestring

import (
	"errors"
	"strconv"
	"strings"
)

func parseString(stringForUnpack string) (string, error) {
	// Check if the string empty
	if len(stringForUnpack) == 0 {
		return "", nil
	}
	// Create new string builder
	sb := new(strings.Builder)
	// Convert string to lowercase rune array
	runeArray := []rune(strings.ToLower(stringForUnpack))
	// Check for wrong first character
	if (runeArray[0] < 97 && runeArray[0] != 92) || runeArray[0] > 122 {
		return "", errors.New("wrong first symbol")
	}
	for i := 0; i <= len(runeArray)-1; i++ {
		// If the character is a '/'
		if runeArray[i] == 92 {
			sb.WriteRune(runeArray[i+1])
			i++
			// If the character is a liter
		} else if runeArray[i] >= 97 && runeArray[i] <= 122 {
			sb.WriteRune(runeArray[i])
			// If the character is a digit
		} else if runeArray[i] >= 48 && runeArray[i] <= 57 {
			num := new(strings.Builder)
			for _, symb := range runeArray[i:] {
				if symb >= 48 && symb <= 57 {
					num.WriteRune(symb)
				} else {
					break
				}
			}
			numAsInt, err := strconv.Atoi(num.String())
			if err != nil {
				return "", err
			}
			sb.WriteString(strings.Repeat(string(runeArray[i-1]), numAsInt-1))
		}
	}
	return sb.String(), nil
}
