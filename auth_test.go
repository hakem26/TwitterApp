package twitterapp

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Sanitize(t *testing.T)  {
	input := RegisterInput{
		Username: " bob ",
		Email: " Bob@gmail.Com ",
		Password: "Pas123",
		ConfirmPassword: "Pas123",
	}

	want := RegisterInput{
		Username: "bob",
		Email: "bob@gmail.com",
		Password: "Pas123",
		ConfirmPassword: "Pas123",
	}

	input.Sanitize()
	require.Equal(t, want, input)
}

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
		{
			name: "invalid email",
			input: RegisterInput{
				Username: "hkm",
				Email: "hkmmail.co",
				Password: "pas123",
				ConfirmPassword: "pas123",
			},
			err: ErrValidation,
		},
		{
			name: "too short username",
			input: RegisterInput{
				Username: "h",
				Email: "hkm@mail.co",
				Password: "pas123",
				ConfirmPassword: "pas123",
			},
			err: ErrValidation,
		},
		{
			name: "too short password",
			input: RegisterInput{
				Username: "hkm",
				Email: "hkm@mail.co",
				Password: "pas1",
				ConfirmPassword: "pas1",
			},
			err: ErrValidation,
		},
		{
			name: "confirm password did not match",
			input: RegisterInput{
				Username: "hkm",
				Email: "hkm@mail.co",
				Password: "pas123",
				ConfirmPassword: "pas456",
			},
			err: ErrValidation,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err:= tc.input.Validate()
			if err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}