package ramstorage

import (
	"context"
	"sync"
	"time"

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

func (r *StatRepository) GetEventsAvr(_ context.Context, timeFrom, timeTo time.Time) (int64, error) {
	data := r.storage.getEventsWithDatetime(timeFrom.Unix(), timeTo.Unix())

	if len(data) == 0 {
		return 0, statrepository.ErrEventsNotFound
	}

	var sum int64
	for _, v := range data {
		sum += v.value
	}

	return sum / int64(len(data)), nil
}
