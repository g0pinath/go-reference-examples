package main

import (
	"database/sql"
	"fmt"
	"log"
    "os"
    _ "github.com/lib/pq"
)


type User struct {
	ID       int
	Email    string
	Password string
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

	DB_DSN := os.Getenv("DB_DSN")
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
    err = db.Ping()
    CheckError(err)
    fmt.Println("Successfully connected!")

    // insert
    // hardcoded
    insertStmt := `insert into "students"("name", "roll") values('John', 3)`
    _, e := db.Exec(insertStmt)
    CheckError(e)
 
    // dynamic
    insertDynStmt := `insert into "students"("name", "roll") values($1, $2)`
    _, e = db.Exec(insertDynStmt, "Jane", 4)
    CheckError(e)

	defer db.Close()

}
