package stat

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/milQA/tzstat/internal/tzstat/storage/statrepository"
)

func (a *ApiStat) GetEventsAverage() http.Handler {
	type (
		result struct {
			Value int64 `json:"value"`
		}
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		argDateTo := r.URL.Query().Get("date_to")
		dateTo, err := time.Parse(time.RFC3339Nano, argDateTo)
		if err != nil {
			dateTo = time.Now()
		}

		dateFrom := dateTo.Add(-time.Minute)

		average, err := a.storage.GetEventsAvr(r.Context(), dateFrom, dateTo)
		if err != nil {
			if !errors.Is(err, statrepository.ErrEventsNotFound) {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("cannot get events average from storage: %s", err)

				return
			}
		}

		result := result{Value: average}
		if err = json.NewEncoder(w).Encode(&result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("cannot encode json: %s", err)

			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
