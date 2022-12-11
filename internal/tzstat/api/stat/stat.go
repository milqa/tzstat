package stat

import (
	"github.com/milQA/tzstat/internal/tzstat/api"
	"github.com/milQA/tzstat/internal/tzstat/storage/statrepository"
)

type ApiStat struct {
	storage statrepository.StatRepository
}

func NewApiStat(storage statrepository.StatRepository) *ApiStat {
	return &ApiStat{storage: storage}
}

// TODO: переписать и использовать это
func (a *ApiStat) Handlers() []api.Method {
	return []api.Method{
		{
			Name:    "set_event",
			Path:    "/set",
			Handler: a.SetEvent(),
		},
		{
			Name:    "get_events_average",
			Path:    "/get",
			Handler: a.GetEventsAverage(),
		},
	}
}
