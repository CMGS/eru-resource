package t

import (
	"context"

	enginetypes "github.com/projecteru2/core/engine/types"
	"github.com/projecteru2/core/resource3/plugins"
	plugintypes "github.com/projecteru2/core/resource3/plugins/types"
	"github.com/sanity-io/litter"
)

func AddNode(ctx context.Context, p plugins.Plugin) {
	//AddNode(ctx context.Context, nodename string, resource *plugintypes.NodeResourceRequest, info *enginetypes.Info) (*plugintypes.AddNodeResponse, error)
	nodename := "test"
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
