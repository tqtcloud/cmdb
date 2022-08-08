package impl_test

import (
	"context"
	_ "github/tqtcloud/cmdb/apps/all"
	"github/tqtcloud/cmdb/apps/secret"
	"github/tqtcloud/cmdb/conf"
	"os"
	"testing"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	ins secret.ServiceServer
)

func TestQuerySecret(t *testing.T) {
	ss, err := ins.QuerySecret(context.Background(), secret.NewQuerySecretRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestDescribeSecret(t *testing.T) {
	ss, err := ins.DescribeSecret(context.Background(), secret.NewDescribeSecretRequest("ca8sio13n7plv10h6fu0"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestCreateSecret(t *testing.T) {
	req := secret.NewCreateSecretRequest()
	req.Description = "测试用例"
	req.ApiKey = os.Getenv("TX_CLOUD_SECRET_ID")
	req.ApiSecret = os.Getenv("TX_CLOUD_SECRET_KEY")
	req.Vendor = 1
	req.AllowRegions = []string{"*"}
	ss, err := ins.CreateSecret(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func init() {
	// 通过环境变量加载测试配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	// 全局日志对象初始化
	zap.DevelopmentSetup()

	// 初始化所有实例
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}

	ins = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
}
