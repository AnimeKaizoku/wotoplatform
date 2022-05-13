/*
 * This file is part of wp-server project (https://github.com/AnimeKaizoku/wotoplatform).
 * Copyright (c) 2021 ALiwoto.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package wotoConfig

import (
	"encoding/json"
	"io/ioutil"
	"wp-server/wotoPacks/core/utils/logging"
)

// GetConfig will create a new config from "config.json" file;
// if this file does not exist at the root of your project,
// it will use a default configuration.
func GetConfig() (*Config, error) {
	if WConfig != nil {
		return WConfig, nil
	} else {
		WConfig = new(Config)
	}

	// read config file and parse it with json
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		// there is a problem in configuration?
		// not found the proper config file?
		// continue your work with default configuration
		// (WARNIGN: usage: GitHub workflows ONLY).
		logging.Warn("got an error while tried to open config.json file: ")
		logging.Warn(err)
		logging.Warn("continuing with a default config variable.")

		return getDefaultConfig(), nil
	}

	err = json.Unmarshal(configFile, WConfig)
	if err != nil {
		return nil, err
	}

	return WConfig, nil
}

// GetConfig will create a new config from path file;
// if the file does not exist at the root of your project,
// it will use a default configuration.
func GetConfigByPath(path string) (*Config, error) {
	if WConfig != nil {
		return WConfig, nil
	} else {
		WConfig = new(Config)
	}

	// read config file and parse it with json
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		// there is a problem in configuration?
		// not found the proper config file?
		// continue your work with default configuration
		// (WARNIGN: usage: GitHub workflows ONLY).
		logging.Warn("got an error while tried to open config.json file: ")
		logging.Warn(err)
		logging.Warn("continuing with a default config variable.")

		return getDefaultConfig(), nil
	}

	err = json.Unmarshal(configFile, &WConfig)
	if err != nil {
		return nil, err
	}

	return WConfig, nil
}

// IsClientIDValid will iterate over the valid client IDs in
// the config value, if the given client ID does not exist there,
// it will return false; otherwise it will return true.
func IsClientIDValid(cli string) bool {
	if WConfig == nil {
		return false
	} else if len(WConfig.ClientIDs) == 0 {
		return WConfig.IsDefault
	}

	if len(WConfig.ClientIDs) == 1 {
		return WConfig.ClientIDs[0] == cli
	}

	for _, id := range WConfig.ClientIDs {
		if id == cli {
			return true
		}
	}

	return false
}

// getDefaultConfig will create a new config and assign it to
// `WConfig` variable.
func getDefaultConfig() *Config {
	if WConfig == nil {
		// WConfig variable should already be initialized when calling
		// this function.
		return nil
	}

	// default name of the project.
	WConfig.Name = "woto-platform"

	WConfig.Description = "online music service"

	// connection type by default is set to "tcp"
	WConfig.Network = "tcp"

	// localhost will only accepts connections from within
	// the local machine, if you want to accept public connections
	// from the outside of the local machine, you need to create a
	// proper `config.json` file and put the proper variables there.
	WConfig.Bind = "localhost"

	// the default port is 50100
	WConfig.Port = "50100"

	// use a default db name for config value
	WConfig.DatabaseName = DefaultDatabaseName

	// use SQL-Lite by default, so we don't need any sql url and
	// password.
	WConfig.UseSQLLite = true

	// WARNING: testing purpose only
	WConfig.IsDefault = true

	return WConfig
}

func UseSqlite() bool {
	if WConfig != nil {
		return WConfig.UseSQLLite
	}
	return false
}

func UseCrypto() bool {
	if WConfig != nil {
		return WConfig.UseCrypto
	}
	return false
}

func GetDbPath() string {
	if WConfig != nil {
		return WConfig.DatabaseName + ".db"
	}
	return DefaultDatabaseName + ".db"
}

func GetDatabaseURL() string {
	if WConfig == nil {
		return ""
	}
	return WConfig.DatabaseUrl
}

func SkipDefaultTransaction() bool {
	if WConfig == nil {
		return false
	}
	return WConfig.SkipDefaultTransaction
}

func IsDebug() bool {
	if WConfig == nil {
		return true
	}
	return WConfig.IsDebug || WConfig.IsDefault
}

func GetOwners() []WotoOwner {
	if WConfig == nil {
		return nil
	}
	return WConfig.Owners
}
