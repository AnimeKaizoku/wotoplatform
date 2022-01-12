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

package interfaces

type QString interface {
	Length() int
	IsEmpty() bool
	GetValue() string
	GetIndexV(int) rune
	IsEqual(QString) bool
	Split(...QString) []QString
	SplitN(int, ...QString) []QString
	SplitFirst(qs ...QString) []QString
	SplitStr(...string) []QString
	SplitStrN(int, ...string) []QString
	SplitStrFirst(...string) []QString
	Contains(...QString) bool
	ContainsStr(...string) bool
	ContainsAll(...QString) bool
	ContainsStrAll(...string) bool
	TrimPrefix(...QString) QString
	TrimPrefixStr(...string) QString
	TrimSuffix(...QString) QString
	TrimSuffixStr(...string) QString
	Trim(qs ...QString) QString
	TrimStr(...string) QString
	Replace(qs, newS QString) QString
	ReplaceStr(qs, newS string) QString
	LockSpecial()
	UnlockSpecial()
}
