package checks

import (
	"sync"

	"github.com/jexia/semaphore/pkg/core/instance"
	"github.com/jexia/semaphore/pkg/core/logger"
	"github.com/jexia/semaphore/pkg/core/trace"
	"github.com/jexia/semaphore/pkg/specs"
	"github.com/jexia/semaphore/pkg/specs/template"
)

// ReservedKeywords represents a list with reserved keywords
var ReservedKeywords = []string{
	template.InputResource,
	template.ErrorResource,
	template.StackResource,
}

// FlowDuplicates checks for duplicate definitions
func FlowDuplicates(ctx instance.Context, flows specs.FlowListInterface) error {
	ctx.Logger(logger.Core).Info("Checking manifest duplicates")

	tracker := sync.Map{}

	for _, flow := range flows {
		_, duplicate := tracker.LoadOrStore(flow.GetName(), flow)
		if duplicate {
			return trace.New(trace.WithMessage("duplicate flow '%s'", flow.GetName()))
		}

		err := NodeDuplicates(ctx, flow.GetName(), flow.GetNodes())
		if err != nil {
			return err
		}
	}

	return nil
}

// NodeDuplicates checks for duplicate definitions
func NodeDuplicates(ctx instance.Context, flow string, nodes []*specs.Node) error {
	ctx.Logger(logger.Core).Info("Checking flow duplicates")

	calls := sync.Map{}

	for _, node := range nodes {
		_, duplicate := calls.LoadOrStore(node.ID, node)
		if duplicate {
			return trace.New(trace.WithMessage("duplicate resource '%s' in flow '%s'", node.ID, flow))
		}

		for _, key := range ReservedKeywords {
			if key != node.ID {
				continue
			}

			return trace.New(trace.WithMessage("flow with the name '%s' is a reserved keyword", node.ID))
		}
	}

	return nil
}
