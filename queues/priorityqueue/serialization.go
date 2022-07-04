// Copyright (c) 2022, Kinson Chow. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package priorityqueue

// Assert Serialization implementation
//var _ containers.JSONSerializer = (*Queue)(nil)
//var _ containers.JSONDeserializer = (*Queue)(nil)

// ToJSON outputs the JSON representation of the queue.
func (queue *Queue[E]) ToJSON() ([]byte, error) {
	return queue.heap.ToJSON()
}

// FromJSON populates the queue from the input JSON representation.
func (queue *Queue[E]) FromJSON(data []byte) error {
	return queue.heap.FromJSON(data)
}

// UnmarshalJSON @implements json.Unmarshaler
func (queue *Queue[E]) UnmarshalJSON(bytes []byte) error {
	return queue.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (queue *Queue[E]) MarshalJSON() ([]byte, error) {
	return queue.ToJSON()
}
