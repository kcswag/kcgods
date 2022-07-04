// Copyright (c) 2022, Kinson Chow. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arrayqueue

// Assert Serialization implementation
//var _ containers.JSONSerializer = (*Queue)(nil)
//var _ containers.JSONDeserializer = (*Queue)(nil)

// ToJSON outputs the JSON representation of the queue.
func (queue *Queue[T]) ToJSON() ([]byte, error) {
	return queue.list.ToJSON()
}

// FromJSON populates the queue from the input JSON representation.
func (queue *Queue[T]) FromJSON(data []byte) error {
	return queue.list.FromJSON(data)
}

// UnmarshalJSON @implements json.Unmarshaler
func (queue *Queue[T]) UnmarshalJSON(bytes []byte) error {
	return queue.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (queue *Queue[T]) MarshalJSON() ([]byte, error) {
	return queue.ToJSON()
}
