package connectivity_test

import (
	"fmt"
	"github/tqtcloud/cmdb/provider/aliyun/connectivity"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	client := &connectivity.AliCloudClient{}
	err := env.Parse(client)
	fmt.Println(client.AccessKey, client.AccessSecret, err)
	if err == nil {
		if should.NoError(err) {
			t.Log(err)
			fmt.Println(client.Account())
		}
	} else {
		t.Log(err)
	}

}
