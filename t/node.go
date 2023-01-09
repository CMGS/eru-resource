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

func GetNodesDeployCapacity(ctx context.Context, p plugins.Plugin, nodenames []string) {
	//GetNodesDeployCapacity(ctx context.Context, nodenames []string, resource *plugintypes.WorkloadResourceRequest) (*plugintypes.GetNodesDeployCapacityResponse, error)
	req := &plugintypes.WorkloadResourceRequest{
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

func SetNodeResourceCapacity(ctx context.Context, p plugins.Plugin, nodename string) {
	//SetNodeResourceCapacity(ctx context.Context, nodename string, resource *plugintypes.NodeResource, resourceRequest *plugintypes.NodeResourceRequest, delta bool, incr bool) (*plugintypes.SetNodeResourceCapacityResponse, error)
	//	resourceReq := &plugintypes.NodeResourceRequest{
	//		//"cpu":    "1:50,2:50",
	//		"memory": "2gb",
	//	}

	resource := &plugintypes.NodeResource{
		"cpu": 4.0,
		"cpu_map": map[string]int64{
			"1": 50,
			"0": 100,
		},
		"memory": 102400,
	}

	r, err := p.SetNodeResourceCapacity(ctx, nodename, resource, nil, true, true)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func GetNodeResourceInfo(ctx context.Context, p plugins.Plugin, nodename string) {
	//GetNodeResourceInfo(ctx context.Context, nodename string, workloadsResource []*plugintypes.WorkloadResource) (*plugintypes.GetNodeResourceInfoResponse, error)
	req := []*plugintypes.WorkloadResource{
		{
			"cpu_request":    0.5, // TODO maybe not _
			"cpu_limit":      1,
			"memory_request": 100,
			"memory_limit":   100,
		},
	}
	r, err := p.GetNodeResourceInfo(ctx, nodename, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func SetNodeResourceInfo(ctx context.Context, p plugins.Plugin, nodename string) {
	//SetNodeResourceInfo(ctx context.Context, nodename string, capacity *plugintypes.NodeResource, usage *plugintypes.NodeResource) (*plugintypes.SetNodeResourceInfoResponse, error)
	cap := &plugintypes.NodeResource{
		"cpu": 4.0,
		"cpu_map": map[string]int64{
			"1": 50,
			"0": 100,
			"2": 40,
			"3": 60,
		},
		"memory": 102400,
	}
	usage := &plugintypes.NodeResource{}
	r, err := p.SetNodeResourceInfo(ctx, nodename, cap, usage)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func SetNodeResourceUsage(ctx context.Context, p plugins.Plugin, nodename string) {
	//	resourceReq := &plugintypes.NodeResourceRequest{
	//		//"cpu":    "1:50,2:50",
	//		"memory": "2gb",
	//	}

	//resource := &plugintypes.NodeResource{
	//	"cpu": 4.0,
	//	"cpu_map": map[string]int64{
	//		"1": 50,
	//		"0": 100,
	//	},
	//	"memory": 102400,
	//}

	workloadResourcs := []*plugintypes.WorkloadResource{
		{"cpu_request": 1.0},
		{"cpu_request": 1.0},
	}

	r, err := p.SetNodeResourceUsage(ctx, nodename, nil, nil, workloadResourcs, true, true)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func GetMostIdleNode(ctx context.Context, p plugins.Plugin, nodenames []string) {
	r, err := p.GetMostIdleNode(ctx, nodenames)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func FixNodeResource(ctx context.Context, p plugins.Plugin, nodename string) {
	//FixNodeResource(ctx context.Context, nodename string, workloadsResource []*plugintypes.WorkloadResource) (*plugintypes.GetNodeResourceInfoResponse, error)
	req := []*plugintypes.WorkloadResource{
		{
			"cpu_request":    0.5,
			"cpu_limit":      1,
			"memory_request": 100,
			"memory_limit":   100,
		},
	}
	r, err := p.FixNodeResource(ctx, nodename, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}
