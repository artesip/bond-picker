package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation(t *testing.T) {
	tests := []struct {
		name    string
		user    *User
		mustErr bool
	}{
		{"nil user", nil, true},
		{"empty struct", &User{}, true},
		{"empty username", &User{Username: ""}, true},
		{"whitespace username", &User{Username: "     "}, true},
		{"missing password", &User{Username: "test"}, true},
		{"empty password string", &User{Username: "test", Password: ""}, true},
		{"both empty", &User{Username: "", Password: ""}, true},
		{"valid short user", &User{Username: "test", Password: "testtest"}, false},
		{"valid long user", &User{Username: "testtest", Password: "testtest123"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()

			if tt.mustErr {
				assert.Error(t, err)
				assert.Equal(t, ValidationErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
