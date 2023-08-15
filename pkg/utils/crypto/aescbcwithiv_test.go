package crypto

import (
	"testing"

	"github.com/upmio/upm-pkg/pkg/vars"
)

func TestAescbcwithiv(t *testing.T) {
	origin := "hello world"

	crypted, err := AES_CBC_Encrypt([]byte(origin), []byte(vars.AescbcwithivKey))
	if err != nil {
		t.Fatalf("error:%v", err)
	}
	t.Logf("pwd:%s, crypted:%v (len:%d)", origin, crypted, len(crypted))

	out, err := AES_CBC_Decrypt(crypted, []byte(vars.AescbcwithivKey))
	if err != nil {
		t.Fatalf("error:%v", err)
	}
	t.Logf("origin:%s", string(out))

	if origin != string(out) {
		t.Logf("%s != %s", origin, out)
	}
}