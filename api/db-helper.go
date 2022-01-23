package api

import (
	"database/sql"
	"log"
)

// initlaize DB
func (s *SchabloneServer) initializeDB() error {
	rows, _ := s.queryDB("SELECT count(*) FROM TemplateGroup")

	var value int
	for rows.Next() {
		err := rows.Scan(&value)
		if err != nil {
			return err
		}
	}
	// Create Admin Group
	if value < 1 {
		_, err := s.executeOnDB("INSERT INTO TemplateGroup(Name) values (?)", "admin")
		return err
	}
	return nil
}

// Execute request on DB
func (s *SchabloneServer) executeOnDB(prepateStatement string, args ...interface{}) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(prepateStatement)
	if err != nil {
		log.Printf("Error %s", err)
		return -1, err
	}
	data, err := stmt.Exec(args...)
	if err != nil {
		log.Printf("Error %s", err)
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	id, err := data.LastInsertId()
	return id, err
}

// Query DB
func (s *SchabloneServer) queryDB(queryStatement string, args ...interface{}) (*sql.Rows, error) {
	rows, err := s.db.Query(queryStatement, args...)
	if err != nil {
		log.Printf("Error %s", err)
		return nil, err
	}

	return rows, nil
}
