package filex

type FileOrDir struct {
	IsDir			bool
	Childs 			[]*FileOrDir
	Data			[]byte
}

func (FileOrDir)Write() error {
	return nil
}