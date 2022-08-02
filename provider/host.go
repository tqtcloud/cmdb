package provider

import (
	"context"
	"github.com/infraboard/mcube/pager"
	"github/tqtcloud/cmdb/apps/host"
)

type HostOperator interface {
	QueryHost(req *QueryHostRequest) pager.Pager
	QueryDisk(req *QueryDiskRequest) pager.Pager
	DescribeHost(ctx context.Context, req *DescribeHostRequest) (*host.Host, error)
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		Rate: 5,
	}
}

func NewQueryHostRequestWithRate(rate int32) *QueryHostRequest {
	return &QueryHostRequest{
		Rate: float64(rate),
	}
}

type QueryHostRequest struct {
	Rate float64 `json:"rate"`
}

type DescribeHostRequest struct {
	Id string `json:"id"`
}

func NewQueryDiskRequest() *QueryDiskRequest {
	return &QueryDiskRequest{
		Rate: 5,
	}
}

type QueryDiskRequest struct {
	Rate float64 `json:"rate"`
}
