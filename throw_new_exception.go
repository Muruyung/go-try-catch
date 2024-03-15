package gotry

func throwNewException(e any) {
	if e != nil {
		panic(e)
	}
}
