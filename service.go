package todogo

type Service interface {
	Add()
	FindOne()
	FindAll()
	Count()
	Update()
	Delete()
}