// Copyright (c) 2022, Kinson Chow. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treeset

import (
	"encoding/json"
)

// Assert Serialization implementation
//var _ containers.JSONSerializer = (*Set)(nil)
//var _ containers.JSONDeserializer = (*Set)(nil)

// ToJSON outputs the JSON representation of the set.
func (set *Set[E]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Set[E]) FromJSON(data []byte) error {
	var elements []E
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (set *Set[E]) UnmarshalJSON(bytes []byte) error {
	return set.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (set *Set[E]) MarshalJSON() ([]byte, error) {
	return set.ToJSON()
}
