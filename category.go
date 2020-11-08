package main 

type category struct{
	name string
}

func (c *category) setCategoryName(categoryName string){
	c.name = categoryName
}