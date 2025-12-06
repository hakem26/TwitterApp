package twitterapp

import "testing"

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct{
		name string
		input RegisterInput
		err error
	} {
		{
			name: "valid",
			input: RegisterInput{
				Username: "hkm",
				Email: "hkm@mail.co",
				Password: "pas123",
				ConfirmPassword: "pas123",
			},
			err: nil,
		},
	}
}