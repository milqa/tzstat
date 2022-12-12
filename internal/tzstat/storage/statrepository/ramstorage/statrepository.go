package ramstorage

import (
	"context"
	"sync"
	"time"

	"github.com/shopspring/decimal"

	"github.com/milQA/tzstat/internal/tzstat/storage/statrepository"
)

type StatRepository struct {
	storage *storage
}

func NewStatRepository() statrepository.StatRepository {
	return &StatRepository{
		storage: &storage{
			mu:   sync.RWMutex{},
			data: make([]event, 0, 1024),
		},
	}
}

func (r *StatRepository) AddEvent(_ context.Context, datetime time.Time, value int64) error {
	r.storage.insert(datetime.Unix(), value)

	return nil
}

func (r *StatRepository) GetEventsAvr(_ context.Context, timeFrom, timeTo time.Time) (decimal.Decimal, error) {
	data := r.storage.getEventsWithDatetime(timeFrom.Unix(), timeTo.Unix())

	if len(data) == 0 {
		return decimal.New(0, 0), statrepository.ErrEventsNotFound
	}

	var sum int64
	for _, v := range data {
		sum += v.value
	}

	return decimal.New(sum, 0).Div(decimal.New(int64(len(data)), 0)), nil
}
