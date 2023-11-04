package utils

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("vino@test.com", "vino")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token:%v\n", token)
}

func TestVerifyJWT(t *testing.T) {
	token, err := GenerateJWT("vino@test.com", "vino")
	if err != nil {
		t.Error(err)
		return
	}

	claim, ok, err := VerifyJWT(token)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("verify failed")
		return
	}
	t.Logf("claim: %v\n", JsonMarshal(claim))
}
