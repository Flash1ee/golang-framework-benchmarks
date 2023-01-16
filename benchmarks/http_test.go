package benchmarks_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"main/benchmarks"
)

func BenchmarkHttpSimple(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	benchRequest(b, benchmarks.GetHandler("http"), req)
}

func BenchmarkHttpParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodPost, "/param?name=abc", nil)
	benchRequest(b, benchmarks.GetHandler("http"), req)
}

func BenchmarkHttpPostData(b *testing.B) {
	jsonData, err := json.Marshal(genReqStrings(50))
	if err != nil {
		b.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	benchRequest(b, benchmarks.GetHandler("http"), req)
}
