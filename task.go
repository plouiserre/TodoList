package main 

type task struct {
	name string 
	status string
	category category
}

func (t *task) setTask(name string, status string, taskCategory category){
	t.name = name
	t.updateStatus(status)
	t.category = taskCategory
}

func (t *task) updateStatus (status string){
	isValid := isValidStatus(status)
	if isValid{
		t.status = status
	}
}

func isValidStatus(statusToCheck string) bool{
	result := false
	if statusToCheck == "TODO" || statusToCheck == "DOING" || statusToCheck == "DID" || statusToCheck == "CANCEL"{
		result = true
	}	
	return result
}