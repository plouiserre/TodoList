package main 

import(
	"fmt"
	"io/ioutil"
	"os"
)

type fileManager struct{
	nameFile string 
}

func (fm *fileManager) flushFile(){
	fm.saveDataFile("")
}

func (fm fileManager) loadContentFile() string{
	bs, err := ioutil.ReadFile(fm.nameFile)
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	content := string(bs)
	return content
}

func (fm fileManager) saveDataFile(msg string){
	err := ioutil.WriteFile(fm.nameFile, []byte(msg), 0666)
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func (fm *fileManager) displayAllTask() string{
	contentFile := fm.loadContentFile()
	return contentFile
}
