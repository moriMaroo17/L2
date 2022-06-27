package main

// Empty struct for map 'data'
type void struct{}

// Set implementation
type set struct {
	data map[rune]void
}

// put is a method for put value into set
func (s *set) put(value rune) {
	s.data[value] = void{}
}

// checkContain is a method for check containing value into the set
func (s *set) checkContain(value rune) bool {
	_, ok := s.data[value]
	return ok
}
