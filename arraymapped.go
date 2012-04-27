package arraylists

type AMElement struct {
	Value interface{}
	List  *ArrayMapped
}

type ArrayMapped struct {
	mapp map[int]*AMElement
	l    int
}

func NewArrayMapped(n int) *ArrayMapped {
	a := &ArrayMapped{mapp: make(map[int]*AMElement, n), l: 0}
	a.l = n
	return a
}

func (a *ArrayMapped) Add(value interface{}) *AMElement {
	e:= &AMElement{Value:value, List:a}

	return e
}
