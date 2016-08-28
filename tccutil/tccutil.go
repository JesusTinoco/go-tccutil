package tccutil

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"strings"

	version "github.com/hashicorp/go-version"
	// Required to access to the TCC.db sqlite3 database
	_ "github.com/mattn/go-sqlite3"
)

var ttc = "/Library/Application Support/com.apple.TCC/TCC.db"

// Client ...
type Client struct {
	Name string
}

func getOSXVersion() string {
	cmd := exec.Command("sw_vers", "-productVersion")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(out.String())
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

// ListClients ...
func ListClients() []Client {
	db := initDB(ttc)
	defer db.Close()

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
	return result
}

// GetClient ...
func GetClient() {
}

func getQueryToInsertClient(client string) string {
	localOSXVersion, err := version.NewVersion(getOSXVersion())
	if err != nil {
		log.Fatal(err)
	}
	osxVersion, err := version.NewVersion("10.11")
	if err != nil {
		log.Fatal(err)
	}

	var clientType int
	if string(client[0]) == "/" {
		clientType = 1
	} else {
		clientType = 0
	}

	var query string
	if localOSXVersion.Equal(osxVersion) || localOSXVersion.GreaterThan(osxVersion) {
		query = fmt.Sprintf("INSERT or REPLACE INTO access VALUES('kTCCServiceAccessibility','%s',%d,1,1,NULL,NULL)", client, clientType)
	} else {
		query = fmt.Sprintf("INSERT or REPLACE INTO access VALUES('kTCCServiceAccessibility','%s',%d,1,1,NULL)", client, clientType)
	}
	return query
}

// InsertClient ...
func InsertClient(client string) {
	db := initDB(ttc)
	defer db.Close()

	query := getQueryToInsertClient(client)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// RemoveClient ...
func RemoveClient(client string) {
	db := initDB(ttc)
	defer db.Close()

	query := fmt.Sprintf("DELETE from access where client IS '%s'", client)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// EnableClient ...
func EnableClient(client string) {
	db := initDB(ttc)
	defer db.Close()

	query := fmt.Sprintf("UPDATE access SET allowed='1' WHERE client='%s'", client)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// DisableClient ...
func DisableClient(client string) {
	db := initDB(ttc)
	defer db.Close()

	query := fmt.Sprintf("UPDATE access SET allowed='0' WHERE client='%s'", client)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
