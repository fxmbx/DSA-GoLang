package exmaple

import "fmt"

type Component interface {
	Search(string)
}

type File struct {
	Name string
}

func (f *File) Search(key string) {
	fmt.Printf("searching for : %s, in file: %s\n", f.Name, key)
}

type Folder struct {
	Name       string
	Components []Component
}

func (folder *Folder) Search(key string) {
	fmt.Println("folder search in pogress in folder: " + folder.Name)
	for _, component := range folder.Components {
		component.Search(key)
	}
}

func (folder *Folder) AddComponent(component Component) {
	folder.Components = append(folder.Components, component)
}
