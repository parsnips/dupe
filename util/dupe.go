package util

import (
	"bytes"
	"io/ioutil"
)

//Dupe checks for duplicate files
type Dupe interface {
	IsDupe(filePathA, filePathB string) bool
}

type RealDupeChecker struct{}

func (d *RealDupeChecker) IsDupe(a, b string) bool {
	fileA, err := ioutil.ReadFile(a)
	if err != nil {
		panic(err)
	}
	fileB, err := ioutil.ReadFile(b)
	if err != nil {
		panic(err)
	}

	return bytes.Compare(fileA, fileB) == 0
}

type StaticDupe struct {
	ReturnValue bool
}

func (d *StaticDupe) IsDupe(a, b string) bool {
	return d.ReturnValue
}
