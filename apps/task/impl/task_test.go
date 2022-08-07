package impl_test

import (
	"context"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"github/tqtcloud/cmdb/apps/resource"
	"github/tqtcloud/cmdb/apps/task"
	"github/tqtcloud/cmdb/conf"
	"os"
	"testing"
)

var (
	ins task.ServiceServer
)

func TestCreateTask(t *testing.T) {
	req := task.NewCreateTaskRequst()
	req.Type = task.Type_RESOURCE_SYNC
	req.Region = os.Getenv("TX_CLOUD_REGION")
	req.ResourceType = resource.Type_HOST
	req.SecretId = "ca8sio13n7plv10h6fu0"
	taskIns, err := ins.CreateTask(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(taskIns)
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

	ins = app.GetGrpcApp(task.AppName).(task.ServiceServer)
}
