package benchmarks

import (
	"net/http"
	"net/http/httptest"
	"testing"
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
