package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return &DatabasePostgreSQL{Name: "PostgreSQL"}
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return &DatabaseMongoDB{Name: "MongoDB"}
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongoDB}
}
