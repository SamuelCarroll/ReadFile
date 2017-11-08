package readFile

//package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/SamuelCarroll/DataTypes"
)

func handleLine(line string) *dataTypes.Data {
	var newData dataTypes.Data
	items := strings.Split(line, ",")

	newData.Class, _ = strconv.Atoi(items[0])

	for i := 1; i < len(items); i++ {
		fval, err1 := strconv.ParseFloat(items[i], 64)
		bval, err2 := strconv.ParseBool(items[i])

		//try converting to s Float then a bool finally accept a string type
		if err1 == nil {
			newData.FeatureSlice = append(newData.FeatureSlice, fval)
		} else if err2 == nil {
			newData.FeatureSlice = append(newData.FeatureSlice, bval)
		} else {
			newData.FeatureSlice = append(newData.FeatureSlice, items[i])
		}
	}

	return &newData
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Read reads in a specified file
func Read(inFile string) []*dataTypes.Data {
	//func main () {
	dat, err := ioutil.ReadFile(inFile)
	check(err)

	sDat := fmt.Sprintf("%s", dat)
	datLines := strings.Split(sDat, "\n")

	var lineData []*dataTypes.Data
	for _, line := range datLines {
		if line != "" {
			temp := handleLine(line)
			lineData = append(lineData, temp)
		}
	}

	return lineData
}
