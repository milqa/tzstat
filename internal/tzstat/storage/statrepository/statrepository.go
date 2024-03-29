package statrepository

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

type StatRepository interface {
	AddEvent(ctx context.Context, datetime time.Time, value int64) error
	GetEventsAvr(ctx context.Context, timeFrom, timeTo time.Time) (decimal.Decimal, error)
}
