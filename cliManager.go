package main 

import(
	"bufio"
	"fmt"
	"strings"
	"os"
)

//TODO 
//corriger quand on sauvegarde de ne 
//corriger la suppression de tâches 

type cliManager struct{
	categoryManager categoryManager
	taskManager taskManager
}

func (cli *cliManager) initCliManager(){
	cli.categoryManager = categoryManager{
		categories : make(map[string]category),
	}
	cli.taskManager = taskManager{
		tasks : make(map[string]task),
	}

	fmt.Println("Start process")	

	cli.categoryManager.initCategoryManager("data/todolist.txt")

	fmt.Println("Bonjour humain. Je suis une IA super intelligente pour t'aider à faire quelques choses de ta misérable vie!!")
	
	cli.taskManager.initTaskManager("data/todolist.txt")
}

func (cli *cliManager) loopActions(){
	i:= 0

	for i<1{
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Que veux-tu faire? Choisis la bonne réponse")
		fmt.Println("1 : Créer une tâche en lui donnant un simple nom")
		fmt.Println("2 : Créer une catégorie en lui donnant un simple nom")
		fmt.Println("3 : Mettre à jour le status d'une tâche")
		fmt.Println("4 : Tout sauvegarder")
		fmt.Println("5 : Afficher le contenu de la sauvegarde")
		fmt.Println("6 : Supprimer une tâche")
		fmt.Println("7 : Supprimer une catégorie")
		fmt.Println("8 : Te barrer car tu as peur de moi")

		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n","",-1)

		if strings.Compare("1", answer) == 0{
			cli.createTask(reader)
		} else if strings.Compare("2", answer ) == 0 {
			cli.createCategory(reader)
		} else if strings.Compare("3", answer) == 0{
			cli.updateStatusTask(reader)
		} else if strings.Compare("4", answer) == 0{
			cli.SaveDatas()
		} else if strings.Compare("5", answer) == 0{
			cli.displayEntities()
		}else if strings.Compare("6", answer) == 0{
			cli.deleteTask(reader)
		}else if strings.Compare("7", answer) == 0{
			cli.deleteCategory(reader)
		}else if strings.Compare("8", answer) == 0{
			fmt.Println("Terminé tout le monde descend!!!")
			break;
		}else{
			fmt.Println("Abruti ta réponse ne fait pas parti des choix, on recommence tout.")
		}
	}
	fmt.Println("Finish process")	
}

func (cli *cliManager)  createTask(reader *bufio.Reader){
	fmt.Println("Merci de choisir la catégorie de la future tâche")
	categoryNewTask  := cli.findCategory(reader)
	if(categoryNewTask != category{}){
		if(categoryNewTask != (category{}) ){
			fmt.Println("La catégorie ",categoryNewTask.name," a été choisie")

			fmt.Println("Merci de choisir le nom pour ta tâche")
			taskName, _ := reader.ReadString('\n')
			taskName = strings.Replace(taskName, "\n","",-1)
			
			customizeTask := task{}	
			customizeTask.setTask(taskName,"TODO",categoryNewTask)
			cli.taskManager.addTask(customizeTask)	
		}else {
			fmt.Println(categoryNewTask.name," n'existe pas merci de choisir une catégorie qui existe la prochaine fois")
		}
	} else {
		fmt.Println("Aucune catégorie n'a été créé merci de le faire et de revenir après dans la création de tâche")
	}	
}

func (cli *cliManager)  deleteTask(reader *bufio.Reader){
	fmt.Println("Merci de choisir la tâche à supprimer")
	taskToDelete := cli.findTask(reader)
	if(taskToDelete != task{}){
		if(taskToDelete != (task{}) ){
			fmt.Println("La task ",taskToDelete.name," va être supprimée")			
			cli.taskManager.removeTask(taskToDelete.name)
		}else {
			fmt.Println(taskToDelete.name," n'existe pas merci de choisir une tâche qui existe pour la supprimer")
		}
	} else {
		fmt.Println("Aucune tâche n'a été créé merci de le faire et de revenir après dans la suppression de tâche")
	}	
}

func (cli *cliManager) updateStatusTask(reader *bufio.Reader){
	fmt.Println("Quelle est la tâche à mettre à jour")
	taskToUpdate:= cli.findTask(reader)
	if(taskToUpdate != task{}){
		if(taskToUpdate != (task{})){
			fmt.Println("Quel est le nouveau status de la tâche")
			fmt.Println("TODO")
			fmt.Println("DOING")
			fmt.Println("DID")
			fmt.Println("CANCEL")
			newStatus, _ := reader.ReadString('\n')
			newStatus = strings.Replace(newStatus, "\n","",-1)
			
			cli.taskManager.updateStatus(taskToUpdate, newStatus)		
		}
	} else {
		fmt.Println("Ta saisie ne correspond à aucune tâche existante crétin")
	}
}

func (cli *cliManager) findTask(reader *bufio.Reader) task{
	tasks := cli.taskManager.tasks
	tasksCount := len(tasks)
	taskToFind := task{}
	tasksByName  := make(map[string] task)
	if(tasksCount > 0){
		for i, task := range tasks{
			tasksByName[task.name] = task
			fmt.Println(i,": ", task.name)
		}
		taskToFindName, _ := reader.ReadString('\n')
		taskToFindName = strings.Replace(taskToFindName, "\n","",-1)
		taskToFind = tasksByName[taskToFindName]
	}
	return taskToFind
}

func (cli *cliManager)  createCategory(reader *bufio.Reader){
	fmt.Println("Merci de choisir le nom pour ta catégorie")
	categoryName, _ := reader.ReadString('\n')
	categoryName = strings.Replace(categoryName, "\n","",-1)

	ownCategory := category{}
	ownCategory.setCategoryName(categoryName)
	
	//Attendre juste après
	cli.categoryManager.addCategory(ownCategory)
}

func (cli *cliManager)  deleteCategory(reader *bufio.Reader){
	fmt.Println("Merci de choisir la catégorie à supprimer")
	categoryToDelete  := cli.findCategory(reader)
	if(categoryToDelete != category{}){
		if(categoryToDelete != (category{})){
			result := cli.categoryManager.removeCategory(cli.taskManager.tasks, categoryToDelete.name)
			if(result == true){
				fmt.Println(categoryToDelete.name, " a été supprimée")
			}else {
				fmt.Println(categoryToDelete.name, " n'a été supprimée car elle possède encore des tâches")
			}
		}else {
			fmt.Println(categoryToDelete.name," n'existe pas merci de choisir une catégorie existante pour la supprimée")
		}		
	}else {
		fmt.Println("Aucune catégorie n'a été créé merci de le faire et de revenir après dans la suppression de catégorie")
	}
}

func (cli *cliManager) findCategory(reader *bufio.Reader) category{
	categoryToFind := category{}
	categoriesByName  := make(map[string] category)
	categories := cli.categoryManager.categories
	categoriesCount := len(categories)
	if(categoriesCount > 0){
		for i, category := range categories{
			categoriesByName[category.name] = category
			fmt.Println(i,": ", category.name)
		}

		categoryName, _ := reader.ReadString('\n')
		categoryName = strings.Replace(categoryName, "\n","",-1)
		categoryToFind = categoriesByName[categoryName]
	}
	return categoryToFind
}

func (cli *cliManager) SaveDatas(){
	cli.taskManager.fileManager.flushFile()
	cli.taskManager.saveAllTasks()
	cli.categoryManager.saveAllCategories()
}

func (cli *cliManager) displayEntities(){
	content := cli.taskManager.fileManager.displayAllTask()
	fmt.Println(content)
}