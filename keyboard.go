// Package keyboard reads user input from keyboard.
package keyboard

import (
	"bufio"
	"strings"
	"strconv"
	"os"
)
// GetFloat reads floating point from keyboard.
// Returs the number read and any error occured.
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}
	return number, nil
}
