package api

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/auth"
	at "github.com/mustafa-korkmaz/goapitemplate/pkg/api/auth/transport"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/healthcheck"
	hct "github.com/mustafa-korkmaz/goapitemplate/pkg/api/healthcheck/transport"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/api/olive"
	ot "github.com/mustafa-korkmaz/goapitemplate/pkg/api/olive/transport"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/config"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/middleware/jwt"
	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/server"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {

	dbClient, err := mongodb.New(cfg.Db.Conn, cfg.Db.Timeout, cfg.Db.LogQueries)
	if err != nil {
		return err
	}

	// sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	// rbac := rbac.New()
	//
	//log := zlog.New()

	e := server.New()
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	//at.NewHTTP(al.New(auth.Initialize(db, jwt, sec, rbac), log), e, jwt.MWFunc())

	//create jwt token validation middleware
	jwt := jwt.New(cfg.Jwt.Secret, cfg.Jwt.SigningAlgorithm, cfg.Jwt.Duration)

	//group api versions
	v1 := e.Group("/v1")
	v2 := e.Group("/v2")

	// ut.NewHTTP(ul.New(user.Initialize(db, rbac, sec), log), v1)
	// pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	at.NewHTTP(auth.New(jwt, dbClient, cfg.Db.Name), jwt.MWFunc(), v1)
	hct.NewHTTP(healthcheck.New(), v1, v2)
	ot.NewHTTP(olive.New(dbClient, cfg.Db.Name), jwt.MWFunc(), v1)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}
