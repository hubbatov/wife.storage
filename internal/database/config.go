package database

type DatabaseConfig struct {
	//Host - database host
	Host string
	//Port - database port
	Port int
	//User - database user
	User string
	//Password - database password
	Password string
	//Database - database db name
	Database string
	// Provider - database provider
	Provider string
}
