package main

import (
	"bytes"
	"fmt"
)

// IntSet: 負ではない小さな整数のセット
// そのゼロ値は空セットを表す
type IntSet struct {
	words []uint64
}

// Has: 負ではない値 x をセットが含んでいるか否かを報告
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Len: 要素数を返す
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

// Add: セットに負ではない値 x を追加
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(n ...int) {
	for _, x := range n {
		word, bit := x/64, uint(x%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

// Remove: セットから x を取り除く
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit
}

// Clear: セットから全ての要素を取り除く
func (s *IntSet) Clear() {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		s.words[i] = 0
	}
}

// Copy: セットのコピーを返す
func (s *IntSet) Copy() *IntSet {
	var copied IntSet
	for i, sword := range s.words {
		copied.words = append(copied.words, 0)
		if sword == 0 {
			continue
		}
		copied.words[i] = sword
	}
	return &copied

	//return &IntSet{append([]uint64(nil), s.words...)}
}

// UnionWith は s, t の和集合を s に設定
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

// s, t の対称差
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String: "{1 2 3}" の形式の文字列としてセットを返す
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if ((word >> uint(j)) & 1) == 1 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String(), x.Len())
	fmt.Println(x.Copy())
	x.Remove(144)
	fmt.Println(x.String(), x.Len())
	fmt.Println(x.Copy())
	x.Clear()
	fmt.Println(x.String(), x.Len())
	x.AddAll(1, 4, 5345)
	fmt.Println(x.String(), x.Len())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String(), y.Len())

	x.UnionWith(&y)
	fmt.Println(x.String(), x.Len())

	fmt.Println(x.Has(9), x.Has(123))
}
