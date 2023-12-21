package security

import "testing"

func TestGeneratorJWT(t *testing.T) {
	token, err := GeneratorJWT("yym", "12@qq.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
