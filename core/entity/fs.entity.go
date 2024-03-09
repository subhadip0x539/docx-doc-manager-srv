package entity

type FSFileObject struct {
	DocumentId   string
	DocumentName string
	Size         int64
	Type         string
	Content      []byte
}
