package utils

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func GetPwd() (string, error) {
	validatePassword := func(input string) error {
		if len(input) < 6 {
			return errors.New("Password must have more than 6 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validatePassword,
		Mask:     '*',
	}

	result, err := prompt.Run()
	return result, err
}

func GetUsername() (string, error) {
	prompt := promptui.Prompt{
		Label: "Username",
	}

	result, err := prompt.Run()
	return result, err
}
