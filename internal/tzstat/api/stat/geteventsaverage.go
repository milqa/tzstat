package stat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (a *ApiStat) GetEventsAverage() http.Handler {
	type (
		args struct {
			DateFrom time.Time `json:"date_from"`
			DateTo   time.Time `json:"date_to"`
		}

		result struct {
			Value int64 `json:"value"`
		}
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var args args
		if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			log.Printf("cannot parse body: %s", err)

			return
		}

		average, err := a.storage.GetEventsAvr(r.Context(), args.DateFrom, args.DateTo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("cannot get events average from storage: %s", err)

			return
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
