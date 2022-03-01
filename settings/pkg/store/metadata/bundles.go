// Package store implements the go-micro store interface
package store

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	settingsmsg "github.com/owncloud/ocis/protogen/gen/ocis/messages/settings/v0"
	"github.com/owncloud/ocis/settings/pkg/store/defaults"
)

// ListBundles returns all bundles in the dataPath folder that match the given type.
func (s *Store) ListBundles(bundleType settingsmsg.Bundle_Type, bundleIDs []string) ([]*settingsmsg.Bundle, error) {
	// TODO: this is needed for initialization - we need to find a better way to fix this
	if s.mdc == nil && len(bundleIDs) == 1 {
		return defaultBundle(bundleType, bundleIDs[0]), nil
	}
	s.Init()

	if len(bundleIDs) == 0 {
		bIDs, err := s.mdc.ReadDir(nil, bundleFolderLocation)
		if err != nil {
			return nil, err
		}

		bundleIDs = bIDs
	}
	var bundles []*settingsmsg.Bundle
	for _, id := range bundleIDs {
		b, err := s.mdc.SimpleDownload(nil, bundlePath(id))
		if err != nil {
			return nil, err
		}

		bundle := &settingsmsg.Bundle{}
		err = json.Unmarshal(b, bundle)
		if err != nil {
			return nil, err
		}

		if bundle.Type == bundleType {
			bundles = append(bundles, bundle)
		}

	}
	return bundles, nil
}

// ReadBundle tries to find a bundle by the given id within the dataPath.
func (s *Store) ReadBundle(bundleID string) (*settingsmsg.Bundle, error) {
	s.Init()
	b, err := s.mdc.SimpleDownload(nil, bundlePath(bundleID))
	if err != nil {
		return nil, err
	}

	bundle := &settingsmsg.Bundle{}
	return bundle, json.Unmarshal(b, bundle)
}

// ReadSetting tries to find a setting by the given id within the dataPath.
func (s *Store) ReadSetting(settingID string) (*settingsmsg.Setting, error) {
	fmt.Println("ReadSetting not implemented")
	return nil, errors.New("not implemented")
}

// WriteBundle sends the givens record to the metadataclient. returns `record` for legacy reasons
func (s *Store) WriteBundle(record *settingsmsg.Bundle) (*settingsmsg.Bundle, error) {
	s.Init()
	b, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	return record, s.mdc.SimpleUpload(nil, bundlePath(record.Id), b)
}

// AddSettingToBundle adds the given setting to the bundle with the given bundleID.
func (s *Store) AddSettingToBundle(bundleID string, setting *settingsmsg.Setting) (*settingsmsg.Setting, error) {
	s.Init()
	b, err := s.ReadBundle(bundleID)
	if err != nil {
		// TODO: How to differentiate 'not found'?
		b = new(settingsmsg.Bundle)
		b.Id = bundleID
		b.Type = settingsmsg.Bundle_TYPE_DEFAULT
	}

	if setting.Id == "" {
		setting.Id = uuid.Must(uuid.NewV4()).String()
	}

	b.Settings = append(b.Settings, setting)
	_, err = s.WriteBundle(b)
	return setting, err
}

// RemoveSettingFromBundle removes the setting from the bundle with the given ids.
func (s *Store) RemoveSettingFromBundle(bundleID string, settingID string) error {
	fmt.Println("RemoveSettingFromBundle not implemented")
	return errors.New("not implemented")
}

func bundlePath(id string) string {
	return fmt.Sprintf("%s/%s", bundleFolderLocation, id)
}

func defaultBundle(bundleType settingsmsg.Bundle_Type, bundleID string) []*settingsmsg.Bundle {
	var bundles []*settingsmsg.Bundle
	for _, b := range defaults.GenerateBundlesDefaultRoles() {
		if b.Type == bundleType && b.Id == bundleID {
			bundles = append(bundles, b)
		}
	}
	return bundles
}
