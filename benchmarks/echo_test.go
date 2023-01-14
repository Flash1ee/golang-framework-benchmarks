package benchmarks

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
)

func BenchmarkEchoSimple(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	benchRequest(b, GetHandler("echo"), req)
}

func BenchmarkEchoParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/param/abc", nil)
	benchRequest(b, GetHandler("echo"), req)
}

func BenchmarkEchoPostData(b *testing.B) {
	jsonData, err := json.Marshal(genReqStrings(10))
	if err != nil {
		b.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	benchRequest(b, GetHandler("echo"), req)
}
