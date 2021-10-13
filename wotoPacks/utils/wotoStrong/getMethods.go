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

package wotoStrong

import (
	"reflect"
	"strings"

	tf "wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/utils/wotoStrings"
	wv "wp-server/wotoPacks/utils/wotoValues"
)

// GetValue will give you the real value of this StrongString.
func (s *StrongString) GetValue() string {
	return string(s._value)
}

// length method, will give you the length-as-int of this StrongString.
func (s *StrongString) Length() int {
	return len(s._value)
}

// isEmpty will check if this StrongString is empty or not.
func (s *StrongString) IsEmpty() bool {
	return s._value == nil || len(s._value) == wv.BaseIndex
}

// isEqual will check if the passed-by-value in the arg is equal to this
// StrongString or not.
func (s *StrongString) IsEqual(_q tf.QString) bool {
	if reflect.TypeOf(_q) != reflect.TypeOf(s) {
		return _q.GetValue() == s.GetValue()
	}

	_strong, _ok := _q.(*StrongString)
	if !_ok {
		return false
	}
	// check if the length of them are equal or not.
	if len(s._value) != len(_strong._value) {
		//fmt.Println(len(_s._value), len(_strong._value))
		return false
	}
	for i := 0; i < len(s._value); i++ {
		if s._value[i] != _strong._value[i] {
			//fmt.Println(_s._value[i], _strong._value[i])
			return false
		}
	}
	return true
}

// GetIndexV method will give you the rune in _index.
func (s *StrongString) GetIndexV(_index int) rune {
	if s.IsEmpty() {
		return wv.BaseIndex
	}

	l := len(s._value)

	if _index >= l || l < wv.BaseIndex {

		return s._value[wv.BaseIndex]
	}

	return s._value[_index]
}

// HasSuffix will check if at least there is one suffix is
// presents in this StrongString not.
// the StrongString should ends with at least one of these suffixes.
func (s *StrongString) HasSuffix(values ...string) bool {
	for _, str := range values {
		if strings.HasSuffix(s.GetValue(), str) {
			return true
		}
	}

	return false
}

// HasSuffixes will check if all of the suffixes are
// present in this StrongString or not.
// the StrongString should ends with all of these suffixes.
// usage of this method is not recommended, since you can use
// HasSuffix method with only one string (the longest string).
// this way you will just use too much cpu resources.
func (s *StrongString) HasSuffixes(values ...string) bool {
	for _, str := range values {
		if !strings.HasSuffix(s.GetValue(), str) {
			return false
		}
	}

	return true
}

// HasPrefix will check if at least there is one prefix is
// presents in this StrongString or not.
// the StrongString should starts with at least one of these prefixes.
func (s *StrongString) HasPrefix(values ...string) bool {
	for _, str := range values {
		if strings.HasPrefix(s.GetValue(), str) {
			return true
		}
	}

	return false
}

// HasPrefixes will check if all of the prefixes are
// present in this StrongString or not.
// the StrongString should starts with all of these suffixes.
// usage of this method is not recommended, since you can use
// HasSuffix method with only one string (the longest string).
// this way you will just use too much cpu resources.
func (s *StrongString) HasPrefixes(values ...string) bool {
	for _, str := range values {
		if !strings.HasPrefix(s.GetValue(), str) {
			return false
		}
	}

	return true
}

func (s *StrongString) Split(qs ...tf.QString) []tf.QString {
	strs := wotoStrings.SplitSlice(s.GetValue(), ToStrSlice(qs))
	return ToQSlice(strs)
}

func (s *StrongString) SplitN(n int, qs ...tf.QString) []tf.QString {
	strs := wotoStrings.SplitSliceN(s.GetValue(), ToStrSlice(qs), n)
	return ToQSlice(strs)
}

func (s *StrongString) SplitFirst(qs ...tf.QString) []tf.QString {
	strs := wotoStrings.SplitSliceN(s.GetValue(), ToStrSlice(qs), wv.BaseTwoIndex)
	return ToQSlice(strs)
}

func (s *StrongString) SplitStr(qs ...string) []tf.QString {
	strs := wotoStrings.SplitSlice(s.GetValue(), qs)
	return ToQSlice(strs)
}

func (s *StrongString) SplitStrN(n int, qs ...string) []tf.QString {
	strs := wotoStrings.SplitSliceN(s.GetValue(), qs, n)
	return ToQSlice(strs)
}

func (s *StrongString) SplitStrFirst(qs ...string) []tf.QString {
	strs := wotoStrings.SplitSliceN(s.GetValue(), qs, wv.BaseTwoIndex)
	return ToQSlice(strs)
}

func (s *StrongString) ToQString() tf.QString {
	return s
}

