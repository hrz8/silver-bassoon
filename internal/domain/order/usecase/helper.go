package usecase

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func extractPagination(r *http.Request) (page, limit int) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "5"
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}

	return page, limit
}

func extractKeyword(r *http.Request) string {
	keyword := r.URL.Query().Get("keyword")
	if keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", strings.ToLower(keyword))
	}

	return keyword
}

func extractDateFilter(r *http.Request) (bool, pgtype.Timestamp, pgtype.Timestamp) {
	var usingDateFilter bool
	var startDate pgtype.Timestamp
	var endDate pgtype.Timestamp

	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")
	if startDateStr != "" && endDateStr != "" {
		startDateTemp := transformToPGTimestamp(startDateStr)
		endDateTemp := transformToPGTimestamp(endDateStr)

		if startDateTemp.Valid && endDateTemp.Valid {
			usingDateFilter = true
			startDate = startDateTemp
			endDate = endDateTemp
		}
	}

	return usingDateFilter, startDate, endDate
}

func transformToPGTimestamp(timestampString string) pgtype.Timestamp {
	var result pgtype.Timestamp

	timestamp, err := time.Parse("2006-01-02T15:04:05.000", timestampString)
	if err != nil {
		result.Valid = false
	} else {
		result.Valid = true
	}

	result.Time = timestamp

	return result
}
