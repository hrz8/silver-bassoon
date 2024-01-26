package main

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/rs/cors"

	OrderUsecase "github.com/hrz8/silver-bassoon/internal/domain/order/usecase"
	"github.com/hrz8/silver-bassoon/pkg/logger"
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
	router.Handle("/api/orders", c.Handler(http.HandlerFunc(OrderUsecase.GetCustomerOrders(db))))

	server := &http.Server{Addr: fmt.Sprintf("127.0.0.1:%s", "3980")}
	server.Handler = router

	err := make(chan error, 1)

	go func() {
		logger.Info("http start at :3980")
		err <- server.ListenAndServe()
	}()

	return err
}
