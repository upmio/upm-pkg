package crypto

import (
	"testing"

	"tesseract/pkg/vars"
)

func TestCbc256WithSalt(t *testing.T) {
	origin := "root"

	crypted, err := Encrypt(origin, vars.CBC256SeCretAESKey)
	if err != nil {
		t.Fatalf("error:%v", err)
	}
	t.Logf("pwd:%s,crypted:%v(len:%d)", origin, crypted, len(crypted))

	out, err := Decrypt([]byte(crypted), vars.CBC256SeCretAESKey)
	if err != nil {
		t.Fatalf("error:%v", err)
	}
	t.Logf("origin:%s", string(out))

	if origin != string(out) {
		t.Logf("%s != %s", origin, out)
	}
}
