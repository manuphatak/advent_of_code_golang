package shared

import (
	"fmt"
	"strings"
)

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E] { return make(Set[E]) }
func (s Set[E]) Add(v E) {
	s[v] = struct{}{}
}
func (s Set[E]) Remove(v E) {
	delete(s, v)
}

func (s Set[E]) Contains(v E) bool {
	_, ok := s[v]
	return ok
}

func (s Set[E]) String() string {
	keys := []string{}
	for k := range s {
		keys = append(keys, fmt.Sprintf("%v", k))
	}

	return fmt.Sprintf("Set[%v]", strings.Join(keys, " "))
}

func (s Set[E]) Size() int {
	return len(s)
}
