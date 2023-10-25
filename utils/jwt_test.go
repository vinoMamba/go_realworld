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
