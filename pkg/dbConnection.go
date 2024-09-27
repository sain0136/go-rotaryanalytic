package pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sain0136/go-rotaryanalytic/utils"
)

type sqlCredentials struct {
	username     string
	password     string
	hostname     string
	port         string
	databaseName string
}

// generateDsn generates a DSN string that can be used to connect to MySQL based on the fields of the sqlCredentials struct.
func (credentials sqlCredentials) generateDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", credentials.username, credentials.password, credentials.hostname, credentials.port, credentials.databaseName)
}

func ConnectToMySQL() (*sql.DB, error) {
	credentials, err := setCredentials()
	if err != nil {
		return nil, fmt.Errorf("failed to set credentials: %w", err)
	}
	dsn := credentials.generateDsn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	test_connection()
	return db, nil
}

func test_connection() bool {
	db, err := ConnectToMySQL()
	if err != nil {
		log.Println("failed to connect: %w", err)
		return false
	}
	defer db.Close()

	//Check Table Existence
	query := "SHOW TABLES LIKE 'sessions';"
	var tableName string
	err = db.QueryRow(query).Scan(&tableName)
	if err != nil {
		log.Println("Table 'sessions' does not exist.")
		return false
	}
	if tableName == "" {
		log.Println("Table 'sessions' does not exist.")
		return false
	}
	log.Println("Table 'sessions' exists.")

	// Retrieve a Session Record (Optional)
	query = "SELECT * FROM sessions LIMIT 1;"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("failed to retrieve session record: %v", err)
	}
	defer rows.Close()

	// Print the retrieved session record
	for rows.Next() {
		// Assuming the sessions table has columns: session_id, user_id, session_data
		var session_id int
		var full_name int
		var email string
		err := rows.Scan(&session_id, &full_name, &email)
		if err != nil {
			log.Printf("failed to scan session record: %v", err)
		}
		log.Printf("Session Record - ID: %d, UserID: %d, SessionData: %s", session_id, full_name, email)
	}
	return true
}

func setCredentials() (sqlCredentials, error) {
	keys := []string{"MYSQL_USERNAME", "MYSQL_PASSWORD", "MYSQL_HOSTNAME", "MYSQL_PORT", "MYSQL_DATABASE_NAME"}
	values := make(map[string]string) // store values returned from utils.GetConfig

	// Mapping from string to utils.envValue
	envValueMap := map[string]utils.EnvValue{
		"MYSQL_USER":     utils.MYSQL_USER,
		"MYSQL_PASSWORD": utils.MYSQL_PASSWORD,
		"MYSQL_HOST":     utils.MYSQL_HOST,
		"MYSQL_PORT":     utils.MYSQL_PORT,
		"MYSQL_DB_NAME":  utils.MYSQL_DB_NAME,
	}

	for _, key := range keys {
		envKey, ok := envValueMap[key]
		if !ok {
			return sqlCredentials{}, fmt.Errorf("invalid config key: %s", key)
		}
		value, err := utils.GetConfig(envKey)
		if err != nil {
			return sqlCredentials{}, fmt.Errorf("failed to get config for %s: %w", key, err)
		}
		values[key] = value
	}

	credentials := sqlCredentials{
		username:     values["MYSQL_USER"],
		password:     values["MYSQL_PASSWORD"],
		hostname:     values["MYSQL_HOST"],
		port:         values["MYSQL_PORT"],
		databaseName: values["MYSQL_DB_NAME"],
	}

	return credentials, nil
}
