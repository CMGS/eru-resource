package t

import (
	"context"
	"fmt"

	"github.com/projecteru2/core/resource3/plugins"
	"github.com/sanity-io/litter"
)

func Name(p plugins.Plugin) {
	fmt.Println(p.Name())
}

func GetMetrics(ctx context.Context, p plugins.Plugin) {
	litter.Dump(p.GetMetrics(ctx, "pod1", "test"))
}

func GetMetricsDescription(ctx context.Context, p plugins.Plugin) {
	litter.Dump(p.GetMetricsDescription(ctx))
}
