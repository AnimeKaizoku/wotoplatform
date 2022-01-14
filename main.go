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

package main

import (
	"log"
	"net"

	"wp-server/wotoPacks/core/utils/logging"
	"wp-server/wotoPacks/core/wotoConfig"
	"wp-server/wotoPacks/database"
	"wp-server/wotoPacks/entryPoints"

	"go.uber.org/zap"
)

func main() {
	f := loadLogger()
	if f != nil {
		defer f()
	}

	err := runServer()
	if err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	cfg, err := wotoConfig.GetConfig()
	if err != nil {
		return err
	}

	whole := cfg.Bind + ":" + cfg.Port
	ln, err := net.Listen(cfg.Network, whole)
	if err != nil {
		return err
	}

	// do NOT close the listener in this function.
	// it should be done in `entryPoints.Listen` function.
	//defer ln.Close()

	err = database.StartDatabase()
	if err != nil {
		return err
	}

	entryPoints.Listen(ln)

	return nil
}

func loadLogger() func() {
	loggerMgr := logging.InitZapLog(true)
	zap.ReplaceGlobals(loggerMgr)
	logging.SUGARED = loggerMgr.Sugar()
	return func() {
		_ = loggerMgr.Sync()
	}
}