func (s *StrongString) Contains(qs ...tf.QString) bool {
	v := s.GetValue()
	for _, current := range qs {
		if strings.Contains(v, current.GetValue()) {
			return true
		}
	}

	return false
}

func (s *StrongString) ContainsStr(str ...string) bool {
	v := s.GetValue()
	for _, current := range str {
		if strings.Contains(v, current) {
			return true
		}
	}

	return false
}

func (s *StrongString) ContainsAll(qs ...tf.QString) bool {
	v := s.GetValue()
	for _, current := range qs {
		if !strings.Contains(v, current.GetValue()) {
			return false
		}
	}

	return true
}

func (s *StrongString) ContainsStrAll(str ...string) bool {
	v := s.GetValue()
	for _, current := range str {
		if !strings.Contains(v, current) {
			return false
		}
	}

	return true
}

func (s *StrongString) TrimPrefix(qs ...tf.QString) tf.QString {
	final := s.GetValue()
	for _, current := range qs {
		final = strings.TrimPrefix(final, current.GetValue())
	}

	return SsPtr(final)
}

func (s *StrongString) TrimPrefixStr(qs ...string) tf.QString {
	final := s.GetValue()
	for _, current := range qs {
		final = strings.TrimPrefix(final, current)
	}

	return SsPtr(final)
}

func (s *StrongString) TrimSuffix(qs ...tf.QString) tf.QString {
	final := s.GetValue()
	for _, current := range qs {
		final = strings.TrimSuffix(final, current.GetValue())
	}

	return SsPtr(final)
}

func (s *StrongString) TrimSuffixStr(qs ...string) tf.QString {
	final := s.GetValue()
	for _, current := range qs {
		final = strings.TrimSuffix(final, current)
	}

	return SsPtr(final)
}

func (s *StrongString) Trim(qs ...tf.QString) tf.QString {
	final := s.GetValue()
	for _, current := range qs {
		final = strings.Trim(final, current.GetValue())
	}

	return SsPtr(final)
}

func (s *StrongString) TrimStr(qs ...string) tf.QString {
	final := s.GetValue()
	for _, current := range qs {
		final = strings.Trim(final, current)
	}

	return SsPtr(final)
}

func (s *StrongString) Replace(qs, newS tf.QString) tf.QString {
	return s.ReplaceStr(qs.GetValue(), newS.GetValue())
}

func (s *StrongString) ReplaceStr(qs, newS string) tf.QString {
	final := s.GetValue()
	final = strings.ReplaceAll(final, qs, newS)
	return SsPtr(final)
}

// LockSpecial will lock all the defiend special characters.
// This way, you don't actually have to be worry about
// some normal mistakes in spliting strings, cut them out,
// check them. join them, etc...
// WARNING: this method is so dangerous, it's really
// dangerous. we can't say that it's unsafe actually,
// but still it's really dangerous, so if you don't know what the
// fuck are you doing, then please don't use this method.
// this method will not return you a new value, it will effect the
// current value. please consider using it carefully.
func (s *StrongString) LockSpecial() {
	final := s.GetValue()
	// replacing escaped string characters
	// (I mean escaped double quetion mark) is necessary before
	// repairing value.
	final = strings.ReplaceAll(final, wv.BACK_STR, wv.JA_STR)

	// let it repair the string.
	// this function is for repairing these special characters
	// and strings:
	// '=', ':' and "=="
	// it will escape them.
	// if it wasn't for this function, members had to
	// escape all of these bullshits themselves...
	// hahaha, you see, it's actually usefull.
	final = *repairString(&final)

	final = strings.ReplaceAll(final, wv.BACK_FLAG, wv.JA_FLAG)
	final = strings.ReplaceAll(final, wv.BACK_EQUALITY, wv.JA_EQUALITY)
	final = strings.ReplaceAll(final, wv.BACK_DDOT, wv.JA_DDOT)

	s._value = make([]rune, wv.BaseIndex)
	for _, c := range final {
		if c != wv.BaseIndex {
			s._value = append(s._value, c)
		}
	}
}

// UnlockSpecial will unlock all the defiend special characters.
// it will return them to their normal form.
func (s *StrongString) UnlockSpecial() {
	final := s.GetValue()
	final = strings.ReplaceAll(final, wv.JA_FLAG, wv.FLAG_PREFIX)
	final = strings.ReplaceAll(final, wv.JA_STR, wv.STR_SIGN)
	final = strings.ReplaceAll(final, wv.JA_EQUALITY, wv.EqualStr)
	final = strings.ReplaceAll(final, wv.JA_DDOT, wv.DdotSign)

	s._value = make([]rune, wv.BaseIndex)
	for _, c := range final {
		if c != wv.BaseIndex {
			s._value = append(s._value, c)
		}
	}
}
