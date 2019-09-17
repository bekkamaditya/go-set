package myset

// type mystring string
//
// func (s1 mystring) Compare(s2 interface{}) bool {
// 	compared := strings.Compare(fmt.Sprintf("%s", s1), fmt.Sprintf("%s", s2))
// 	return compared == 0
// }
//
// func main() {
// 	fmt.Println("Implementing set data structure")
//
// 	// var arr []Element
// 	// arr[0] = mystring("aditya")
// 	arr := []Element{
// 		mystring("aditya"),
// 		mystring("bekkam"),
// 		mystring("venkata"),
// 		mystring("venkata"),
// 		mystring("venkata"),
// 		mystring("venkata"),
// 	}
// 	s := NewSet(arr)
// 	fmt.Printf("%+v\n", s)
// 	arr1 := []Element{
// 		mystring("foo"),
// 		mystring("bar"),
// 		mystring("bekkam"),
// 		mystring("one"),
// 	}
// 	s1 := NewSet(arr1)
// 	fmt.Printf("%+v\n", s1)
// 	fmt.Printf("%+v\n", s.Union(s1))
// 	fmt.Printf("%+v\n", s.Intersection(s1))
// 	fmt.Printf("%+v\n", s.Difference(s1))
// }

// Element ...
type Element interface {
	Compare(interface{}) bool
}

// Set ...
type Set struct {
	elements []Element
}

// NewSet ...
func NewSet(elem []Element) *Set {
	s := &Set{
		elements: elem,
	}
	s.RemoveDuplicates()
	return s
}

func (s *Set) copySet() *Set {
	resultantSet := &Set{
		elements: make([]Element, 0),
	}
	for _, elem := range s.elements {
		resultantSet.elements = append(resultantSet.elements, elem)
	}
	return resultantSet
}

func (s *Set) isElementPresent(v Element) bool {
	for _, item := range s.elements {
		if item.Compare(v) {
			return true
		}
	}
	return false
}

func (s *Set) indexOf(v Element) int {
	for index, item := range s.elements {
		if item.Compare(v) {
			return index
		}
	}
	return -1
}

// Union ...
func (s *Set) Union(set2 *Set) *Set {
	resultantSet := s.copySet()
	for _, elem := range set2.elements {
		if !resultantSet.isElementPresent(elem) {
			resultantSet.elements = append(resultantSet.elements, elem)
		}
	}
	return resultantSet
}

// Intersection ...
func (s *Set) Intersection(set2 *Set) *Set {
	resultantSet := &Set{
		elements: make([]Element, 0),
	}
	for _, elem := range s.elements {
		if set2.isElementPresent(elem) {
			resultantSet.elements = append(resultantSet.elements, elem)
		}
	}
	return resultantSet
}

// Difference ...
func (s *Set) Difference(set2 *Set) *Set {
	resultantSet := s.copySet()
	for i, elem := range resultantSet.elements {
		if set2.isElementPresent(elem) {
			resultantSet.elements[i] = resultantSet.elements[len(resultantSet.elements)-1]
			resultantSet.elements[len(resultantSet.elements)-1] = nil
			resultantSet.elements = resultantSet.elements[:len(resultantSet.elements)-1]
		}
	}
	return resultantSet
}

// RemoveDuplicates ...
func (s *Set) RemoveDuplicates() {
	tempArr := s.elements
	s.elements = make([]Element, 0)
	for _, elem := range tempArr {
		if !s.isElementPresent(elem) {
			s.elements = append(s.elements, elem)
		}
	}
}
