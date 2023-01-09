package t

import (
	"context"

	"github.com/projecteru2/core/resource3/plugins"
	plugintypes "github.com/projecteru2/core/resource3/plugins/types"
	"github.com/sanity-io/litter"
)

func CalculateDeploy(ctx context.Context, p plugins.Plugin, nodename string) {
	//CalculateDeploy(ctx context.Context, nodename string, deployCount int, resourceRequest *plugintypes.WorkloadResourceRequest) (*plugintypes.CalculateDeployResponse, error)
	req := &plugintypes.WorkloadResourceRequest{
		"cpu-bind":       true,
		"cpu-request":    0.5,
		"cpu-limit":      1,
		"memory-request": "1kb",
		"memory-limit":   "1kb",
	}
	r, err := p.CalculateDeploy(ctx, nodename, 100, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func CalculateRealloc(ctx context.Context, p plugins.Plugin, nodename string) {
	//CalculateRealloc(ctx context.Context, nodename string, resource *plugintypes.WorkloadResource, resourceRequest *plugintypes.WorkloadResourceRequest) (*plugintypes.CalculateReallocResponse, error)
	req := &plugintypes.WorkloadResourceRequest{
		"cpu-bind":       false,
		"cpu-request":    -0.1,
		"cpu-limit":      -0.2,
		"memory-request": "1kb",
		"memory-limit":   "1kb",
	}
	origin := &plugintypes.WorkloadResource{
		"cpu_request": 1,
		"cpu_limit":   1,
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
			"cpu_request": 0.1,
			"cpu_limit":   0.2,
			"cpu_map":     map[string]int64{"1": 20},
		},
		"ID2": {
			"cpu_request": 0.1,
			"cpu_limit":   0.2,
		},
	}
	r, err := p.CalculateRemap(ctx, nodename, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}
