package composite

import "fmt"

type Component interface {
	List()
}

type File struct {
	Name string
}

func (f *File) List() {
	fmt.Printf("File: %s\n", f.Name)
}

type Directory struct {
	Name       string
	Components []Component
}

func (d *Directory) List() {
	fmt.Printf("Directory: %s\n", d.Name)
	for _, component := range d.Components {
		component.List()
	}
}

func (d *Directory) Add(component Component) {
	d.Components = append(d.Components, component)
}

func (d *Directory) Remove(component Component) {
	var updatedComponents []Component
	for _, c := range d.Components {
		if c != component {
			updatedComponents = append(updatedComponents, c)
		}
	}
	d.Components = updatedComponents
}

func GetFileDirectories() {
	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}
	file3 := &File{Name: "file3.txt"}

	dir1 := &Directory{Name: "dir1"}
	dir2 := &Directory{Name: "dir2"}

	dir1.Add(file1)
	dir1.Add(file2)
	dir2.Add(file3)
	dir2.Add(dir1)
	dir2.List()
	dir2.Remove(file3)
	dir2.List()
}
