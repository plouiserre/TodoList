package main
import(
	"bufio"
	"fmt"
	"strings"
	"os"
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
	
	fmt.Println("Bonjour humain. Je suis une IA super intelligente pour t'aider à faire quelques choses de ta misérable vie!!")
	
	taskManager.initTaskManager("data/todolist.txt")

	taskManager.fileManager.flushFile()
	i:= 0

	for i<1{
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Que veux-tu faire? Choisis la bonne réponse")
		fmt.Println("1 : Créer une tâche en lui donnant un simple nom")
		fmt.Println("2 : Créer une catégorie en lui donnant un simple nom")
		fmt.Println("3 : Tout sauvegarder")
		fmt.Println("4 : Afficher le contenu de la sauvegarde")
		fmt.Println("5 : Supprimer une tâche")
		fmt.Println("6 : Supprimer une catégorie")
		fmt.Println("7 : Te barrer car tu as peur de moi")

		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n","",-1)

		if strings.Compare("1", answer) == 0{
			createTask(reader, categoryManager, taskManager)
		} else if strings.Compare("2", answer ) == 0 {
			createCategory(reader, categoryManager)
		} else if strings.Compare("3", answer) == 0{
			SaveDatas(taskManager, categoryManager)
		} else if strings.Compare("4", answer) == 0{
			displayEntities(taskManager)
		}else if strings.Compare("5", answer) == 0{
			deleteTask(reader,taskManager)
		}else if strings.Compare("6", answer) == 0{
			deleteCategory(reader,categoryManager,taskManager)
		}else if strings.Compare("7", answer) == 0{
			fmt.Println("Terminé tout le monde descend!!!")
			break;
		}else{
			fmt.Println("Abruti ta réponse ne fait pas parti des choix, on recommence tout.")
		}
	}
	fmt.Println("Finish process")	
}

func createTask(reader *bufio.Reader, categoryManager categoryManager, taskManager taskManager){
	fmt.Println("Merci de choisir la catégorie de la future tâche")
	categoriesByName  := make(map[string] category)
	categories := categoryManager.categories
	categoriesCount := len(categories)
	if(categoriesCount > 0){
		categoryNewTask := category{}
		for i, category := range categories{
			categoriesByName[category.name] = category
			fmt.Println(i,": ", category.name)
		}

		categoryName, _ := reader.ReadString('\n')
		categoryName = strings.Replace(categoryName, "\n","",-1)
		categoryNewTask = categoriesByName[categoryName]
		if(categoryNewTask != (category{}) ){
			fmt.Println("La catégorie ",categoryName," a été choisie")

			fmt.Println("Merci de choisir le nom pour ta tâche")
			taskName, _ := reader.ReadString('\n')
			taskName = strings.Replace(taskName, "\n","",-1)
			
			customizeTask := task{}	
			customizeTask.setTask(taskName,"TODO",categoryNewTask)
			taskManager.addTask(customizeTask)	
		}else {
			fmt.Println(categoryName," n'existe pas merci de choisir une catégorie qui existe la prochaine fois")
		}
	} else {
		fmt.Println("Aucune catégorie n'a été créé merci de le faire et de revenir après dans la création de tâche")
	}	
}

func deleteTask(reader *bufio.Reader, taskManager taskManager){
	fmt.Println("Merci de choisir la tâche à supprimer")
	tasksByName  := make(map[string] task)
	tasks := taskManager.tasks
	tasksCount := len(tasks)
	if(tasksCount > 0){
		taskToDelete := task{}
		for i, task := range tasks{
			tasksByName[task.name] = task
			fmt.Println(i,": ", task.name)
		}

		taskToDeleteName, _ := reader.ReadString('\n')
		taskToDeleteName = strings.Replace(taskToDeleteName, "\n","",-1)
		taskToDelete = tasksByName[taskToDeleteName]
		if(taskToDelete != (task{}) ){
			fmt.Println("La task ",taskToDeleteName," va être supprimée")			
			taskManager.removeTask(taskToDelete.name)
		}else {
			fmt.Println(taskToDeleteName," n'existe pas merci de choisir une tâche qui existe pour la supprimer")
		}
	} else {
		fmt.Println("Aucune tâche n'a été créé merci de le faire et de revenir après dans la suppression de tâche")
	}	
}

func createCategory(reader *bufio.Reader, categoryManager categoryManager){
	fmt.Println("Merci de choisir le nom pour ta catégorie")
	categoryName, _ := reader.ReadString('\n')
	categoryName = strings.Replace(categoryName, "\n","",-1)

	ownCategory := category{}
	ownCategory.setCategoryName(categoryName)
	
	//Attendre juste après
	categoryManager.addCategory(ownCategory)
}

func deleteCategory(reader *bufio.Reader, categoryManager categoryManager, taskManager taskManager){
	fmt.Println("Merci de choisir la catégorie à supprimer")
	categoriesByName := make(map[string] category)
	categories := categoryManager.categories
	categoriesCount := len(categories)
	if(categoriesCount > 0){
		categoryToDelete := category{}
		for i, category := range categories{
			categoriesByName[category.name] = category
			fmt.Println(i, ": ",category.name)
		}
		categoryToDeleteName, _ := reader.ReadString('\n')
		categoryToDeleteName = strings.Replace(categoryToDeleteName, "\n","",-1)
		categoryToDelete = categoriesByName[categoryToDeleteName]
		if(categoryToDelete != (category{})){
			result := categoryManager.removeCategory(taskManager.tasks, categoryToDelete.name)
			if(result == true){
				fmt.Println(categoryToDelete.name, " a été supprimée")
			}else {
				fmt.Println(categoryToDelete.name, " n'a été supprimée car elle possède encore des tâches")
			}
		}else {
			fmt.Println(categoryToDeleteName," n'existe pas merci de choisir une catégorie existante pour la supprimée")
		}		
	}else {
		fmt.Println("Aucune catégorie n'a été créé merci de le faire et de revenir après dans la suppression de catégorie")
	}
}

func SaveDatas(taskManager taskManager, categoryManager categoryManager){
	taskManager.saveAllTasks()
	categoryManager.saveAllCategories()
}

func displayEntities(tm taskManager){
	content := tm.fileManager.displayAllTask()
	fmt.Println(content)
}

//TODO to delete 
func check(tm taskManager){
	fmt.Println("check")
	for _, task := range tm.tasks{
		fmt.Println(task.name, task.status," ", task.category.name)
	}
	fmt.Println("fin du check")
}

func displayTasksSpecificStatus(tm taskManager, status string){
	fmt.Println("Affichage des tâches du statut ", status)
	tasks := tm.getTasksByStatus(status)
	for _, task := range tasks{
		fmt.Println(task.name," ", task.status," ", task.category.name)
	}
	fmt.Println("Fin des tâches")
}

func displayTasksSpecificCategory(tm taskManager, taskCategory category){
	fmt.Println("Affichage des tâches de la catégorie ", taskCategory)
	tasks := tm.getTasksByCategory(taskCategory)
	for _, task := range tasks{
		fmt.Println(task.name," ", task.status," ", task.category.name)
	}
	fmt.Println("Fin des tâches")
}