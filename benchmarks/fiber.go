package benchmarks

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

var Handler fasthttp.RequestHandler

func init() {
	h := fiber.New()
	h.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})
	h.Post("/", func(c *fiber.Ctx) error {
		var req Request[[]string]
		if err := c.BodyParser(&req); err != nil {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return c.Send([]byte(err.Error()))
		}

		if len(req.Data) == 0 {
			c.Response().SetStatusCode(http.StatusBadRequest)
			return nil
		}

		c.JSON(req.Data[len(req.Data)-1])
		return nil
	})
	h.Get("/param/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString(fmt.Sprintf("Hello, %s", name))
	})
	FiberApp = h
}

func ToHTTPAdaptor(app *fiber.App, b *testing.B) http.HandlerFunc {
	return handlerFunc(app, b)
}

func handlerFunc(app *fiber.App, b *testing.B) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// New fasthttp request
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		// Convert net/http -> fasthttp request
		if r.Body != nil {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, utils.StatusMessage(fiber.StatusInternalServerError), fiber.StatusInternalServerError)
				return
			}
			req.Header.SetContentLength(len(body))
			_, _ = req.BodyWriter().Write(body)
		}
		req.Header.SetMethod(r.Method)
		req.SetRequestURI(r.RequestURI)
		req.SetHost(r.Host)
		for key, val := range r.Header {
			for _, v := range val {
				req.Header.Set(key, v)
			}
		}
		if _, _, err := net.SplitHostPort(r.RemoteAddr); err != nil && err.(*net.AddrError).Err == "missing port in address" {
			r.RemoteAddr = net.JoinHostPort(r.RemoteAddr, "80")
		}
		remoteAddr, err := net.ResolveTCPAddr("tcp", r.RemoteAddr)
		if err != nil {
			http.Error(w, utils.StatusMessage(fiber.StatusInternalServerError), fiber.StatusInternalServerError)
			return
		}

		// New fasthttp Ctx
		var fctx fasthttp.RequestCtx
		fctx.Init(req, remoteAddr, nil)

		// Execute fasthttp Ctx though app.Handler
		app.Handler()(&fctx)

		// Convert fasthttp Ctx > net/http
		fctx.Response.Header.VisitAll(func(k, v []byte) {
			w.Header().Add(string(k), string(v))
		})
		w.WriteHeader(fctx.Response.StatusCode())
		_, _ = w.Write(fctx.Response.Body())
	}
}
