// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestRedBlackTreePut(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5)
	tree.Put(6)
	tree.Put(7)
	tree.Put(3)
	tree.Put(4)
	tree.Put(1)
	tree.Put(2)
	tree.Put(1)
	tree.Put(1)
	tree.Put(2)

	if actualValue := tree.Size(); actualValue != 10 {
		t.Errorf("Got %v expected %v", actualValue, 10)
	}
	if actualValue, expectedValue := fmt.Sprintf("%v", tree.Keys()), "[1 1 1 2 2 3 4 5 6 7]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	tests1 := [][]interface{}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
		{7, true},
		{8, false},
	}

	for _, test := range tests1 {
		// retrievals
		actualFound := tree.Get(test[0])
		if actualFound != test[1] {
			t.Errorf("Got %v expected %v", actualFound, test[1])
		}
	}
}

func TestRedBlackTreeRemove(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5)
	tree.Put(6)
	tree.Put(7)
	tree.Put(3)
	tree.Put(4)
	tree.Put(1)
	tree.Put(2)
	tree.Put(1)

	tree.Remove(5)
	tree.Remove(6)
	tree.Remove(7)
	tree.Remove(8)
	tree.Remove(5)

	if actualValue, expectedValue := fmt.Sprintf("%v", tree.Keys()), "[1 1 2 3 4]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if actualValue := tree.Size(); actualValue != 5 {
		t.Errorf("Got %v expected %v", actualValue, 5)
	}

	tests2 := [][]interface{}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, false},
		{6, false},
		{7, false},
		{8, false},
	}

	for _, test := range tests2 {
		actualFound := tree.Get(test[0])
		if actualFound != test[1] {
			t.Errorf("Got %v expected %v", actualFound, test[1])
		}
	}

	tree.Remove(1)
	tree.Remove(4)
	tree.Remove(2)
	tree.Remove(3)
	tree.Remove(2)
	tree.Remove(2)

	if actualValue, expectedValue := fmt.Sprintf("%v", tree.Keys()), "[1]"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	tree.Remove(1)
	if empty, size := tree.Empty(), tree.Size(); empty != true || size != -0 {
		t.Errorf("Got %v expected %v", empty, true)
	}

}

func TestRedBlackTreeLeftAndRight(t *testing.T) {
	tree := NewWithIntComparator()

	if actualValue := tree.Left(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	if actualValue := tree.Right(); actualValue != nil {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}

	tree.Put(1)
	tree.Put(5)
	tree.Put(6)
	tree.Put(7)
	tree.Put(3)
	tree.Put(4)
	tree.Put(1) // overwrite
	tree.Put(2)

	if actualValue, expectedValue := fmt.Sprintf("%d", tree.Left().Key), "1"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}

	if actualValue, expectedValue := fmt.Sprintf("%d", tree.Right().Key), "7"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeCeilingAndFloor(t *testing.T) {
	tree := NewWithIntComparator()

	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
	if node, found := tree.Ceiling(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	tree.Put(5)
	tree.Put(6)
	tree.Put(7)
	tree.Put(3)
	tree.Put(4)
	tree.Put(1)
	tree.Put(2)

	if node, found := tree.Floor(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Floor(0); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}

	if node, found := tree.Ceiling(4); node.Key != 4 || !found {
		t.Errorf("Got %v expected %v", node.Key, 4)
	}
	if node, found := tree.Ceiling(8); node != nil || found {
		t.Errorf("Got %v expected %v", node, "<nil>")
	}
}

func TestRedBlackTreeIteratorNextOnEmpty(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()
	for it.Next() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestRedBlackTreeIteratorPrevOnEmpty(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()
	for it.Prev() {
		t.Errorf("Shouldn't iterate on empty tree")
	}
}

func TestRedBlackTreeIterator1Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5)
	tree.Put(6)
	tree.Put(7)
	tree.Put(3)
	tree.Put(4)
	tree.Put(1)
	tree.Put(2)
	tree.Put(1)
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1 (2)
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count += it.Count()
		key := it.Key()
		if actualValue, expectedValue := key, count-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator1Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(5)
	tree.Put(6)
	tree.Put(7)
	tree.Put(3)
	tree.Put(4)
	tree.Put(1)
	tree.Put(2)
	tree.Put(1)
	// │   ┌── 7
	// └── 6
	//     │   ┌── 5
	//     └── 4
	//         │   ┌── 3
	//         └── 2
	//             └── 1
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.Size()
	for it.Prev() {
		key := it.Key()
		if actualValue, expectedValue := key, countDown-1; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		countDown -= it.Count()
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator2Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3)
	tree.Put(1)
	tree.Put(2)
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator2Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3)
	tree.Put(1)
	tree.Put(2)
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.Size()
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator3Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(1)
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		switch key {
		case count:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, count; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator3Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(1)
	it := tree.Iterator()
	for it.Next() {
	}
	countDown := tree.Size()
	for it.Prev() {
		key := it.Key()
		switch key {
		case countDown:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		default:
			if actualValue, expectedValue := key, countDown; actualValue != expectedValue {
				t.Errorf("Got %v expected %v", actualValue, expectedValue)
			}
		}
		countDown--
	}
	if actualValue, expectedValue := countDown, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator4Next(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13)
	tree.Put(8)
	tree.Put(17)
	tree.Put(1)
	tree.Put(11)
	tree.Put(15)
	tree.Put(25)
	tree.Put(6)
	tree.Put(22)
	tree.Put(27)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
	}
	if actualValue, expectedValue := count, tree.Size(); actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIterator4Prev(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(13)
	tree.Put(8)
	tree.Put(17)
	tree.Put(1)
	tree.Put(11)
	tree.Put(15)
	tree.Put(25)
	tree.Put(6)
	tree.Put(22)
	tree.Put(27)
	// │           ┌── 27
	// │       ┌── 25
	// │       │   └── 22
	// │   ┌── 17
	// │   │   └── 15
	// └── 13
	//     │   ┌── 11
	//     └── 8
	//         │   ┌── 6
	//         └── 1
	it := tree.Iterator()
	count := tree.Size()
	for it.Next() {
	}
	for it.Prev() {
		count--
	}
	if actualValue, expectedValue := count, 0; actualValue != expectedValue {
		t.Errorf("Size different. Got %v expected %v", actualValue, expectedValue)
	}
}

func TestRedBlackTreeIteratorBegin(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3)
	tree.Put(1)
	tree.Put(2)
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	for it.Next() {
	}

	it.Begin()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Next()
	if key := it.Key(); key != 1 {
		t.Errorf("Got %v expected %v", key, 1)
	}
}

