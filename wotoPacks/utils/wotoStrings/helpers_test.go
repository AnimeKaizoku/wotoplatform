/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
 * Copyright (c) 2021 AmanoTeam.
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

package wotoStrings_test

import (
	"testing"
	"wp-server/wotoPacks/utils/wotoStrings"
)

func TestSplit(t *testing.T) {
	myStr := "Hello! How Are you??\u1201I'm good, what about you?"
	re := wotoStrings.Split(myStr, "!", "??", ",")
	if len(re) != 4 {
		t.Errorf("expected length of 4, but got %d", len(re))
	}

	re = wotoStrings.SplitWhite(myStr, "!", "??", ",", "How Are you")
	if len(re) != 3 {
		t.Errorf("expected length of 3, but got %d", len(re))
	}
}
