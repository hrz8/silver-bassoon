package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	psqlrepo "github.com/hrz8/silver-bassoon/internal/repo/psql"
	"github.com/hrz8/silver-bassoon/pkg/response"
	"github.com/jackc/pgx/v5"
)

func GetCustomerOrders(db *pgx.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		if r.Method == http.MethodGet {
			tz := r.Header.Get("X-Time-Zone")
			if tz == "" {
				tz = "UTC+11:00"
			}

			page, limit := extractPagination(r)
			keyword := extractKeyword(r)

			payload := &psqlrepo.GetCustomerOrdersParams{
				TimeZone:        tz,
				IsSearchTerm:    keyword != "",
				SearchTerm:      &keyword,
				UsingDateFilter: false,
				UsingPagination: false,
			}

			usingDateFilter, startDate, endDate := extractDateFilter(r)
			payload.UsingDateFilter = usingDateFilter
			payload.StartDate = startDate
			payload.EndDate = endDate

			queries := psqlrepo.New(db)
			// count
			itemsAll, err := queries.GetCustomerOrders(ctx, payload)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// rows
			payload.UsingPagination = true
			payload.PageNumber = page
			payload.PageSize = int32(limit)
			items, err := queries.GetCustomerOrders(ctx, payload)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			result := items
			if result == nil {
				result = make([]*psqlrepo.GetCustomerOrdersRow, 0)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response.HTTPResponse[[]*psqlrepo.GetCustomerOrdersRow]{
				Message: "success fetch customer orders",
				Result:  &result,
				Meta: &response.MetaResponse{
					Count: len(items),
					Total: len(itemsAll),
				},
			})
			return
		}

		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}