func TestRedBlackTreeIteratorEnd(t *testing.T) {
	tree := NewWithIntComparator()
	it := tree.Iterator()

	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	tree.Put(3)
	tree.Put(1)
	tree.Put(2)
	it.End()
	if it.node != nil {
		t.Errorf("Got %v expected %v", it.node, nil)
	}

	it.Prev()
	if key := it.Key(); key != 3 {
		t.Errorf("Got %v expected %v", key, 3)
	}
}

func TestRedBlackTreeIteratorFirst(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3)
	tree.Put(1)
	tree.Put(2)
	it := tree.Iterator()
	if actualValue, expectedValue := it.First(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key := it.Key(); key != 1 {
		t.Errorf("Got %v expected %v", key, 1)
	}
}

func TestRedBlackTreeIteratorLast(t *testing.T) {
	tree := NewWithIntComparator()
	tree.Put(3)
	tree.Put(1)
	tree.Put(2)
	it := tree.Iterator()
	if actualValue, expectedValue := it.Last(), true; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
	if key := it.Key(); key != 3 {
		t.Errorf("Got %v expected %v", key, 3)
	}
}

func TestRedBlackTreeSerialization(t *testing.T) {
	tree := NewWithStringComparator()
	tree.Put("c")
	tree.Put("b")
	tree.Put("a")

	var err error
	assert := func() {
		if actualValue, expectedValue := tree.Size(), 3; actualValue != expectedValue {
			t.Errorf("Got %v expected %v", actualValue, expectedValue)
		}
		if actualValue := tree.Keys(); actualValue[0].(string) != "a" || actualValue[1].(string) != "b" || actualValue[2].(string) != "c" {
			t.Errorf("Got %v expected %v", actualValue, "[a,b,c]")
		}
		if err != nil {
			t.Errorf("Got error %v", err)
		}
	}

	assert()

	json, err := tree.ToJSON()
	assert()

	err = tree.FromJSON(json)
	assert()
}

