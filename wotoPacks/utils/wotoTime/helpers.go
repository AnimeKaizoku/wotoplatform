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

package wotoTime

import (
	"time"
	wv "wp-server/wotoPacks/utils/wotoValues"
)

func formatTimeToList(t time.Time) []int {
	hour, min, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, min, hour, day, int(month), year}
}

// With initialize Now with time
func With(t time.Time) *Now {
	config := DefaultConfig
	if config == nil {
		config = &Config{
			WeekStartDay: WeekStartDay,
			TimeFormats:  TimeFormats,
		}
	}

	return &Now{Time: t, Config: config}
}

// New initialize Now with time
func New(t time.Time) *Now {
	return With(t)
}

// BeginningOfMinute beginning of minute
func BeginningOfMinute() time.Time {
	return With(time.Now()).BeginningOfMinute()
}

// BeginningOfHour beginning of hour
func BeginningOfHour() time.Time {
	return With(time.Now()).BeginningOfHour()
}

// BeginningOfDay beginning of day
func BeginningOfDay() time.Time {
	return With(time.Now()).BeginningOfDay()
}

// BeginningOfWeek beginning of week
func BeginningOfWeek() time.Time {
	return With(time.Now()).BeginningOfWeek()
}

// BeginningOfMonth beginning of month
func BeginningOfMonth() time.Time {
	return With(time.Now()).BeginningOfMonth()
}

// BeginningOfQuarter beginning of quarter
func BeginningOfQuarter() time.Time {
	return With(time.Now()).BeginningOfQuarter()
}

// BeginningOfYear beginning of year
func BeginningOfYear() time.Time {
	return With(time.Now()).BeginningOfYear()
}

// EndOfMinute end of minute
func EndOfMinute() time.Time {
	return With(time.Now()).EndOfMinute()
}

// EndOfHour end of hour
func EndOfHour() time.Time {
	return With(time.Now()).EndOfHour()
}

// EndOfDay end of day
func EndOfDay() time.Time {
	return With(time.Now()).EndOfDay()
}

// EndOfWeek end of week
func EndOfWeek() time.Time {
	return With(time.Now()).EndOfWeek()
}

// EndOfMonth end of month
func EndOfMonth() time.Time {
	return With(time.Now()).EndOfMonth()
}

// EndOfQuarter end of quarter
func EndOfQuarter() time.Time {
	return With(time.Now()).EndOfQuarter()
}

// EndOfYear end of year
func EndOfYear() time.Time {
	return With(time.Now()).EndOfYear()
}

// Monday monday
func Monday() time.Time {
	return With(time.Now()).Monday()
}

// Sunday sunday
func Sunday() time.Time {
	return With(time.Now()).Sunday()
}

// EndOfSunday end of sunday
func EndOfSunday() time.Time {
	return With(time.Now()).EndOfSunday()
}

// Parse parse string to time
func Parse(strs ...string) (time.Time, error) {
	return With(time.Now()).Parse(strs...)
}

// ParseInLocation parse string to time in location
func ParseInLocation(loc *time.Location, strs ...string) (time.Time, error) {
	return With(time.Now().In(loc)).Parse(strs...)
}

// MustParse must parse string to time or will panic
func MustParse(strs ...string) time.Time {
	return With(time.Now()).MustParse(strs...)
}

// MustParseInLocation must parse string to time in location or will panic
func MustParseInLocation(loc *time.Location, strs ...string) time.Time {
	return With(time.Now().In(loc)).MustParse(strs...)
}

// Between check now between the begin, end time or not
func Between(time1, time2 string) bool {
	return With(time.Now()).Between(time1, time2)
}

func GenerateCurrentDateTime() string {
	// dd/MM/yyyy HH:mm:ss

	t := time.Now()

	str := wv.MakeSureNum(t.Day(), 2) + "/"
	str += wv.MakeSureNum(int(t.Month()), 2) + "/"
	str += wv.MakeSureNum(t.Year(), 4) + " "
	str += wv.MakeSureNum(t.Hour(), 2) + ":"
	str += wv.MakeSureNum(t.Minute(), 2) + ":"
	str += wv.MakeSureNum(t.Second(), 2)

	return str
}
