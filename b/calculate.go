package b

import (
	"context"

	"github.com/projecteru2/core/resource3/plugins"
	plugintypes "github.com/projecteru2/core/resource3/plugins/types"
	"github.com/sanity-io/litter"
)

func CalculateDeploy(ctx context.Context, p plugins.Plugin, nodename string) {
	//CalculateDeploy(ctx context.Context, nodename string, deployCount int, resourceRequest *plugintypes.WorkloadResourceRequest) (*plugintypes.CalculateDeployResponse, error)
	req := &plugintypes.WorkloadResourceRequest{
		"volume": []string{"AUTO:/save:rw:10G:100:100:1M:1M"},
	}
	r, err := p.CalculateDeploy(ctx, nodename, 1, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func CalculateRealloc(ctx context.Context, p plugins.Plugin, nodename string) {
	//CalculateRealloc(ctx context.Context, nodename string, resource *plugintypes.WorkloadResource, resourceRequest *plugintypes.WorkloadResourceRequest) (*plugintypes.CalculateReallocResponse, error)
	req := &plugintypes.WorkloadResourceRequest{
		"storage": "2mb",
	}
	origin := &plugintypes.WorkloadResource{
		"storage_request": 1024,
		"storage_limit":   1024,
	}
	r, err := p.CalculateRealloc(ctx, nodename, origin, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func CalculateRemap(ctx context.Context, p plugins.Plugin, nodename string) {
	req := map[string]*plugintypes.WorkloadResource{
		"ID1": {
			"storage_request": 1024,
			"storage_limit":   1024,
		},
		"ID2": {
			"storage_request": 1024,
			"storage_limit":   1024,
		},
	}
	r, err := p.CalculateRemap(ctx, nodename, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}
