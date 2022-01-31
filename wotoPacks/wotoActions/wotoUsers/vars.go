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

package wotoUsers

import (
	wv "wp-server/wotoPacks/core/wotoValues"
	wa "wp-server/wotoPacks/wotoActions"
)

var (
	_batchHandlers = map[wa.BatchExecution]wv.ReqHandler{
		BATCH_REGISTER_USER:           batchRegisterUser,
		BATCH_REGISTER_VIRTUAL_USER:   batchRegisterVirtualUser,
		BATCH_LOGIN_USER:              batchLoginUser,
		BATCH_GET_ME:                  batchGetMe,
		BATCH_CHANGE_USER_BIO:         batchChangeUserBio,
		BATCH_CHANGE_NAMES:            batchChangeNames,
		BATCH_GET_USER_INFO:           batchGetUserInfo,
		BATCH_GET_USER_BY_TELEGRAM_ID: batchGetUserByTelegramID,
		BATCH_GET_USER_BY_EMAIL:       batchGetUserByEmail,
		BATCH_RESOLVE_USERNAME:        batchResolveUsername,
		BATCH_CHANGE_USER_PERMISSION:  batchChangeUserPermission,
		BATCH_GET_USER_FAVORITE:       batchGetUserFavorite,
		BATCH_GET_USER_FAVORITE_COUNT: batchGetUserFavoriteCount,
		BATCH_SET_USER_FAVORITE:       batchSetUserFavorite,
	}
)
