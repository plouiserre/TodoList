package main 

import(
	"fmt"
	"io/ioutil"
	"os"
)

type file struct{
	nameFile string
}

func (f file) loadContentFile() string{
	bs, err := ioutil.ReadFile(f.nameFile)
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	content := string(bs)
	return content
}

func (f file) saveDataFile(msg string){
	err := ioutil.WriteFile(f.nameFile, []byte(msg), 0666)
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}