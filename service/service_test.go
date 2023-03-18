package service

import "testing"

func TestSetToken(t *testing.T) {
	token, err := SetToken("joker")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s", token)
}
