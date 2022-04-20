package config

import (
	"context"

	"github.com/owncloud/ocis/ocis-pkg/shared"
)

// Config combines all available configuration parts.
type Config struct {
	*shared.Commons `yaml:"-"`

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing,omitempty"`
	Log     *Log     `yaml:"log,omitempty"`
	Debug   Debug    `yaml:"debug,omitempty"`

	HTTP HTTP `yaml:"http,omitempty"`

	TokenManager TokenManager `yaml:"token_manager,omitempty"`
	Reva         Reva         `yaml:"reva,omitempty"`

	IdentityManagement IdentityManagement `yaml:"identity_management,omitempty"`

	AccountBackend     string `yaml:"account_backend,omitempty" env:"OCS_ACCOUNT_BACKEND_TYPE"`
	StorageUsersDriver string `yaml:"storage_users_driver,omitempty" env:"STORAGE_USERS_DRIVER;OCS_STORAGE_USERS_DRIVER"`
	MachineAuthAPIKey  string `yaml:"machine_auth_api_key,omitempty" env:"OCIS_MACHINE_AUTH_API_KEY;OCS_MACHINE_AUTH_API_KEY"`

	Context context.Context `yaml:"-"`
}

// IdentityManagement keeps track of the OIDC address. This is because Reva requisite of uniqueness for users
// is based in the combination of IDP hostname + UserID. For more information see:
// https://github.com/cs3org/reva/blob/4fd0229f13fae5bc9684556a82dbbd0eced65ef9/pkg/storage/utils/decomposedfs/node/node.go#L856-L865
type IdentityManagement struct {
	Address string `yaml:"address" env:"OCIS_URL;OCS_IDM_ADDRESS"`
}