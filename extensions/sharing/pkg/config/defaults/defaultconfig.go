package defaults

import (
	"path/filepath"

	"github.com/owncloud/ocis/extensions/sharing/pkg/config"
	"github.com/owncloud/ocis/ocis-pkg/config/defaults"
)

func FullDefaultConfig() *config.Config {
	cfg := DefaultConfig()

	EnsureDefaults(cfg)

	return cfg
}

func DefaultConfig() *config.Config {
	return &config.Config{
		Debug: config.Debug{
			Addr:   "127.0.0.1:9151",
			Token:  "",
			Pprof:  false,
			Zpages: false,
		},
		GRPC: config.GRPCConfig{
			Addr:     "127.0.0.1:9150",
			Protocol: "tcp",
		},
		Service: config.Service{
			Name: "sharing",
		},
		GatewayEndpoint:   "127.0.0.1:9142",
		JWTSecret:         "Pive-Fumkiu4",
		UserSharingDriver: "json",
		UserSharingDrivers: config.UserSharingDrivers{
			JSON: config.UserSharingJSONDriver{
				File: filepath.Join(defaults.BaseDataPath(), "storage", "shares.json"),
			},
			SQL: config.UserSharingSQLDriver{
				DBUsername:                 "",
				DBPassword:                 "",
				DBHost:                     "",
				DBPort:                     1433,
				DBName:                     "",
				PasswordHashCost:           11,
				EnableExpiredSharesCleanup: true,
				JanitorRunInterval:         60,
			},
			CS3: config.UserSharingCS3Driver{
				ProviderAddr:   "127.0.0.1:9215",
				ServiceUserID:  "95cb8724-03b2-11eb-a0a6-c33ef8ef53ad",
				ServiceUserIDP: "https://localhost:9200",
			},
		},
		PublicSharingDriver: "json",
		PublicSharingDrivers: config.PublicSharingDrivers{
			JSON: config.PublicSharingJSONDriver{
				File: filepath.Join(defaults.BaseDataPath(), "storage", "publicshares.json"),
			},
			SQL: config.PublicSharingSQLDriver{
				DBUsername:                 "",
				DBPassword:                 "",
				DBHost:                     "",
				DBPort:                     1433,
				DBName:                     "",
				PasswordHashCost:           11,
				EnableExpiredSharesCleanup: true,
				JanitorRunInterval:         60,
			},
			CS3: config.PublicSharingCS3Driver{
				ProviderAddr:   "127.0.0.1:9215",
				ServiceUserID:  "95cb8724-03b2-11eb-a0a6-c33ef8ef53ad",
				ServiceUserIDP: "https://localhost:9200",
			},
		},
		Events: config.Events{
			Addr:      "127.0.0.1:9233",
			ClusterID: "ocis-cluster",
		},
	}
}

func EnsureDefaults(cfg *config.Config) {
	// provide with defaults for shared logging, since we need a valid destination address for BindEnv.
	if cfg.Logging == nil && cfg.Commons != nil && cfg.Commons.Log != nil {
		cfg.Logging = &config.Logging{
			Level:  cfg.Commons.Log.Level,
			Pretty: cfg.Commons.Log.Pretty,
			Color:  cfg.Commons.Log.Color,
			File:   cfg.Commons.Log.File,
		}
	} else if cfg.Logging == nil {
		cfg.Logging = &config.Logging{}
	}
	// provide with defaults for shared tracing, since we need a valid destination address for BindEnv.
	if cfg.Tracing == nil && cfg.Commons != nil && cfg.Commons.Tracing != nil {
		cfg.Tracing = &config.Tracing{
			Enabled:   cfg.Commons.Tracing.Enabled,
			Type:      cfg.Commons.Tracing.Type,
			Endpoint:  cfg.Commons.Tracing.Endpoint,
			Collector: cfg.Commons.Tracing.Collector,
		}
	} else if cfg.Tracing == nil {
		cfg.Tracing = &config.Tracing{}
	}
}
