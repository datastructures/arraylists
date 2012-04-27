package arraylists

type Element struct {
	Value interface{}
	List  *ArrayStack
}

type ArrayStack struct {
	array []*Element
}

func NewArrayStack() *ArrayStack {
	a := &ArrayStack{}
	a.array = make([]*Element, 0, 10)
	return a
}

func (a *ArrayStack) Push(value interface{}) (e *Element) {
	e = &Element{Value:value, List:a}
	a.array = append(a.array, e)
	return e
}

func (a *ArrayStack) Pop() (e *Element) {
	if a.Len() > 0 {
		e = a.array[len(a.array)-1]
		a.array = a.array[:len(a.array)-1]
		return e
	}
	return e
}

func (a *ArrayStack) Remove() (e *Element) {
	if a.Len() >0 {
		e = a.array[0]
		a.array= a.array[1:len(a.array)]
		return
	}
	return
}

func (a ArrayStack) Len() int { return len(a.array) }
