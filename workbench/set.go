package workbench

type Set map[int]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(item int) {
	s[item] = struct{}{}
}

func (s Set) Remove(item int) {
	delete(s, item)
}

func (s Set) Contains(item int) bool {
	_, exists := s[item]
	return exists
}

func (s Set) Size() int {
	return len(s)
}

// TODO: look up generic set
// type Set map[E comperable] map[E]struct{}
