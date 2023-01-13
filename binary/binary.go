package main

import (
	"context"
	"time"

	"github.com/CMGS/eru-resource/b"
	"github.com/CMGS/eru-resource/t"
	"github.com/projecteru2/core/log"
	"github.com/projecteru2/core/resource3/plugins/binary"
	coretypes "github.com/projecteru2/core/types"
)

func main() {
	ctx := context.TODO()
	path := "test"

	log.SetupLog(ctx, "debug", "")

	config := coretypes.Config{
		ResourcePlugin: coretypes.ResourcePluginConfig{
			Dir:         "./binary",
			CallTimeout: time.Second * 300,
		},
	}
	bi, err := binary.NewPlugin(ctx, path, config)
	if err != nil {
		panic(err)
	}

	//t.Name(bi)
	//t.GetMetricsDescription(ctx, bi)
	b.AddNode(ctx, bi, "test")
	defer t.RemoveNode(ctx, bi, "test")
	//
	// t.GetMetrics(ctx, bi, "test")
	//
	//	b.SetNodeResourceCapacity(ctx, bi, "test")
	//	b.GetNodesDeployCapacity(ctx, bi, []string{"test"})
	b.GetNodeResourceInfo(ctx, bi, "test")
	b.SetNodeResourceInfo(ctx, bi, "test")
	b.GetNodeResourceInfo(ctx, bi, "test")
}
