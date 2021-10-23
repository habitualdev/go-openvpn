/*
 * Copyright (C) 2018 The "MysteriumNetwork/go-openvpn" Authors.
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

package config

import (
	"fmt"
)

// ToArguments serializes openvpn configuration structure to a list of command line arguments which can be passed to openvpn process
// it also serialize file style options to appropriate files inside given config directory
func (config GenericConfig) ToArguments() ([]string, error) {
	arguments := make([]string, 0)

	for _, item := range config.options {
		option, ok := item.(optionCliSerializable)
		if !ok {
			return nil, fmt.Errorf("unserializable option '%s': %#v", item.getName(), item)
		}

		optionValues, err := option.toCli()
		if err != nil {
			return nil, err
		}

		arguments = append(arguments, optionValues...)
	}

	return arguments, nil
}

type optionCliSerializable interface {
	toCli() ([]string, error)
}
