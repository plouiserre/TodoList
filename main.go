package main
import(
	"fmt"
)

func main(){
	dataManagement := file{
		nameFile : "data/todolist.txt",
	}

	dataManagement.saveDataFile("todolist saved!!")

	contentFile := dataManagement.loadContentFile()

	fmt.Println("Display loaded file ", contentFile)
}