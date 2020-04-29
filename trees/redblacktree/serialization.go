// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import (
	"encoding/json"

	"github.com/afiodorov/countedredblacktree/containers"
	"github.com/afiodorov/countedredblacktree/utils"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Tree)(nil)
	var _ containers.JSONDeserializer = (*Tree)(nil)
}

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree) ToJSON() ([]byte, error) {
	elements := make(map[string]int)
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Count()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree) FromJSON(data []byte) error {
	elements := make(map[float64]int)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			for i := 0; i < value; i++ {
				tree.Put(key)
			}
		}
	}
	return err
}
