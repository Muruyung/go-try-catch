package try

func ThrowNewException(e any) {
	if e != nil {
		panic(e)
	}
}
