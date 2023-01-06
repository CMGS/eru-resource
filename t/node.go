package t

import (
	"context"

	enginetypes "github.com/projecteru2/core/engine/types"
	"github.com/projecteru2/core/resource3/plugins"
	plugintypes "github.com/projecteru2/core/resource3/plugins/types"
	"github.com/sanity-io/litter"
)

func AddNode(ctx context.Context, p plugins.Plugin, nodename string) {
	//AddNode(ctx context.Context, nodename string, resource *plugintypes.NodeResourceRequest, info *enginetypes.Info) (*plugintypes.AddNodeResponse, error)
	req := &plugintypes.NodeResourceRequest{
		"cpu":         2,
		"share":       0,
		"memory":      "1gb",
		"numa-cpu":    []string{"0", "1"},
		"numa-memory": []string{"512mb", "512mb"},
		"xxxx":        "uuuu",
	}
	info := &enginetypes.Info{NCPU: 8, MemTotal: 2048}
	r, err := p.AddNode(ctx, nodename, req, info)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func RemoveNode(ctx context.Context, p plugins.Plugin, nodename string) {
	//RemoveNode(ctx context.Context, nodename string) (*plugintypes.RemoveNodeResponse, error)
	r, err := p.RemoveNode(ctx, nodename)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func GetDeployCapacity(ctx context.Context, p plugins.Plugin, nodenames []string) {
	//GetNodesDeployCapacity(ctx context.Context, nodenames []string, resource *plugintypes.WorkloadResource) (*plugintypes.GetNodesDeployCapacityResponse, error)
	req := &plugintypes.WorkloadResource{
		"keep-cpu-bind":  true,
		"cpu-bind":       false,
		"cpu-request":    0.5,
		"cpu-limit":      1,
		"memory-request": 100,
		"memory-limit":   100,
	}
	r, err := p.GetNodesDeployCapacity(ctx, nodenames, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}
