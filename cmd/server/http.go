package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	psqlrepo "github.com/hrz8/silver-bassoon/internal/repo/psql"
	"github.com/jackc/pgx/v5"
	"github.com/rs/cors"
)

func start(db *pgx.Conn) chan error {
	var err chan error

	go func() {
		err = deliver(db)
	}()

	return err
}

func deliver(db *pgx.Conn) chan error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
	})

	router := http.NewServeMux()
	router.Handle("/", c.Handler(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()

			if r.Method == http.MethodGet {
				w.Header().Set("Content-Type", "application/json")

				keyword := fmt.Sprintf("%%%s%%", "")

				queries := psqlrepo.New(db)
				items, err := queries.GetCustomerOrders(ctx, &psqlrepo.GetCustomerOrdersParams{
					IsSearchTerm:    false,
					SearchTerm:      &keyword,
					UsingDateFilter: false,
					UsingPagination: false,
					PageNumber:      1,
				})
				if err != nil {
					http.Error(w, "", http.StatusBadRequest)
					return
				}

				json.NewEncoder(w).Encode(items)
				return
			}

			http.Error(w, "", http.StatusBadRequest)
		},
	)))

	server := &http.Server{Addr: fmt.Sprintf("127.0.0.1:%s", "3900")}
	server.Handler = router

	err := make(chan error, 1)

	go func() {
		err <- server.ListenAndServe()
	}()

	return err
}
