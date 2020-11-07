package main
import(
	"fmt"
)

func main(){
	fmt.Println("Start")
	taskManager := taskManager{
		tasks : make(map[string]task),
	}
	taskManager.initTaskManager("data/todolist.txt")
	
	firstTask := task{}
	firstTask.setTask("eat","DOING")
	taskManager.addTask(firstTask)

	secondTask := task{}
	secondTask.setTask("clean","TODO")
	taskManager.addTask(secondTask)	

	thirsdTask := task{}
	thirsdTask.setTask("run","DID")
	taskManager.addTask(thirsdTask)	
	
	taskManager.removeTask(firstTask.name)

	taskManager.updateStatus(secondTask,"DOING")	
		
	taskManager.saveAllTask()

	displayTaks(taskManager)

	fmt.Println("Finish")
}

func displayTaks(tm taskManager){
	content := tm.displayAllTask()
	fmt.Println(content)
}

//TODO to delete 
func check(tm taskManager){
	fmt.Println("check")
	for _, task := range tm.tasks{
		fmt.Println(task.name, task.status)
	}
	fmt.Println("fin du check")
}