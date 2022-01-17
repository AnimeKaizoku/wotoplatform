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

import wv "wp-server/wotoPacks/core/wotoValues"

//---------------------------------------------------------

func (b *ChangeBioData) IsBioTooLong() bool {
	return wv.IsBioTooLong(b.Bio)
}

func (b *ChangeBioData) HasNotModified(info *wv.UserInfo) bool {
	return b.Bio == info.Bio
}

//---------------------------------------------------------

func (n *ChangeNamesData) IsFirstNameTooLong() bool {
	return wv.IsFirstNameTooLong(n.FirstName)
}

func (n *ChangeNamesData) IsLastNameTooLong() bool {
	return wv.IsLastNameTooLong(n.LastName)
}

func (n *ChangeNamesData) HasNotModified(info *wv.UserInfo) bool {
	return n.FirstName == info.FirstName && n.LastName == info.LastName
}

//---------------------------------------------------------
func (i *GetUserInfoData) IsInvalid() bool {
	return !wv.IsCorrectUsernameFormat(i.Username) && i.UserId.IsZero()
}
