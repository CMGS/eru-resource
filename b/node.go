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
		"volumes": []string{"/data:10G", "/data1:10G", "/data2:10G"},
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
