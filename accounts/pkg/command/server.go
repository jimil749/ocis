package command

import (
	"context"
	"strings"
	"time"

	"github.com/owncloud/ocis/ocis-pkg/sync"

	"contrib.go.opencensus.io/exporter/jaeger"
	"contrib.go.opencensus.io/exporter/ocagent"
	"contrib.go.opencensus.io/exporter/zipkin"
	"github.com/micro/cli/v2"
	"github.com/oklog/run"
	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/owncloud/ocis/accounts/pkg/config"
	"github.com/owncloud/ocis/accounts/pkg/flagset"
	"github.com/owncloud/ocis/accounts/pkg/metrics"
	"github.com/owncloud/ocis/accounts/pkg/server/grpc"
	"github.com/owncloud/ocis/accounts/pkg/server/http"
	svc "github.com/owncloud/ocis/accounts/pkg/service/v0"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

// Server is the entry point for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:        "server",
		Usage:       "Start ocis accounts service",
		Description: "uses an LDAP server as the storage backend",
		Flags:       flagset.ServerWithConfig(cfg),
		Before: func(ctx *cli.Context) error {
			logger := NewLogger(cfg)
			if cfg.HTTP.Root != "/" {
				cfg.HTTP.Root = strings.TrimSuffix(cfg.HTTP.Root, "/")
			}

			// When running on single binary mode the before hook from the root command won't get called. We manually
			// call this before hook from ocis command, so the configuration can be loaded.
			if !cfg.Supervised {
				return ParseConfig(ctx, cfg)
			}
			logger.Debug().Str("service", "accounts").Msg("ignoring config file parsing when running supervised")
			return nil
		},
		Action: func(c *cli.Context) error {
			logger := NewLogger(cfg)
			if cfg.Tracing.Enabled {
				switch t := cfg.Tracing.Type; t {
				case "agent":
					exporter, err := ocagent.NewExporter(
						ocagent.WithReconnectionPeriod(5*time.Second),
						ocagent.WithAddress(cfg.Tracing.Endpoint),
						ocagent.WithServiceName(cfg.Tracing.Service),
					)
					if err != nil {
						logger.Error().
							Err(err).
							Str("endpoint", cfg.Tracing.Endpoint).
							Str("collector", cfg.Tracing.Collector).
							Msg("Failed to create agent tracing")
						return err
					}
					trace.RegisterExporter(exporter)
					view.RegisterExporter(exporter)
				case "jaeger":
					exporter, err := jaeger.NewExporter(
						jaeger.Options{
							AgentEndpoint:     cfg.Tracing.Endpoint,
							CollectorEndpoint: cfg.Tracing.Collector,
							Process: jaeger.Process{
								ServiceName: cfg.Tracing.Service,
							},
						},
					)
					if err != nil {
						logger.Error().
							Err(err).
							Str("endpoint", cfg.Tracing.Endpoint).
							Str("collector", cfg.Tracing.Collector).
							Msg("Failed to create jaeger tracing")
						return err
					}
					trace.RegisterExporter(exporter)
				case "zipkin":
					endpoint, err := openzipkin.NewEndpoint(
						cfg.Tracing.Service,
						cfg.Tracing.Endpoint,
					)
					if err != nil {
						logger.Error().
							Err(err).
							Str("endpoint", cfg.Tracing.Endpoint).
							Str("collector", cfg.Tracing.Collector).
							Msg("Failed to create zipkin tracing")
						return err
					}
					exporter := zipkin.NewExporter(
						zipkinhttp.NewReporter(
							cfg.Tracing.Collector,
						),
						endpoint,
					)
					trace.RegisterExporter(exporter)
				default:
					logger.Warn().
						Str("type", t).
						Msg("Unknown tracing backend")
				}
				trace.ApplyConfig(
					trace.Config{
						DefaultSampler: trace.AlwaysSample(),
					},
				)
			} else {
				logger.Debug().
					Msg("Tracing is not enabled")
			}
			var (
				gr          = run.Group{}
				ctx, cancel = func() (context.Context, context.CancelFunc) {
					if cfg.Context == nil {
						return context.WithCancel(context.Background())
					}
					return context.WithCancel(cfg.Context)
				}()
				mtrcs = metrics.New()
			)

			defer cancel()

			mtrcs.BuildInfo.WithLabelValues(cfg.Server.Version).Set(1)

			handler, err := svc.New(svc.Logger(logger), svc.Config(cfg))
			if err != nil {
				logger.Error().Err(err).Msg("handler init")
				return err
			}

			{
				server := http.Server(
					http.Logger(logger),
					http.Name(cfg.Server.Name),
					http.Context(ctx),
					http.Config(cfg),
					http.Metrics(mtrcs),
					http.Handler(handler),
				)

				gr.Add(server.Run, func(_ error) {
					logger.Info().Str("server", "http").Msg("Shutting down server")
					cancel()
				})
			}

			{
				server := grpc.Server(
					grpc.Logger(logger),
					grpc.Name(cfg.Server.Name),
					grpc.Context(ctx),
					grpc.Config(cfg),
					grpc.Metrics(mtrcs),
					grpc.Handler(handler),
				)

				gr.Add(server.Run, func(_ error) {
					logger.Info().Str("server", "grpc").Msg("Shutting down server")
					cancel()
				})
			}

			if !cfg.Supervised {
				sync.Trap(&gr, cancel)
			}

			return gr.Run()
		},
	}
}