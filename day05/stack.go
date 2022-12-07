package day05

type stack struct {
	slice []rune
}

func newStack() stack {
	return stack{make([]rune, 0)}
}

func (st *stack) push(item rune) {
	st.slice = append(st.slice, item)
}

func (st *stack) pop() rune {
	n := len(st.slice) - 1
	ret := st.slice[n]

	st.slice = st.slice[:n]

	return ret
}

func (st *stack) isEmpty() bool {
	return len(st.slice) == 0
}
