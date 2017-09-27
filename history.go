package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"gitlab.com/jinfagang/colorgo"
	"encoding/gob"
	"strconv"
)

// Set Global Variable
var usr, _ = user.Current()
var configDir = filepath.Join(usr.HomeDir, ".config")
var hisGobPath = filepath.Join(configDir, "history.gob")


// every upload will save into a record
type Record struct {
	FilePath string
	Url string
}


func checkError(err error) {
	if err != nil {
		cg.PrintlnRed(err.Error())
	}
}

// Here read must set with concrete type
// TODO: this function why are the fucking every time only read 1 elem??????
func readFromLocal() []Record{

	data := []Record{}
	hisGob, err := os.Open(hisGobPath)
	checkError(err)

	dataDecoder := gob.NewDecoder(hisGob)
	err = dataDecoder.Decode(&data)
	checkError(err)
	fmt.Println("every time read from local: ", data)

	return data
}

func saveToLocal(data []Record) {
	if Exists(hisGobPath) {
		hisGob, err := os.Open(hisGobPath)
		checkError(err)
		defer hisGob.Close()

		dataEncoder := gob.NewEncoder(hisGob)
		dataEncoder.Encode(data)
		fmt.Println("every time save to local: ", data)
	} else {
		hisGob, err := os.Create(hisGobPath)
		checkError(err)
		defer hisGob.Close()

		dataEncoder := gob.NewEncoder(hisGob)
		dataEncoder.Encode(data)
	}
}

func SaveToHistory(newRecord Record) {
	// insert a new string to local

	if !Exists(configDir) {
		os.MkdirAll(configDir, 0777)
	}

	if Exists(hisGobPath) {
		data := readFromLocal()
		data = append(data, newRecord)
		saveToLocal(data)
	} else {
		data := []Record{newRecord}
		saveToLocal(data)
	}
}

func ShowHistory() {
	// print out the local saved urls
	// data can be any type
	if Exists(hisGobPath) {
		data := readFromLocal()
		for i, record := range data{
			cg.PrintYellow( strconv.Itoa(i+1) + ". ")
			cg.PrintYellow(record.Url)
			fmt.Print(" ")
			cg.PrintlnBlue(record.FilePath)
		}

	} else {
		cg.PrintlnYellow("no history find yet.")
	}

}
