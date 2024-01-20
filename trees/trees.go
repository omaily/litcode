package trees

type element struct {
	data interface{}
	next *element
}

func StartTask() {
	a := &element{}
	println(a)
}
