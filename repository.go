package main

type TaskRepository interface {
	Save(task Task)
	Update(task Task, id string)
	Delete(task Task)
}

type ProductRepository interface {
	Save(product Product)
	Update(product Product, id string)
	Delete(product Product)
}

type ProductRepositoryImpl struct {
	//aqqq11111qqqqvvvvv
}

func Save(product Product) {
	//db.query("INSERT INTO p")
}

func (p Product) Save(product Product) {

}
