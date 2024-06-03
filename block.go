package gotry

// Block try catch using block format
type Block struct {
	Try   func(ThrowNewException func(any))
	Catch func(e error, st StackTrace)
}

// Do running block try catch
func (b Block) Do() Exception {
	return Try(b.Try).Catch(b.Catch)
}