func TestRedBlackTreeCounts(t *testing.T) {
	countGreaterOrEqual := func(a int, array []int) (ret int) {
		for _, b := range array {
			if b >= a {
				ret++
			}
		}
		return
	}

	countGreater := func(a int, array []int) (ret int) {
		for _, b := range array {
			if b > a {
				ret++
			}
		}
		return
	}

	countSmallerOrEqual := func(a int, array []int) (ret int) {
		for _, b := range array {
			if b <= a {
				ret++
			}
		}
		return
	}

	countSmaller := func(a int, array []int) (ret int) {
		for _, b := range array {
			if b < a {
				ret++
			}
		}
		return
	}

	r := rand.New(rand.NewSource(17))
	nTests := 20
	arrSize := 100
	upperBound := 30

	for i := 0; i < nTests; i++ {
		tree := NewWithIntComparator()
		array := make([]int, arrSize)
		for j := 0; j < len(array); j++ {
			randVal := r.Intn(upperBound)
			for randVal == 5 {
				// skip 5
				randVal = r.Intn(upperBound)
			}
			array[j] = randVal
			tree.Put(randVal)
		}

		testVals := make([]int, len(array))
		copy(testVals, array)
		testVals = append(testVals, -10, upperBound+10, 5)

		for _, e := range array {
			if expected, actual := countGreaterOrEqual(e, array), tree.CountGreaterOrEqual(e); expected != actual {
				t.Errorf("GreaterOrEqual: Got %v expected %v", actual, expected)
			}
			if expected, actual := countGreater(e, array), tree.CountGreater(e); expected != actual {
				t.Errorf("Greater: Got %v expected %v", actual, expected)
			}
			if expected, actual := countSmallerOrEqual(e, array), tree.CountSmallerOrEqual(e); expected != actual {
				t.Errorf("SmallerOrEqual: Got %v expected %v", actual, expected)
			}
			if expected, actual := countSmaller(e, array), tree.CountSmaller(e); expected != actual {
				t.Errorf("Smaller: Got %v expected %v", actual, expected)
			}
		}
	}
}

func TestRedBlackTreeNumGreater(t *testing.T) {
	countGreater := func(a int, array []int) (ret int) {
		for _, b := range array {
			if b > a {
				ret++
			}
		}
		return
	}

	r := rand.New(rand.NewSource(17))
	nTests := 20
	arrSize := 100
	upperBound := 30

	for i := 0; i < nTests; i++ {
		tree := NewWithIntComparator()
		array := make([]int, arrSize)
		for j := 0; j < len(array); j++ {
			randVal := r.Intn(upperBound)
			array[j] = randVal
			tree.Put(randVal)
		}

		sort.Ints(array)

		it := tree.Iterator()
		it.Begin()
		for i := 0; it.Next(); {
			for j := 0; j < it.Count(); i, j = i+1, j+1 {
				if expected, actual := array[i], it.Key().(int); expected != actual {
					t.Errorf("Iterator: Got %v expected %v", actual, expected)
				}
				if expected, actual := countGreater(array[i], array), it.NumGreater(); expected != actual {
					t.Errorf("NumGreater: Got %v expected %v", actual, expected)
				}
			}
		}
	}
}

func TestRedBlackTreeNoFloor(t *testing.T) {
	tree := NewWithFloat64Comparator()
	tree.Put(10.0)
	tree.Put(20.0)
	tree.Put(30.0)

	if expected, actual := tree.CountSmallerOrEqual(5.0), 0; expected != actual {
		t.Errorf("SmallerOrEqual: Got %v expected %v", actual, expected)
	}

	if expected, actual := tree.CountSmaller(5.0), 0; expected != actual {
		t.Errorf("Smaller: Got %v expected %v", actual, expected)
	}
}

func TestRedBlackTreeNoCeiling(t *testing.T) {
	tree := NewWithFloat64Comparator()
	tree.Put(1.0)
	tree.Put(2.0)
	tree.Put(3.0)

	if expected, actual := tree.CountGreaterOrEqual(5.0), 0; expected != actual {
		t.Errorf("GreaterOrEqual: Got %v expected %v", actual, expected)
	}

	if expected, actual := tree.CountGreater(5.0), 0; expected != actual {
		t.Errorf("Greater: Got %v expected %v", actual, expected)
	}
}

func benchmarkGet(b *testing.B, tree *Tree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Get(n)
		}
	}
}

func benchmarkPut(b *testing.B, tree *Tree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Put(n)
		}
	}
}

func benchmarkRemove(b *testing.B, tree *Tree, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			tree.Remove(n)
		}
	}
}

func BenchmarkRedBlackTreeGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreeGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkGet(b, tree, size)
}

func BenchmarkRedBlackTreePut100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreePut100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkPut(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}

func BenchmarkRedBlackTreeRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	tree := NewWithIntComparator()
	for n := 0; n < size; n++ {
		tree.Put(n)
	}
	b.StartTimer()
	benchmarkRemove(b, tree, size)
}
