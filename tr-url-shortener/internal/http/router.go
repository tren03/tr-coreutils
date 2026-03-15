package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	mux    *chi.Mux
	logger *slog.Logger
	srv    *http.Server
}

func New(logger *slog.Logger) *Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check endpoint
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "HTTP API Ready"}`))
	})

	return &Router{
		mux:    r,
		logger: logger,
	}
}

// Group creates a new route group for easy middleware stacking and path prefixing.
func (rt *Router) Group(prefix string) *chi.Mux {
	g := chi.NewRouter()
	rt.mux.Mount(prefix, g)
	return g
}

// MountHandler adds a handler func to a specific path. Use this for DI: pass handler factories that take deps.
func (rt *Router) Mount(path, method string, handler http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		rt.mux.Get(path, handler)
	case http.MethodPost:
		rt.mux.Post(path, handler)
	// Add more methods as needed: Put, Delete, etc.
	default:
		rt.logger.Error("unsupported method", "method", method)
	}
}

// Serve starts the HTTP server on the given address.

func (rt *Router) ListenAndServe(addr string) error {
	rt.srv = &http.Server{
		Addr:    addr,
		Handler: rt.mux,
	}
	rt.logger.Info("starting HTTP server", "addr", addr)
	if err := rt.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (rt *Router) Shutdown(ctx context.Context) error {
	if rt.srv != nil {
		return rt.srv.Shutdown(ctx)
	}
	return nil
}
