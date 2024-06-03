package gotry

// throwNewException running panic error
func throwNewException(e any) {
	if e != nil {
		panic(e)
	}
}
