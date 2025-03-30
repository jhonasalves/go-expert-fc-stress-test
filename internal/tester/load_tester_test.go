package tester

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunLoadTest_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	totalRequests := 10
	concurrency := 5

	result := RunLoadTest(server.URL, totalRequests, concurrency)

	assert.Equal(t, totalRequests, result.TotalRequests, "TotalRequests mismatch")
	assert.Equal(t, totalRequests, result.SuccessfulRequests, "SuccessfulRequests mismatch")

	count, ok := result.StatusCodes.Load(http.StatusOK)
	assert.True(t, ok, "Status code 200 not found")
	assert.Equal(t, totalRequests, count.(int), "Status code 200 count mismatch")
}

func TestRunLoadTest_PartialSuccess(t *testing.T) {
	var counter int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if counter%2 == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		counter++
	}))
	defer server.Close()

	totalRequests := 10
	concurrency := 5

	result := RunLoadTest(server.URL, totalRequests, concurrency)

	expectedSuccess := totalRequests / 2

	assert.Equal(t, totalRequests, result.TotalRequests, "TotalRequests mismatch")
	assert.Equal(t, expectedSuccess, result.SuccessfulRequests, "SuccessfulRequests mismatch")

	count200, ok := result.StatusCodes.Load(http.StatusOK)
	assert.True(t, ok, "Status code 200 not found")
	assert.Equal(t, expectedSuccess, count200.(int), "Status code 200 count mismatch")

	count500, ok := result.StatusCodes.Load(http.StatusInternalServerError)
	assert.True(t, ok, "Status code 500 not found")
	assert.Equal(t, expectedSuccess, count500.(int), "Status code 500 count mismatch")
}

func TestRunLoadTest_AllFailures(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	totalRequests := 10
	concurrency := 5

	result := RunLoadTest(server.URL, totalRequests, concurrency)

	assert.Equal(t, totalRequests, result.TotalRequests, "TotalRequests mismatch")
	assert.Equal(t, 0, result.SuccessfulRequests, "SuccessfulRequests mismatch")

	count, ok := result.StatusCodes.Load(http.StatusInternalServerError)
	assert.True(t, ok, "Status code 500 not found")
	assert.Equal(t, totalRequests, count.(int), "Status code 500 count mismatch")
}
