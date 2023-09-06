package set

type Set struct {
	m map[string]struct{}
}

var exists = struct{}{}

/*
	Creates a new set Instance
*/
func NewSet() *Set {
	s := &Set{}
	s.m = make(map[string]struct{})
	return s
}

/*
	Adds a value to the Set
*/
func (s *Set) Add(value string) {
	s.m[value] = exists
}

/*
	Removes a value from the Set
*/
func (s *Set) Remove(value string) {
	delete(s.m, value)
}

/*
	Indicates if the value exists in the state
*/
func (s *Set) Contains(value string) bool {
	_, present := s.m[value]
	return present
}

func (s *Set) GetValues() []string {

	values := make([]string, 0, len(s.m))

	for k := range s.m {
		values = append(values, k)
	}
	return values
}

func (s *Set) Size() int {
	return len(s.m)
}
