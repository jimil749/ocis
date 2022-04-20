package command

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path"
	"strings"

	"github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/ocis/pkg/register"
	cli "github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"

	accounts "github.com/owncloud/ocis/extensions/accounts/pkg/config"
	graph "github.com/owncloud/ocis/extensions/graph/pkg/config"
	idm "github.com/owncloud/ocis/extensions/idm/pkg/config"
	notifications "github.com/owncloud/ocis/extensions/notifications/pkg/config"
	ocs "github.com/owncloud/ocis/extensions/ocs/pkg/config"
	proxy "github.com/owncloud/ocis/extensions/proxy/pkg/config"
	settings "github.com/owncloud/ocis/extensions/settings/pkg/config"
	thumbnails "github.com/owncloud/ocis/extensions/thumbnails/pkg/config"
)

const configFilename string = "ocis.yml"
const passwordLength int = 32

// InitCommand is the entrypoint for the init command
func InitCommand(cfg *config.Config) *cli.Command {
	// TODO: remove homedir get
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("could not get homedir")
	}
	return &cli.Command{
		Name:  "init",
		Usage: "initialise an ocis config",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "insecure",
				EnvVars: []string{"OCIS_INSECURE"},
				Value:   "ask",
			},
			&cli.BoolFlag{
				Name:    "force-overwrite",
				Aliases: []string{"f"},
				EnvVars: []string{"OCIS_FORCE_CONFIG_OVERWRITE"},
				Value:   false,
			},
			&cli.StringFlag{
				Name: "config-path",
				//Value: cfg.ConfigPath, // TODO: as soon as PR 3480 is merged, remove quotes
				Value: path.Join(homeDir, ".ocis"), // TODO: this is temporary for experimenting, line above is relevant
				Usage: "config path for the ocis runtime",
				// Destination: &cfg.ConfigFile, // TODO: same as above
			},
		},
		Action: func(c *cli.Context) error {
			insecureFlag := c.String("insecure")
			insecure := false
			if insecureFlag == "ask" {
				answer := strings.ToLower(stringPrompt("Insecure Backends? [Yes|No]"))
				if answer == "yes" || answer == "y" {
					insecure = true
				}
			} else if insecureFlag == "true" {
				insecure = true
			}
			err := createConfig(insecure, c.Bool("force-overwrite"), c.String("config-path"))
			if err != nil {
				log.Fatalf("Could not create config: %s", err)
			}
			return nil
		},
	}
}

func init() {
	register.AddCommand(InitCommand)
}

func checkConfigPath(configPath string) error {
	targetPath := path.Join(configPath, configFilename)
	if _, err := os.Stat(targetPath); err == nil {
		return fmt.Errorf("Config in %s already exists", targetPath)
	}
	return nil
}

func createConfig(insecure, forceOverwrite bool, configPath string) error {
	err := checkConfigPath(configPath)
	if err != nil && !forceOverwrite {
		return err
	}
	err = os.MkdirAll(configPath, 0700)
	if err != nil {
		return err
	}
	cfg := config.Config{
		Accounts: &accounts.Config{},
		//Audit:    &audit.Config{},
		//GLAuth:        &glauth.Config{},
		//GraphExplorer: &graphExplorer.Config{},
		Graph: &graph.Config{},
		IDM:   &idm.Config{},
		//IDP:           &idp.Config{},
		//Nats:          &nats.Config{},
		Notifications: &notifications.Config{},
		Proxy:         &proxy.Config{},
		OCS:           &ocs.Config{},
		Settings:      &settings.Config{},
		//Storage:       &storage.Config{},
		Thumbnails: &thumbnails.Config{},
		//Web:           &web.Config{},
		//WebDAV:        &webdav.Config{},
	}

	if insecure {
		cfg.Proxy.InsecureBackends = insecure
	}

	idmServicePassword, err := generateRandomPassword(passwordLength)
	if err != nil {
		return fmt.Errorf("Could not generate random password for idm: %s", err)
	}
	idpServicePassword, err := generateRandomPassword(passwordLength)
	if err != nil {
		return fmt.Errorf("Could not generate random password for idp: %s", err)
	}
	ocisAdminServicePassword, err := generateRandomPassword(passwordLength)
	if err != nil {
		return fmt.Errorf("Could not generate random password for ocis admin: %s", err)
	}
	revaServicePassword, err := generateRandomPassword(passwordLength)
	if err != nil {
		return fmt.Errorf("Could not generate random password for reva: %s", err)
	}
	tokenManagerJwtSecret, err := generateRandomPassword(passwordLength)
	if err != nil {
		return fmt.Errorf("Could not generate random password for tokenmanager: %s", err)
	}
	machineAuthSecret, err := generateRandomPassword(passwordLength)
	if err != nil {
		return fmt.Errorf("Could not generate random password for machineauthsecret: %s", err)
	}
	thumbnailTransferTokenSecret, err := generateRandomPassword(passwordLength)
	if err != nil {
		return fmt.Errorf("Could not generate random password for machineauthsecret: %s", err)
	}

	cfg.TokenManager.JWTSecret = tokenManagerJwtSecret
	cfg.Accounts.TokenManager.JWTSecret = tokenManagerJwtSecret
	cfg.Graph.TokenManager.JWTSecret = tokenManagerJwtSecret
	cfg.IDM.ServiceUserPasswords.Idm = idmServicePassword
	cfg.IDM.ServiceUserPasswords.Idp = idpServicePassword
	cfg.IDM.ServiceUserPasswords.OcisAdmin = ocisAdminServicePassword
	cfg.IDM.ServiceUserPasswords.Reva = revaServicePassword
	cfg.Notifications.Notifications.MachineAuthSecret = machineAuthSecret
	cfg.OCS.MachineAuthAPIKey = machineAuthSecret
	cfg.Proxy.TokenManager.JWTSecret = tokenManagerJwtSecret
	cfg.Proxy.MachineAuthAPIKey = machineAuthSecret
	cfg.Settings.Metadata.MachineAuthAPIKey = machineAuthSecret
	cfg.Settings.TokenManager.JWTSecret = tokenManagerJwtSecret
	cfg.Thumbnails.Thumbnail.TransferTokenSecret = thumbnailTransferTokenSecret
	yamlOutput, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Could not marshall config into yaml: %s", err)
	}
	targetPath := path.Join(configPath, configFilename)
	err = ioutil.WriteFile(targetPath, yamlOutput, 0600)
	if err != nil {
		return err
	}
	fmt.Printf(
		"======================================\n"+
			" generated OCIS Config\n"+
			"======================================\n"+
			" configpath : %s\n"+
			" user       : admin\n"+
			" password   : %s\n",
		targetPath, ocisAdminServicePassword)
	return nil
}

func stringPrompt(label string) string {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		input, _ = reader.ReadString('\n')
		if input != "" {
			break
		}
	}
	return strings.TrimSpace(input)
}

func generateRandomPassword(length int) (string, error) {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-=+!@#$%^&*."
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		ret[i] = chars[num.Int64()]
	}

	return string(ret), nil
}