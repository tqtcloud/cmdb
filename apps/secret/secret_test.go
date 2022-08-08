package secret_test

import (
	"github/tqtcloud/cmdb/apps/secret"
	"testing"
)

var (
	encryptKey = "qwerty"
)

func TestSecretEncrypt(t *testing.T) {
	ins := secret.NewDefaultSecret()
	ins.Data.ApiSecret = "123456"
	ins.Data.EncryptAPISecret(encryptKey)
	t.Log(ins.Data.ApiSecret)

	ins.Data.DecryptAPISecret(encryptKey)
	t.Log(ins.Data.ApiSecret)
}
