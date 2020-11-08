package main
import(
	"fmt"
)

func main(){
	fmt.Println("Start task")
	taskManager := taskManager{
		tasks : make(map[string]task),
	}

	taskManager.initTaskManager("data/todolist.txt")

	taskManager.fileManager.flushFile()
	
	firstTask := task{}
	firstTask.setTask("eat","DOING")
	taskManager.addTask(firstTask)

	secondTask := task{}
	secondTask.setTask("code","TODO")
	taskManager.addTask(secondTask)	

	thirsdTask := task{}
	thirsdTask.setTask("run","DID")
	taskManager.addTask(thirsdTask)	
	
	taskManager.removeTask(firstTask.name)

	taskManager.updateStatus(secondTask,"DOING")	
		
	taskManager.saveAllTasks()

	fmt.Println("Finish task")

	fmt.Println("Start category")
	categoryManager := categoryManager{
		categories : make(map[string]category),
	}

	categoryManager.initCategoryManager("data/todolist.txt")
	
	firstCategory := category{}
	firstCategory.setCategoryName("FOOD")
	categoryManager.addCategory(firstCategory)

	secondCategory := category{}
	secondCategory.setCategoryName("PHYSICAL")
	categoryManager.addCategory(secondCategory)	
	
	categoryManager.removeCategory(firstCategory.name)

	categoryManager.updateStatus(secondCategory,"VARIED")	
		
	categoryManager.saveAllCategories()	

	displayEntities(taskManager)

	fmt.Println("Finish task")
}

func displayEntities(tm taskManager){
	content := tm.fileManager.displayAllTask()
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