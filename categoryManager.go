package main 

type categoryManager struct{
	categories map[string]category 
	fileManager fileManager
}

func(c *categoryManager) initCategoryManager(fileName string){
	c.fileManager = fileManager{
		nameFile : fileName,
	}
}

func (c *categoryManager) addCategory(cat category){
	isExist := c.isElementExist(cat.name)
	if isExist == false {
		c.categories[cat.name] = cat
	}
}

//TODO factorise with the same method in taskManager.go
func (c *categoryManager) isElementExist(name string) bool{
	isExist := false
	cat := c.categories[name]
	if (cat == category{}){
		isExist = false
	}
	return isExist
}

func (c *categoryManager) updateStatus(cat category, newCategoryName string){
	oldName := cat.name
	cat.setCategoryName(newCategoryName)
	c.removeCategory(oldName)
	c.addCategory(cat)
}

func (c *categoryManager) removeCategory(name string){
	delete(c.categories, name)
}

func(c *categoryManager) saveAllCategories(){
	contentToWrite := c.fileManager.loadContentFile()
	for _, category := range c.categories{
		if (contentToWrite==""){
			contentToWrite = category.name
		} else {
			contentToWrite += "\n"+category.name
		}
	}
	c.fileManager.saveDataFile(contentToWrite)
}
