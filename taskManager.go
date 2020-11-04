package main

type taskManager struct {
	tasks map[string]task
}

func (tm *taskManager) addTask(t task){
	isExist := tm.isElementExist(t.name)
	if isExist == false {
		tm.tasks[t.name] = t
	}
}

func (tm *taskManager) updateStatus(t task, status string){
	t.updateStatus(status)
	tm.removeTask(t.name)
	tm.addTask(t)
}

func (tm *taskManager) isElementExist(name string) bool{
	isExist := false
	t := tm.tasks[name]
	if (t == task{}){
		isExist = false
	}
	return isExist
}

func (tm *taskManager) removeTask(name string){
	delete(tm.tasks, name)
}
