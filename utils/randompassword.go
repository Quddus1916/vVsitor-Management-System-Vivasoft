package utils

import (
	"github.com/sethvargo/go-password/password"
)

func GenerateRandomPassword() (string, error) {
	res, err := password.Generate(64, 10, 10, false, false)

	return res, err

}
