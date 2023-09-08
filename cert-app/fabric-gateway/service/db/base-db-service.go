package db

type BaseDBService interface {
	Connect(host string, user string, password string, dbname string, port string) error
}
