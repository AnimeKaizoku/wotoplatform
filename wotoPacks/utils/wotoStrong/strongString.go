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
	tfc "wp-server/wotoPacks/interfaces"
)

const (
	StrongOffSet = 13
)

// the StrongString used in the program for High-security!
type StrongString struct {
	_value []rune
}

// Ss will generate a new StrongString
// with the specified non-encoded string value.
func Ss(s string) StrongString {
	_strong := StrongString{}
	_strong._setValue(s)
	return _strong
}

// Qss will generate a new QString
// with the specified non-encoded string value.
func Qss(s string) tfc.QString {
	str := Ss(s)
	return &str
}

// Sb will generate a new StrongString
// with the specified non-encoded bytes value.
func Sb(b []byte) StrongString {
	return Ss(string(b))
}

// QSb will generate a new QString
// with the specified non-encoded bytes value.
func Qsb(b []byte) tfc.QString {
	str := Ss(string(b))
	return &str
}

// SS will generate a new StrongString
// with the specified non-encoded string value.
func SsPtr(s string) *StrongString {
	strong := StrongString{}
	strong._setValue(s)
	return &strong
}

func ToStrSlice(qs []tfc.QString) []string {
	tmp := make([]string, len(qs))
	for i, current := range qs {
		tmp[i] = current.GetValue()
	}
	return tmp
}

func ToQSlice(strs []string) []tfc.QString {
	tmp := make([]tfc.QString, len(strs))
	for i, current := range strs {
		tmp[i] = SsPtr(current)
	}
	return tmp
}
