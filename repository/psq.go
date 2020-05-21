package repository

//TODO mudar para o nome referente ao negocio
type User interface {
	Find()
	Get()
	Delete() error
}
