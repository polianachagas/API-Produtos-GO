package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// CONSTANTES QUE SÃO DADOS DA CONEXÃO
const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	//aqui estamos construindo a tring de conexão usando as constantes definidas acima
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//aqui a conexão com o postgres é aberta
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//ping para ver se a conexão terá sucesso, caso sim retorna o db, que é a conexão com o banco
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
