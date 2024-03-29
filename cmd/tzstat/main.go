package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/milQA/tzstat/internal/tzstat/api/stat"
	"github.com/milQA/tzstat/internal/tzstat/storage/statrepository/ramstorage"
	"github.com/milQA/tzstat/pkg/grsdown"
	"github.com/milQA/tzstat/pkg/metrics"
)

func main() {
	grsdown.Run(context.Background(),
		func(ctx context.Context) error {
			m := metrics.NewMetrics(Application, Version)
			storage := ramstorage.NewStatRepository()
			statApiMethods := stat.NewApiStat(storage)

			r := chi.NewRouter()
			r.Use(middleware.Logger, middleware.Recoverer)

			r.Route("/api", func(r chi.Router) {
				r.Route("/stat", func(r chi.Router) {
					r.Get("/", m.Http.Middleware("get_events_average", statApiMethods.GetEventsAverage()).ServeHTTP)
					r.Post("/", m.Http.Middleware("set_event", statApiMethods.SetEvent()).ServeHTTP)
				})
			})

			r.Get("/metrics", m.Handler.ServeHTTP)

			log.Printf("app %s started. version: %s", Application, Version)

			return http.ListenAndServe(":8080", r)
		})
}
