package main

import (
	"net/http"
	"net/url"
	"os"

	"movie-search/gateway"
	"movie-search/infrastructure"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "gateway",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	route := chi.NewRouter()
	route.Use(middleware.Logger)
	r := gateway.InjectRoutes()
	r.RegisterRoutes(route)

	level.Info(logger).Log("msg", "gateway started")
	defer level.Info(logger).Log("msg", "gateway ended")
	level.Info(logger).Log("msg", "Server ready at 8080")
	level.Info(logger).Log(http.ListenAndServe(":8080", route))
}

func init() {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	env := infrastructure.Environment{}
	env.SetEnvironment()
	env.LoadConfig()
}
