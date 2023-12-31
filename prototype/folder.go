package main

import "fmt"

type Folder struct {
	children []Inode
	name     string
}

var _ Inode = (*Folder)(nil)

func (f *Folder) print(indentation string) {
	fmt.Printf("%s%s\n", indentation, f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
