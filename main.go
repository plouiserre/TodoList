package main
import(
	"fmt"
)

func main(){
	categoryManager := categoryManager{
		categories : make(map[string]category),
	}
	taskManager := taskManager{
		tasks : make(map[string]task),
	}
	fmt.Println("Start process")	

	categoryManager.initCategoryManager("data/todolist.txt")
	
	firstCategory := category{}
	firstCategory.setCategoryName("FOOD")
	categoryManager.addCategory(firstCategory)

	secondCategory := category{}
	secondCategory.setCategoryName("PHYSICAL")
	categoryManager.addCategory(secondCategory)	

	thirdCategory := category{}
	thirdCategory.setCategoryName("COMPUTING")
	categoryManager.addCategory(thirdCategory)	
	
	taskManager.initTaskManager("data/todolist.txt")

	taskManager.fileManager.flushFile()
	
	firstTask := task{}
	firstTask.setTask("eat","DOING",firstCategory)
	taskManager.addTask(firstTask)

	secondTask := task{}
	secondTask.setTask("code","TODO", thirdCategory)
	taskManager.addTask(secondTask)	

	thirsdTask := task{}
	thirsdTask.setTask("run","DID", secondCategory)
	taskManager.addTask(thirsdTask)	
	
	taskManager.removeTask(firstTask.name)

	taskManager.updateStatus(secondTask,"DOING")	
		
	taskManager.saveAllTasks()
	
	categoryManager.removeCategory(taskManager.tasks, firstCategory.name)
	categoryManager.removeCategory(taskManager.tasks, secondCategory.name)
		
	categoryManager.saveAllCategories()	

	fmt.Println("Finish process")

	displayEntities(taskManager)
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