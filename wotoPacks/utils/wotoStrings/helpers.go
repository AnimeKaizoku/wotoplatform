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

package wotoStrings

import (
	"regexp"
	"strings"
	"unicode"

	wv "wp-server/wotoPacks/utils/wotoValues"
)

func Split(s string, separator ...string) []string {
	return SplitSliceN(s, separator, -1)
}

// SplitWhite splits the string with the given separator
// and will remove the white spaces slices from the results
func SplitWhite(s string, separator ...string) []string {
	return SplitSliceNWhite(s, separator, -1)
}

func SplitN(s string, n int, separator ...string) []string {
	return SplitSliceN(s, separator, n)
}

func SplitSlice(s string, separator []string) []string {
	return SplitSliceN(s, separator, -1)
}

func SplitSliceN(s string, separator []string, n int) []string {
	if len(separator) == wv.BaseIndex {
		return []string{s}
	}

	var m string
	for i, f := range separator {
		if i != len(separator)-1 {
			m += regexp.QuoteMeta(f) + wv.OrRegexp
		} else {
			m += regexp.QuoteMeta(f)
		}
	}

	re, err := regexp.Compile(m)
	if err != nil {
		return []string{s}
	}

	return FixSplit(re.Split(s, n))
}

func SplitSliceNWhite(s string, separator []string, n int) []string {
	if len(separator) == wv.BaseIndex {
		return []string{s}
	}

	var m string
	for i, f := range separator {
		if i != len(separator)-1 {
			m += regexp.QuoteMeta(f) + wv.OrRegexp
		} else {
			m += regexp.QuoteMeta(f)
		}
	}

	re, err := regexp.Compile(m)
	if err != nil {
		return []string{s}
	}

	return FixSplitWhite(re.Split(s, n))
}

// FixSplit will fix the bullshit bug in the
// Split function (which is not ignoring the spaces between strings).
func FixSplit(myStrings []string) []string {
	final := make([]string, wv.BaseIndex, cap(myStrings))

	for _, current := range myStrings {
		if !IsEmpty(&current) {
			final = append(final, current)
		}
	}

	return final
}

// FixSplit will fix the bullshit bug in the
// Split function (which is not ignoring the spaces between strings).
func FixSplitWhite(myStrings []string) []string {
	final := make([]string, wv.BaseIndex, cap(myStrings))

	for _, current := range myStrings {
		if strings.TrimSpace(current) != "" {
			final = append(final, current)
		}
	}

	return final
}

// IsEmpty function will check if the passed-by
// string value is empty or not.
func IsEmpty(s *string) bool {
	return s == nil || len(*s) == wv.BaseIndex
}

// AreEqual will check if two string ptr are equal to each other or not.
func AreEqual(s1, s2 *string) bool {
	if s1 == nil && s2 != nil {
		return len(*s2) == 0
	} else if s1 != nil && s2 == nil {
		return len(*s1) == 0
	}

	return s1 == s2 || *s1 == *s2
}

// YesOrNo returns yes if v is true, otherwise no.
func YesOrNo(v bool) string {
	if v {
		return wv.Yes
	} else {
		return wv.No
	}
}

func ToArray(strs ...string) []string {
	return strs
}

func IsAllNumber(str string) bool {
	for _, s := range str {
		if !IsRuneNumber(s) {
			return false
		}
	}

	return true
}

func IsAllNumbers(str ...string) bool {
	for _, ss := range str {
		if !IsAllNumber(ss) {
			return false
		}
	}

	return true
}

func IsRuneNumber(r rune) bool {
	if r <= unicode.MaxLatin1 {
		return '0' <= r && r <= '9'
	}

	return false
}
