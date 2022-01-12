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

package versioning

const (
	// MaxVersion is the maximum version of the game. if the client's version is
	// more than this version, it's not acceptable.
	currentVersion = "2.1.1.5014"

	versionHash = "f302bd7ffacbd295194f86620002b8250e8e9be0233a8055bcebc82c8612843ff9e0f09e42015d5e75581cc93d4c29a91388ed411641b543c8fb7b5a26a2a8b8"
)

const (
	//-----------------------------------------------------
	userAgentValue = "wp-client" // not optional

	//-----------------------------------------------------

	// same domain header field means we are sending the
	// HTTP request to the same domain of our referer domain
	// (and origin).
	// we expect it to send us the respond using the same protocol.
	// The same-origin policy is a critical security mechanism that
	// restricts how a document or script loaded by one origin can interact with
	// a resource from another origin.
	// we should set its value to 1.
	//
	//  > see also: https://developer.mozilla.org/ja/docs/Web/Security/Same-origin_policy
	//xSameDomainKey   = "X-Same-Domain" // not optional
	//xSameDomainValue = "1"             // not optional

	//-----------------------------------------------------

	//ltwVersionKey     = "madokaplay_version"
	//ltwVersionHashKey = "madokaplay_version_hash"
)

// batch execution values
const (
	Batch_CHECK_VERSION = "check_version"
)
