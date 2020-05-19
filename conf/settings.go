package settings


type databaseSettings struct {
	dbURL string
	dbName string
}

var DatabaseSettings databaseSettings

func init() {
	DatabaseSettings.dbName = ""
}