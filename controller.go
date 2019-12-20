package todogo

type Controller interface {
	Create()
	Read()
	Update()
	Delete()
}