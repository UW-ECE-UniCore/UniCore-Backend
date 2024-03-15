package utils

import (
	"fmt"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	fmt.Println(generateRandomCode())
}

func TestSendValidationEmail(t *testing.T) {
	res := SendValidationEmail("jhz.travis@outlook.com", generateRandomCode())
	if res != nil {
		panic(res)
	}
}
