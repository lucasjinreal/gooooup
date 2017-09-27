package main

import (
	"log"
	"fmt"
	"os"
	"encoding/gob"
	"gitlab.com/jinfagang/colorgo"
	"path/filepath"
)


var hisGobPath = filepath.Join("./", "test.gob")
func Exists(path string) bool{
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		// got error, panic it
		panic(err)
	}
	return true
}

type Record struct {
	FilePath string
	Url string
}

func checkError(err error) {
	if err != nil {
		cg.PrintlnRed(err.Error())
	}
}

func readFromLocal() []Record{

	data := []Record{}
	hisGob, err := os.Open(hisGobPath)
	checkError(err)

	dataDecoder := gob.NewDecoder(hisGob)
	err = dataDecoder.Decode(&data)
	checkError(err)
	return data
}

func saveToLocal(data []Record) {
	if Exists(hisGobPath) {
		hisGob, err := os.Open(hisGobPath)
		checkError(err)
		defer hisGob.Close()

		dataEncoder := gob.NewEncoder(hisGob)
		gob.Register(Record{})
		dataEncoder.Encode(data)
		fmt.Println("every time save to local: ", data)
	} else {
		hisGob, err := os.Create(hisGobPath)
		checkError(err)
		defer hisGob.Close()

		dataEncoder := gob.NewEncoder(hisGob)
		gob.Register(Record{})

		dataEncoder.Encode(data)
	}
}

func SaveToHistory(newRecord Record) {
	// insert a new string to local

	if Exists(hisGobPath) {
		data := readFromLocal()
		data = append(data, newRecord)
		saveToLocal(data)
	} else {
		data := []Record{newRecord}
		saveToLocal(data)
	}
}

func main() {
	fmt.Println("Test")

	data := []Record{Record{"eee", "baidu.com"},
	Record{"rrrr", "tencent.com"}}

	SaveToHistory(data[0])
	fmt.Println("saved!")


	fmt.Println("================")
	SaveToHistory(Record{"bbbb", "fuck.com"})
	data = readFromLocal()
	fmt.Println(data)

	fmt.Println("================")
	SaveToHistory(Record{"ppppppp", "fuck444444.com"})
	data = readFromLocal()
	fmt.Println(data)



}
