package main

import "fmt"

type Component interface {
	Search(string)
}

type File struct {
	Name string
}

func (f File) Search(keyword string) {
	fmt.Printf("Search for keyword: %s in file: %s\n", keyword, f.Name)
}

type Folder struct {
	Name       string
	components []Component
}

func (f Folder) Search(keyword string) {
	fmt.Printf("Search recursively for keyword: %s in folder %s\n", keyword, f.Name)
	for _, c := range f.components {
		c.Search(keyword)
	}
}

func (f *Folder) Add(component Component) {
	f.components = append(f.components, component)
}

func main() {
	fileA := &File{Name: "File A"}
	fileB := &File{Name: "File B"}
	fileC := &File{Name: "File C"}

	folderA := &Folder{Name: "Folder A"}
	folderA.Add(fileA)

	folderB := &Folder{Name: "Folder B"}
	folderB.Add(fileB)
	folderB.Add(fileC)
	folderB.Add(folderA)

	folderB.Search("apple")
}
