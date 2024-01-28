package usecase

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRequest struct {
	mock.Mock
	request http.Request
}

func (m *MockRequest) URL() *url.URL {
	args := m.Called()
	return args.Get(0).(*url.URL)
}

func NewMockRequest() *MockRequest {
	return &MockRequest{
		request: http.Request{URL: &url.URL{}},
	}
}

func TestExtractPagination(t *testing.T) {
	mockRequest := NewMockRequest()
	mockRequest.request.URL.RawQuery = "page=2&limit=10"

	page, limit := extractPagination(&mockRequest.request)

	assert.Equal(t, 2, page)
	assert.Equal(t, 10, limit)
}

func TestExtractKeyword(t *testing.T) {
	mockRequest := NewMockRequest()
	mockRequest.request.URL.RawQuery = "keyword=test"

	keyword := extractKeyword(&mockRequest.request)

	assert.Equal(t, "%test%", keyword)

	mockRequest.AssertExpectations(t)
}

func TestExtractDateFilter(t *testing.T) {
	mockRequest := NewMockRequest()
	mockRequest.request.URL.RawQuery = "start_date=2022-01-01T00:00:00.000&end_date=2022-01-31T23:59:59.999"

	usingDateFilter, startDate, endDate := extractDateFilter(&mockRequest.request)

	assert.True(t, usingDateFilter)
	assert.True(t, startDate.Valid)
	assert.True(t, endDate.Valid)

	mockRequest.AssertExpectations(t)
}

func TestTransformToPGTimestamp(t *testing.T) {
	result := transformToPGTimestamp("2022-01-01T12:34:56.789")

	assert.True(t, result.Valid)
	expectedTime := time.Date(2022, time.January, 1, 12, 34, 56, 789000000, time.UTC)
	assert.Equal(t, expectedTime, result.Time)

	result = transformToPGTimestamp("invalid_timestamp")

	assert.False(t, result.Valid)
}
