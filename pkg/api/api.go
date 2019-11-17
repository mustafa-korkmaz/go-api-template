package api

// API Docs Go api boilerplate v1
//
//     Schemes: http
//     Version: 2.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Mustafa Korkmaz <mustafakorkmazdev@gmail.com>
//     Host: localhost:1907
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer: []
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/config"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {
	// db, err := postgres.New(cfg.Db.Conn, cfg.Db.Timeout, cfg.Db.LogQueries)
	// if err != nil {
	// 	return err
	// }

	// sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	// rbac := rbac.New()
	// jwt := jwt.New(cfg.Jwt.Secret, cfg.Jwt.SigningAlgorithm, cfg.Jwt.Duration)
	//log := zlog.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	//at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())

	// v1 := e.Group("/v1")
	// v1.Use(jwt.MWFunc())

	// ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	// pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
