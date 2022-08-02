package connectivity

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/stretchr/testify/assert"
	"github.com/tqtcloud/cmdb/provider/aliyun/connectivity"
	"testing"
)

func TestClient(t *testing.T) {
	should := assert.New(t)

	client := &connectivity.AliCloudClient{}
	if err := env.Parse(client); err != nil {
		if should.NoError(err) {
			fmt.Println(client.Account())
		}
	}

}
