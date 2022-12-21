package web_project

type Database struct {
	path     string
	accounts map[string]string
}

func InitDb() *Database {
	db := new(Database)

	db.path = "database/accounts.json"
	db.accounts = make(map[string]string)

	return db
}

/*
func (d Database) CheckAccount(username, password string) bool {
	UsernameHash := sha256.Sum256([]byte(username))
	PasswordHash := sha256.Sum256([]byte(password))

	ExpectedUserHash := sha256.Sum256()
	ExpectedPassHash := sha256.Sum256()

	return false
}
*/
/*
func InitDatabase() {
	db, err := sql.Open("sqlite3", "../database/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS user(
	    username TEXT,
	    password TEXT
	    );
	`
	res, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatal(err)
		return
	}
	println(res)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	query := "INSERT INTO user(username, password) VALUES (?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

}
*/
