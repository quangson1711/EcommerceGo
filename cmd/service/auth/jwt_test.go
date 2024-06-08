package auth

import "testing"

func TestCreateJWT(t *testing.T) {
	token, err := CreateJWT(2)

	if err != nil {
		t.Errorf("error creating token: %v", err)
	}

	if token == "" {
		t.Errorf("expect token to be not empty")
	}

}
