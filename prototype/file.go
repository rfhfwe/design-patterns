package main

import "fmt"

type File struct {
	name string
}

var _ Inode = (*File)(nil)

func (f *File) print(indentation string) {
	fmt.Printf("%s%s\n", indentation, f.name)
}

func (f *File) clone() Inode {
	return &File{
		name: f.name + "_clone",
	}
}
