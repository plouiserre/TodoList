package main
import(
	"fmt"
)

func main(){
	//for the moment all dataManagement files are comments because it will be use later
	/*dataManagement := file{
		nameFile : "data/todolist.txt",
	}

	dataManagement.saveDataFile("todolist saved!!")

	contentFile := dataManagement.loadContentFile()

	fmt.Println("Display loaded file ", contentFile)*/

	fmt.Println("Start")
	taskManager := taskManager{
		tasks : make(map[string]task),
	}

	firstTask := task{}
	firstTask.setTask("eat","DOING")
	taskManager.addTask(firstTask)

	secondTask := task{}
	secondTask.setTask("clean","TODO")
	taskManager.addTask(secondTask)	

	displayTaks(taskManager)

	taskManager.removeTask(firstTask.name)

	displayTaks(taskManager)

	taskManager.updateStatus(secondTask,"DOING")

	displayTaks(taskManager)

	fmt.Println("Finish")
}

func displayTaks(tm taskManager){
	for _, task := range tm.tasks{
		fmt.Println(task.name, task.status)
	}
}