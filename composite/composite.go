package main

import "fmt"

type Component interface {
	Search(string)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("search folder with %s, in name:%s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}

type File struct {
	name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("search file with %s, in name:%s\n", keyword, f.name)
}
