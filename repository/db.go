package repository


// Connection is the database connection interface
type Connection interface {
	URL string
	DBName string
	Connect()
	Disconnect()
	GetCollection()

	Find()
	Get()
	Delete() error
}

