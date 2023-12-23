package security

import "testing"

func TestHashPassword2(t *testing.T) {
	hashPassword, err := HashPassword("123456")

	if err != nil {
		t.Errorf("hash password failed, err: %v", err)
		return
	}

	t.Logf("hash password %v\n", hashPassword)

	isOk := CheckPassword("123456", hashPassword)
	t.Logf("check %v\n", isOk)
}
