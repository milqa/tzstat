package statrepository

import (
	"context"
	"time"
)

type StatRepository interface {
	AddEvent(ctx context.Context, datetime time.Time, value int64) error
	GetEventsAvr(ctx context.Context, timeFrom, timeTo time.Time) (int64, error)
}
