// Copyright (c) 2022, Kinson Chow. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps provides an abstract Map interface.
//
// In computer science, an associative array, map, symbol table, or dictionary is an abstract data type composed of a collection of (key, value) pairs, such that each possible key appears just once in the collection.
//
// Operations associated with this data type allow:
// - the addition of a pair to the collection
// - the removal of a pair from the collection
// - the modification of an existing pair
// - the lookup of a value associated with a particular key
//
// Reference: https://en.wikipedia.org/wiki/Associative_array
package maps

import (
	"github.com/kcswag/kcgods/containers"
)

// Map interface that all maps implement
type Map[K, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K

	containers.Container[V]
}

// BidiMap interface that all bidirectional maps implement (extends the Map interface)
type BidiMap[K, V any] interface {
	GetKey(value K) (key V, found bool)
	Map[K, V]
}

type Entry[K, V any] interface {
	GetKey() K
	GetValue() V
	SetValue(value V) V
}
