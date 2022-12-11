package stat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (a *ApiStat) SetEvent() http.Handler {
	type args struct {
		Date  time.Time `json:"datetime"`
		Value int64     `json:"value"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var args args
		if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			log.Printf("cannot parse body: %s", err)

			return
		}

		if err := a.storage.AddEvent(r.Context(), args.Date, args.Value); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			log.Printf("cannot add event to storage: %s", err)

			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
