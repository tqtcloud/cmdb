package connectivity_test

import (
	"github/tqtcloud/cmdb/provider/txyun/connectivity"
	"testing"
)

func TestTencentCloudClient(t *testing.T) {
	conn := connectivity.C()
	if err := conn.Check(); err != nil {
		t.Fatal(err)
	}

	t.Log(conn.AccountID(), conn.Region)
}

func init() {
	//  初始化client
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
