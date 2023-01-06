package main

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/CMGS/eru-resource/t"
	"github.com/projecteru2/core/resource3/plugins/cpumem"
	coretypes "github.com/projecteru2/core/types"
	"go.etcd.io/etcd/tests/v3/integration"
)

func main() {
	ctx := context.TODO()

	config := coretypes.Config{
		Scheduler: coretypes.SchedulerConfig{
			MaxShare:       -1,
			ShareBase:      100,
			MaxDeployCount: 10000,
		},
		Etcd: coretypes.EtcdConfig{
			Prefix:     "/testing",
			LockPrefix: "__lock__",
			Machines:   []string{"http://127.0.0.1:2379"},
		},
	}

	cm, err := cpumem.NewPlugin(ctx, config)
	if err != nil {
		panic(err)
	}

	t.Name(cm)
	t.GetMetrics(ctx, cm)
	t.GetMetricsDescription(ctx, cm)
}

func etcd() *integration.ClusterV3 {
	os.Args = []string{"test.short=false"}
	testing.Init()
	flag.Parse()

	t := &testing.T{}
	integration.BeforeTestExternal(t)
	return integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
}
