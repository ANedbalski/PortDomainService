package task

import "io"

type NewPortsListFile struct {
	src io.Reader
}

func NewNewPortsListFile(src io.Reader) *NewPortsListFile {
	return &NewPortsListFile{src: src}
}
