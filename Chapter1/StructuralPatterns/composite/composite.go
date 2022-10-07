package main

import (
	"dsa-golang/Chapter1/StructuralPatterns/composite/exmaple"
	"fmt"
)

type IComposite interface {
	perform()
}

type Branch struct {
	name     string
	branches []Branch
	leafs    []Leaflet
}
type Leaflet struct {
	name string
}

func (leaftlet *Leaflet) perform() {
	fmt.Println("branch class: " + leaftlet.name)
}
func (branch *Branch) perform() {
	fmt.Println("branch class: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}
	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) addLeaf(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}
func (branch *Branch) addBranch(branchToadd Branch) {
	branch.branches = append(branch.branches, branchToadd)
}

func (brach *Branch) getLeaflets() []Leaflet {
	return brach.leafs
}
func main() {
	branch1 := &Branch{name: "branch one"}
	leaf1 := Leaflet{name: "leaf  one"}
	branch2 := &Branch{name: "branch two"}
	leaf2 := Leaflet{name: "leaf  two"}
	leaf3 := Leaflet{name: "leaf  three"}
	leaf4 := Leaflet{name: "leaf  four"}
	branch1.addLeaf(leaf1)
	branch1.addLeaf(leaf2)
	branch2.addLeaf(leaf3)
	branch2.addLeaf(leaf4)
	branch1.addBranch(*branch2)
	// branch1.perform()

	//exmaple
	file1 := &exmaple.File{Name: "Riri's file"}
	file2 := &exmaple.File{Name: "Kanye's file"}
	file3 := &exmaple.File{Name: "Hov's file"}

	folder := exmaple.Folder{Name: "All of the Lights"}
	folder.AddComponent(file1)
	folder.AddComponent(file2)
	folder.AddComponent(file3)

	folder.Search("credit card declined")

}
