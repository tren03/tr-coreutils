package boot

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
	ih "github.com/tren03/tr-coreutils/tr-todo/internal/http"
	"github.com/tren03/tr-coreutils/tr-todo/internal/todo"
	"github.com/tren03/tr-coreutils/tr-todo/pkg/config"
)

type App struct {
	Config   *config.Config
	DB       *sql.DB
	Logger   *slog.Logger
	Repos    *Repos
	Clients  *Clients
	Services *Services
	Server   *ih.Router
}

type Repos struct {
	Todo todo.IRepo // Each repo interface goes here
}

type Clients struct {
	// Future client interfaces go here
	// Email email.IClient
	// SMS sms.IClient
}

type Services struct {
	// Services
}

type Option func(*App) error

func WithLogger(level string) Option {
	return func(app *App) error {
		logger, err := setupLogger(level)
		if err != nil {
			return fmt.Errorf("failed to setup logger: %w", err)
		}
		app.Logger = logger
		app.Logger.Info("logger initialized", "level", level)
		return nil
	}
}

func WithDatabase(cfg config.DB) Option {
	return func(app *App) error {
		db, err := setupDatabase(cfg)
		if err != nil {
			return fmt.Errorf("failed to setup database: %w", err)
		}
		app.DB = db
		if err := db.Ping(); err != nil {
			return fmt.Errorf("database health check failed: %w", err)
		}
		if app.Logger != nil {
			app.Logger.Info("database connected and healthy")
		}
		return nil
	}
}

func WithConfig(cfg *config.Config) Option {
	return func(app *App) error {
		app.Config = cfg
		return nil
	}
}

func WithHTTPServer() Option {
	return func(app *App) error {
		if app.Logger == nil {
			return fmt.Errorf("logger must be initialized before HTTP server")
		}
		app.Server = ih.New(app.Logger)
		app.Logger.Info("HTTP router initialized")
		return nil
	}
}

func WithRepos(cfg *config.Config) Option {
	return func(app *App) error {
		if app.DB == nil {
			return fmt.Errorf("database must be initialized before repos")
		}
		if app.Logger == nil {
			return fmt.Errorf("logger must be initialized before repos")
		}

		app.Repos = &Repos{
			Todo: todo.NewRepo(app.DB, app.Logger),
		}

		app.Logger.Info("repositories initialized")
		return nil
	}
}

func WithClients(cfg *config.Config) Option {
	return func(app *App) error {
		if app.DB == nil {
			return fmt.Errorf("database must be initialized before repos")
		}
		if app.Logger == nil {
			return fmt.Errorf("logger must be initialized before repos")
		}

		app.Clients = &Clients{
			// Initialize clients here
		}

		app.Logger.Info("clients initialized")
		return nil
	}
}

func WithServices(cfg *config.Config) Option {
	return func(app *App) error {
		if app.Logger == nil {
			return fmt.Errorf("logger must be initialized before Services")
		}
		if app.Repos == nil {
			return fmt.Errorf("Repos must be initialized before services")
		}
		if app.Clients == nil {
			return fmt.Errorf("Clients must be initialized before services")
		}

		app.Services = &Services{
			// Initialize services here
		}

		app.Logger.Info("Services initialized")
		return nil
	}
}

func helper(opts ...Option) (*App, error) {
	app := &App{}
	for _, opt := range opts {
		if err := opt(app); err != nil {
			return nil, err
		}
	}
	if app.Logger != nil {
		app.Logger.Info("application initialized successfully")
	}
	return app, nil
}

func Initalize() (*App, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("Error fetching config")
	}

	app, err := helper(
		WithConfig(cfg),
		WithLogger("info"),
		WithDatabase(*cfg.DB),
		WithRepos(cfg),
		WithClients(cfg),
		WithServices(cfg),
		WithHTTPServer(),
	)

	if err != nil {
		return nil, fmt.Errorf("Error initlizing application")
	}
	return app, nil
}

func (app *App) Shutdown(ctx context.Context) error {
	if app.Logger != nil {
		app.Logger.Info("shutting down application")
	}

	var errs []error

	if app.Server != nil {
		if err := app.Server.Shutdown(ctx); err != nil {
			errs = append(errs, fmt.Errorf("server shutdown failed: %w", err))
		}
	}

	// Close database connection
	if app.DB != nil {
		if err := app.DB.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close database: %w", err))
		} else if app.Logger != nil {
			app.Logger.Info("database connection closed")
		}
	}

	// Close any other resources (future clients, services, etc.)

	if len(errs) > 0 {
		return fmt.Errorf("shutdown errors: %v", errs)
	}

	if app.Logger != nil {
		app.Logger.Info("application shutdown complete")
	}
	return nil
}
