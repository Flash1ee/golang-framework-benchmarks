package benchmarks

import (
	"fmt"
	"net/http"
	"sync"
)

type Request[T any] struct {
	Data T
}

var (
	httpHandlers map[string]http.Handler
	mu           = sync.RWMutex{}
)

func RegisterHandler(name string, handler http.Handler) {
	if httpHandlers == nil {
		httpHandlers = make(map[string]http.Handler)
	}
	if _, ok := httpHandlers[name]; ok {
		panic("already registered")
	}
	mu.Lock()
	httpHandlers[name] = handler
	mu.Unlock()
}

func GetHandler(name string) http.Handler {
	if httpHandlers == nil {
		return nil
	}
	mu.RLock()
	handler, _ := httpHandlers[name]
	mu.RUnlock()
	return handler
}

func genReqStrings(len int) Request[[]string] {
	data := make([]string, len)

	for i := 0; i < len; i++ {
		data[i] = fmt.Sprintf("data_%d", i)
	}

	return Request[[]string]{Data: data}
}
