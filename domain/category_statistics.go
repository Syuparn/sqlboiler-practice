package domain

import "context"

// TODO: move this model to other package (this does not seem to be domain model...)
type CategoryStatistics struct {
	CategoryID   CategoryID
	CategoryName CategoryName
	NumProducts  uint
}

type SummarizeProductService interface {
	Summarize(context.Context) ([]*CategoryStatistics, error)
}
