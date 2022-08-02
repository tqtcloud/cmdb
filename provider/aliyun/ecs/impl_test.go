package ecs_test

import (
	"context"
	"fmt"
	"github/tqtcloud/cmdb/apps/disk"
	"github/tqtcloud/cmdb/apps/host"
	"github/tqtcloud/cmdb/provider"
	"github/tqtcloud/cmdb/provider/aliyun"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
)

var (
	operator provider.HostOperator
)

func TestQueryEcs(t *testing.T) {
	req := provider.NewQueryHostRequest()
	pager := operator.QueryHost(req)
	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribe(t *testing.T) {
	req := &provider.DescribeHostRequest{Id: "i-t4ne8pajfrx4jmp9porx"}
	ins, err := operator.DescribeHost(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryDisk(t *testing.T) {
	req := provider.NewQueryDiskRequest()
	pager := operator.QueryDisk(req)
	for pager.Next() {
		set := disk.NewDiskSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}

		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().HostOperator()
}
