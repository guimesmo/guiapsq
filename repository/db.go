package repository

// Connection is the database connection interface
type Connection interface {
	Connect()
	Disconnect()
	GetCollection()

	Find()
	Get()
	Delete() error
}
