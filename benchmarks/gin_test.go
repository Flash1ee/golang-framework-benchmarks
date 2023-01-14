package benchmarks

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
)

func BenchmarkGinSimple(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	benchRequest(b, GetHandler("gin"), req)
}

func BenchmarkGinParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/param/abc", nil)
	benchRequest(b, GetHandler("gin"), req)
}

func BenchmarkGinPostData(b *testing.B) {
	jsonData, err := json.Marshal(genReqStrings(10))
	if err != nil {
		b.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	benchRequest(b, GetHandler("gin"), req)
}
