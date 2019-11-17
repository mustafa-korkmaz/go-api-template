package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/middleware/secure"
)

// Config represents server specific config
type Config struct {
	Port                string
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	Debug               bool
}

// CustomBinder struct
type CustomBinder struct {
	b echo.Binder
}

// CustomValidator holds custom validator
type CustomValidator struct {
	V *validator.Validate
}

// New instantates new Echo server
func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover(),
		secure.CORS(), secure.Headers())
	e.GET("/", welcome)
	e.Validator = &CustomValidator{V: validator.New()}
	custErr := &customErrHandler{e: e}
	e.HTTPErrorHandler = custErr.handler
	e.Binder = &CustomBinder{b: &echo.DefaultBinder{}}
	return e
}

func welcome(c echo.Context) error {

	swaggerURL := "/assets/swaggerui"

	welcomeMessage := "Welcome to mute's Golang Api boilerplate project!\n"
	welcomeMessage += "Jump into api docs => " + swaggerURL
	return c.String(http.StatusOK, welcomeMessage)
}

// Start starts echo server
func Start(e *echo.Echo, cfg *Config) {
	s := &http.Server{
		Addr:         cfg.Port,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}
	e.Debug = cfg.Debug

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// NewBinder initializes custom server binder
func NewBinder() *CustomBinder {
	return &CustomBinder{b: &echo.DefaultBinder{}}
}

// Bind tries to bind request into interface, and if it does then validate it
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	if err := cb.b.Bind(i, c); err != nil && err != echo.ErrUnsupportedMediaType {
		return err
	}
	return c.Validate(i)
}

// Validate validates the request
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.V.Struct(i)
}
