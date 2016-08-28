package tccutil

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const dbPath = "./test.db"

func setup() *sql.DB {
	ttc = dbPath
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}

	createTable := `
	CREATE TABLE access (service TEXT NOT NULL, client TEXT NOT NULL, client_type INTEGER NOT NULL, 
	allowed INTEGER NOT NULL, prompt_count INTEGER NOT NULL, csreq BLOB, 
	CONSTRAINT key PRIMARY KEY (service, client, client_type));
	`
	_, err2 := db.Exec(createTable)
	if err2 != nil {
		panic(err2)
	}

	query := "INSERT INTO access VALUES('kTCCServiceAccessibility','com.github.atom',0,1,1,NULL)"
	_, err3 := db.Exec(query)
	if err != nil {
		log.Fatal(err3)
	}

	return db
}

func teardown(db *sql.DB) {
	db.Close()
	os.Remove(dbPath)
}

type ClientTest struct {
	Allowed int
}

func TestTCCUtil_InsertClient(t *testing.T) {
	db := setup()
	defer teardown(db)

	InsertClient("com.microsoft.VSCode")

	query := `
	SELECT client from access
    `
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Client
	for rows.Next() {
		item := Client{}
		err2 := rows.Scan(&item.Name)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}

	if result[1].Name != "com.microsoft.VSCode" {
		t.Fatalf("Should be 'com.microsoft.VSCode'")
	}

	if len(result) != 2 {
		t.Fatalf("Should contain 2 entries")
	}

}

func TestTCCUtil_ListClients(t *testing.T) {
	db := setup()
	defer teardown(db)

	clients := ListClients()
	if clients[0].Name != "com.github.atom" {
		t.Errorf("Should be 'com.github.atom'")
	}
	if len(clients) != 1 {
		t.Errorf("Should contain 1 client")
	}
}

func TestTCCUtil_DisableEnableClient(t *testing.T) {
	db := setup()
	defer teardown(db)

	DisableClient("com.github.atom")

	query := `
	SELECT allowed from access
    `
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []ClientTest
	for rows.Next() {
		item := ClientTest{}
		err2 := rows.Scan(&item.Allowed)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}

	if result[0].Allowed != 0 {
		t.Fatalf("Should be disabled")
	}

	EnableClient("com.github.atom")

	rows, err2 := db.Query(query)
	if err2 != nil {
		panic(err)
	}
	defer rows.Close()

	var result2 []ClientTest
	for rows.Next() {
		item := ClientTest{}
		err2 := rows.Scan(&item.Allowed)
		if err2 != nil {
			panic(err2)
		}
		result2 = append(result2, item)
	}

	if result2[0].Allowed != 1 {
		t.Fatalf("Should be enabled")
	}

}

func TestTCCUtil_RemoveClient(t *testing.T) {
	db := setup()
	defer teardown(db)

	RemoveClient("com.github.atom")

	query := `
	SELECT client from access
    `
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Client
	for rows.Next() {
		item := Client{}
		err2 := rows.Scan(&item.Name)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}

	if len(result) != 0 {
		t.Fatalf("Shouldn't have any entry")
	}

}
