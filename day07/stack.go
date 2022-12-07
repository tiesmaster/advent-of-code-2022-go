package day07

type stack struct {
	slice []dirEntry
}

func newStack() stack {
	return stack{make([]dirEntry, 0)}
}

func (st *stack) push(item ...dirEntry) {
	st.slice = append(st.slice, item...)
}

func (st *stack) pop() dirEntry {
	n := len(st.slice) - 1
	ret := st.slice[n]

	st.slice = st.slice[:n]

	return ret
}

func (st *stack) isEmpty() bool {
	return len(st.slice) == 0
}
