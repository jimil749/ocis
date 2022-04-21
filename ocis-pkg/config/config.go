package config

import (
	"github.com/owncloud/ocis/ocis-pkg/shared"

	accounts "github.com/owncloud/ocis/extensions/accounts/pkg/config"
	appprovider "github.com/owncloud/ocis/extensions/appprovider/pkg/config"
	audit "github.com/owncloud/ocis/extensions/audit/pkg/config"
	authbasic "github.com/owncloud/ocis/extensions/auth-basic/pkg/config"
	authbearer "github.com/owncloud/ocis/extensions/auth-bearer/pkg/config"
	authmachine "github.com/owncloud/ocis/extensions/auth-machine/pkg/config"
	glauth "github.com/owncloud/ocis/extensions/glauth/pkg/config"
	graphExplorer "github.com/owncloud/ocis/extensions/graph-explorer/pkg/config"
	graph "github.com/owncloud/ocis/extensions/graph/pkg/config"
	group "github.com/owncloud/ocis/extensions/group/pkg/config"
	idm "github.com/owncloud/ocis/extensions/idm/pkg/config"
	idp "github.com/owncloud/ocis/extensions/idp/pkg/config"
	nats "github.com/owncloud/ocis/extensions/nats/pkg/config"
	notifications "github.com/owncloud/ocis/extensions/notifications/pkg/config"
	ocs "github.com/owncloud/ocis/extensions/ocs/pkg/config"
	proxy "github.com/owncloud/ocis/extensions/proxy/pkg/config"
	settings "github.com/owncloud/ocis/extensions/settings/pkg/config"
	sharing "github.com/owncloud/ocis/extensions/sharing/pkg/config"
	storagemetadata "github.com/owncloud/ocis/extensions/storage-metadata/pkg/config"
	storagepublic "github.com/owncloud/ocis/extensions/storage-publiclink/pkg/config"
	storageshares "github.com/owncloud/ocis/extensions/storage-shares/pkg/config"
	storageusers "github.com/owncloud/ocis/extensions/storage-users/pkg/config"
	storage "github.com/owncloud/ocis/extensions/storage/pkg/config"
	store "github.com/owncloud/ocis/extensions/store/pkg/config"
	thumbnails "github.com/owncloud/ocis/extensions/thumbnails/pkg/config"
	user "github.com/owncloud/ocis/extensions/user/pkg/config"
	web "github.com/owncloud/ocis/extensions/web/pkg/config"
	webdav "github.com/owncloud/ocis/extensions/webdav/pkg/config"
)

// TokenManager is the config for using the reva token manager
type TokenManager struct {
	JWTSecret string `yaml:"jwt_secret" env:"OCIS_JWT_SECRET"`
}

const (
	// SUPERVISED sets the runtime mode as supervised threads.
	SUPERVISED = iota

	// UNSUPERVISED sets the runtime mode as a single thread.
	UNSUPERVISED
)

type Mode int

// Runtime configures the oCIS runtime when running in supervised mode.
type Runtime struct {
	Port       string `yaml:"port" env:"OCIS_RUNTIME_PORT"`
	Host       string `yaml:"host" env:"OCIS_RUNTIME_HOST"`
	Extensions string `yaml:"extensions" env:"OCIS_RUN_EXTENSIONS"`
}

// Config combines all available configuration parts.
type Config struct {
	*shared.Commons `yaml:"shared"`

	Tracing shared.Tracing `yaml:"tracing"`
	Log     *shared.Log    `yaml:"log"`

	Mode    Mode // DEPRECATED
	File    string
	OcisURL string `yaml:"ocis_url"`

	Registry     string       `yaml:"registry"`
	TokenManager TokenManager `yaml:"token_manager"`
	Runtime      Runtime      `yaml:"runtime"`

	Audit             *audit.Config           `yaml:"audit"`
	Accounts          *accounts.Config        `yaml:"accounts"`
	GLAuth            *glauth.Config          `yaml:"glauth"`
	Graph             *graph.Config           `yaml:"graph"`
	GraphExplorer     *graphExplorer.Config   `yaml:"graph_explorer"`
	IDP               *idp.Config             `yaml:"idp"`
	IDM               *idm.Config             `yaml:"idm"`
	Nats              *nats.Config            `yaml:"nats"`
	Notifications     *notifications.Config   `yaml:"notifications"`
	OCS               *ocs.Config             `yaml:"ocs"`
	Web               *web.Config             `yaml:"web"`
	Proxy             *proxy.Config           `yaml:"proxy"`
	Settings          *settings.Config        `yaml:"settings"`
	Storage           *storage.Config         `yaml:"storage"`
	AuthBasic         *authbasic.Config       `yaml:"auth_basic"`
	AuthBearer        *authbearer.Config      `yaml:"auth_bearer"`
	AuthMachine       *authmachine.Config     `yaml:"auth_machine"`
	User              *user.Config            `yaml:"user"`
	Group             *group.Config           `yaml:"group"`
	AppProvider       *appprovider.Config     `yaml:"app_provider"`
	Sharing           *sharing.Config         `yaml:"app_provider"`
	StorageMetadata   *storagemetadata.Config `yaml:"storage_metadata"`
	StoragePublicLink *storagepublic.Config   `yaml:"storage_public"`
	StorageUsers      *storageusers.Config    `yaml:"storage_users"`
	StorageShares     *storageshares.Config   `yaml:"storage_shares"`
	Store             *store.Config           `yaml:"store"`
	Thumbnails        *thumbnails.Config      `yaml:"thumbnails"`
	WebDAV            *webdav.Config          `yaml:"webdav"`
}
