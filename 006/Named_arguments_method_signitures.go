package main

// In the section we look at how we can name the returnes of our method signitures
// as well as naming the return types

type File_handle interface {
	Copy(sourceFile string, destinationFile string) (byteCode int)
}

type file struct {
	sourcefile any
	destfile   any
}

func (f file) Copy() bool {
	return true
}

func main() {

}
