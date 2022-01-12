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

package wotoActions

import "strings"

func ParseBatchExecute(data string) ([]BatchExecution, error) {
	if !strings.HasPrefix(data, BatchStr) {
		return nil, ErrBatchParseFailed
	}

	data = strings.TrimPrefix(data, BatchStr)
	if !strings.Contains(data, SepBatchStr) {
		return []BatchExecution{BatchExecution(data)}, nil
	}

	strs := strings.Split(data, SepBatchStr)

	b := make([]BatchExecution, 0)
	for _, s := range strs {
		b = append(b, BatchExecution(s))
	}

	return b, nil
}

func IsActionValid(action RequestAction) bool {
	return _actionsMap[action]
}
