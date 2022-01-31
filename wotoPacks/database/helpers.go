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

package database

import (
	"wp-server/wotoPacks/core/utils/logging"
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database/groupCallsDatabase"
	"wp-server/wotoPacks/database/mediaDatabase"
	"wp-server/wotoPacks/database/usersDatabase"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SESSION *gorm.DB

func StartDatabase() error {
	var err error
	var db *gorm.DB

	if wotoConfig.UseSqlite() {
		db, err = gorm.Open(
			sqlite.Open(wotoConfig.GetDbPath()),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Error),
			},
		)
	} else {
		url := wotoConfig.GetDatabaseURL()
		db, err = gorm.Open(
			postgres.Open(url),
			&gorm.Config{
				Logger: logger.Default.LogMode(logger.Error),
			},
		)
	}

	if err != nil {
		return err
	}

	SESSION = db
	wv.SESSION = SESSION

	logging.Info("Database connected ")

	//Create tables if they don't exist
	err = SESSION.AutoMigrate(
		usersDatabase.ModelUserInfo,
		usersDatabase.ModelUserFavorite,
	)
	if err != nil {
		return err
	}

	groupCallsDatabase.LoadGroupCallsDatabase()
	mediaDatabase.LoadMediaDatabase()
	usersDatabase.LoadUsersDatabase()

	logging.Info("Auto-migrated database schema")

	return nil
}
