package benchmarks

import (
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Request[T any] struct {
	Data T
}

var (
	httpHandlers map[string]http.Handler
	mu           = sync.RWMutex{}
	FiberApp     *fiber.App
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
