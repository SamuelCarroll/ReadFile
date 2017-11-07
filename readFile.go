package readFile

//package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Wine basic structure of the wine interface type
type Wine struct {
	Class        int
	FeatureSlice []interface{}
}

func handleLine(line string) *Wine {
	var newWine Wine
	items := strings.Split(line, ",")

	newWine.Class, _ = strconv.Atoi(items[0])

	for i := 1; i < len(items); i++ {
		fval, err1 := strconv.ParseFloat(items[i], 64)
		bval, err2 := strconv.ParseBool(items[i])

		//try converting to s Float then a bool finally accept a string type
		if err1 == nil {
			newWine.FeatureSlice = append(newWine.FeatureSlice, fval)
		} else if err2 == nil {
			newWine.FeatureSlice = append(newWine.FeatureSlice, bval)
		} else {
			newWine.FeatureSlice = append(newWine.FeatureSlice, items[i])
		}
	}

	return &newWine
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Read reads in a file of the Wine dataset
func Read(inFile string) []*Wine {
	//func main () {
	dat, err := ioutil.ReadFile(inFile)
	check(err)

	sDat := fmt.Sprintf("%s", dat)
	datLines := strings.Split(sDat, "\n")

	var wines []*Wine
	for _, line := range datLines {
		if line != "" {
			temp := handleLine(line)
			wines = append(wines, temp)
		}
	}

	return wines
}
