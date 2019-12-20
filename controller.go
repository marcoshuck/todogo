package todogo

type Controller interface {
	Create()
	Read()
	ReadAll()
	Update()
	Delete()
}