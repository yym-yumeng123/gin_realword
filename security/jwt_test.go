package security

import (
	"gin_realword/utils"
	"testing"
)

func TestGeneratorJWT(t *testing.T) {
	token, err := GeneratorJWT("yym", "12@qq.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestVerifyJWT(t *testing.T) {
	claim, valid, err := VerifyJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDMzOTcwNDIsImlhdCI6MTcwMzMxMDY0MiwidXNlciI6eyJlbWFpbCI6IjEyQHFxLmNvbSIsInVzZXJuYW1lIjoieXltIn19.xGvz4K4mWE_JvP_Qad3pC59i77jS6j8Yj7iJQm6Q-Jw")

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("verify jwt: %v\n, claim: %v\n", valid, utils.JsonMarshal(claim))
}
