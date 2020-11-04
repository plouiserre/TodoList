package main 

type task struct {
	name string 
	status string
}

func (t *task) setTask(name string, status string){
	t.name = name
	t.updateStatus(status)
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