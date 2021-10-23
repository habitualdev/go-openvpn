/*
 * Copyright (C) 2019 The "MysteriumNetwork/go-openvpn" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package openvpn

import "github.com/mysteriumnetwork/go-openvpn/openvpn/log"

// UseLogger sets go-openvpn library logger.
func UseLogger(l log.Logger) {
	log.UseLogger(l)
}

// UseDefaultLogger resets logger to the default logger.
func UseDefaultLogger() {
	log.UseDefaultLogger()
}
