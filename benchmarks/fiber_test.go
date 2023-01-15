package benchmarks_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/goccy/go-json"

	"benchmarks"
)

func BenchmarkFiberSimple(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	benchRequest(b, benchmarks.ToHTTPAdaptor(benchmarks.FiberApp, b), req)
}

func BenchmarkFiberParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/param/abc", nil)
	benchRequest(b, benchmarks.ToHTTPAdaptor(benchmarks.FiberApp, b), req)
}

func BenchmarkFiberPostData(b *testing.B) {
	jsonData, err := json.Marshal(genReqStrings(10))
	if err != nil {
		b.Fatal(err)
	}
	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	benchRequest(b, benchmarks.ToHTTPAdaptor(benchmarks.FiberApp, b), req)
}
