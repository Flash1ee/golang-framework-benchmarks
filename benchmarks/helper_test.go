package benchmarks_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"benchmarks"
)

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := httptest.NewRecorder()
	u := r.URL
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, r)

		r.Form = nil
		r.PostForm = nil
		r.MultipartForm = nil
	}
}

func genReqStrings(len int) benchmarks.Request[[]string] {
	data := make([]string, len)

	for i := 0; i < len; i++ {
		data[i] = fmt.Sprintf("data_%d", i)
	}

	return benchmarks.Request[[]string]{Data: data}
}
