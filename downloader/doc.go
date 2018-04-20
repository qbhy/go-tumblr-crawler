package downloader

import "io"

type Document struct {
	io.Reader
	str string
}

func NewDocument(str string) *Document {
	return &Document{
		str: str,
	}
}

func (this *Document) Read() string {
	return this.str
}
