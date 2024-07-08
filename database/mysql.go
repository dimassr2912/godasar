package database

var connection string

func init() { // Akan langsung diakses karena package diakses
	connection = "Mysql"
}

func GetDatabase() string {
	return connection
}
