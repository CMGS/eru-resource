package b

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
		"xxxx":    "uuuu",
		"volum":   "xxxx",
		"volumes": []string{"/data:11G", "/data1:10G", "/data2:10G"},
		"storage": "10G",
		"disks":   []string{"/dev/vda:/data1:1000:1000:10M:10M", "/dev/vdb:/data,/data2:1000:1000:10M:10M"},
	}
	info := &enginetypes.Info{StorageTotal: 1024000}
	r, err := p.AddNode(ctx, nodename, req, info)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func GetNodesDeployCapacity(ctx context.Context, p plugins.Plugin, nodenames []string) {
	//GetNodesDeployCapacity(ctx context.Context, nodenames []string, resource *plugintypes.WorkloadResourceRequest) (*plugintypes.GetNodesDeployCapacityResponse, error)
	req := &plugintypes.WorkloadResourceRequest{
		"volume": []string{"AUTO:/save:rw:10G:100:100:1M:1M"},
	}
	r, err := p.GetNodesDeployCapacity(ctx, nodenames, req)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}

func SetNodeResourceCapacity(ctx context.Context, p plugins.Plugin, nodename string) {
	req := &plugintypes.NodeResourceRequest{
		"xxxx":    "uuuu",
		"volum":   "xxxx",
		"volumes": []string{"/data:1", "/data1:1", "/data2:1"},
		"storage": "1",
		"disks":   []string{"/dev/vda:/data1:1:1:1:1", "/dev/vdb:/data,/data2:1:1:1:1"},
	}

	r, err := p.SetNodeResourceCapacity(ctx, nodename, nil, req, true, true)
	if err != nil {
		panic(err)
	}
	litter.Dump(r)
}
