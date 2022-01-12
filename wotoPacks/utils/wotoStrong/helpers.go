/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
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

package wotoStrong

import (
	wv "wp-server/wotoPacks/utils/wotoValues"
)

func repairString(value *string) *string {
	entered := false
	ignoreNext := false
	final := wv.EMPTY
	last := len(*value) - wv.BaseIndex
	next := wv.BaseIndex
	for i, current := range *value {
		if ignoreNext {
			ignoreNext = false
			continue
		}

		if current == wv.CHAR_STR {
			if !entered {
				entered = true
			} else {
				entered = false
			}

			final += string(current)
			continue
		} else {
			if !entered {
				final += string(current)
				continue
			}

			if isSpecial(current) {
				final += wv.BackSlash + string(current)
				continue
			} else {
				if current == wv.LineChar {
					if i != last {
						next = i + wv.BaseOneIndex
						if (*value)[next] == wv.LineChar {
							final += wv.BackSlash +
								string(current) + string(current)
							ignoreNext = true
							continue
						}
					}
				}
			}
		}

		final += string(current)
	}

	return &final
}

func isSpecial(r rune) bool {
	switch r {
	case wv.EqualChar, wv.DPointChar:
		return true
	default:
		return false
	}

}
