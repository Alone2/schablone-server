package api

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	// Create Admin Group and User
	if value < 1 {
		templateId, err := s.executeOnDB("INSERT INTO TemplateGroup(Name) values (?)", "admin")
		if err != nil {
			return err
		}
		userId, err := s.executeOnDB("INSERT INTO User(Firstname, Lastname, Username, Password) values (?,?,?,?)", "Admin", "Admin", "admin", "$2a$10$TSNVLUrKmA4vIG24w7I0wugKCPkSs.7M6E1R9iiZz6v1dHEWaeQ4e")
		if err != nil {
			return err
		}
		_, err = s.executeOnDB("INSERT INTO User_TemplateGroup(BelongsToUser, TemplateGroup, UserHasWriteAccess, UserHasUserModifyAccess) values (?,?,?,?)", userId, templateId, 1, 1)
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

func (s *SchabloneServer) createUser(username string, firstname string, lastname string, password string) (int64, error) {
	// Hash password
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error %s", err)
		return 0, err
	}
	hashedPasswordString := string(hashedPassword)

	// Add user to DB
	userId, err := s.executeOnDB("INSERT INTO User(Firstname, Lastname, Username, Password) values (?,?,?,?)", firstname, lastname, username, hashedPasswordString)
	if err != nil {
		log.Printf("Error %s", err)
		return 0, err
	}

	return userId, err
}

// returns API Token
func (s *SchabloneServer) verifyUser(username string, password string) (string, error) {
	// Get password out of database
	rows, err := s.queryDB("SELECT Id, Password FROM User WHERE Username=?", username)
	if err != nil {
		log.Printf("Error %s", err)
		return "", err
	}
	var hashedPassword string
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId, &hashedPassword)
		if err != nil {
			log.Printf("Error %s", err)
			return "", err
		}
	}

	// Compare to hash
	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(hashedPassword)
	err = bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
	if err != nil {
		// Password incorrect
		log.Printf("Password incorrect")
		return "", err
	}

	// Generate API Key
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 40)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	apiKey := string(b)
	log.Println(apiKey)

	nowTime := time.Now().Unix()
	_, err = s.executeOnDB("INSERT INTO ActiveAPIKey(BelongsToUser, CreationDate, Content) values (?,?,?)", userId, nowTime, apiKey)
	if err != nil {
		log.Printf("Error %s", err)
		return "", err
	}

	return apiKey, nil
}

// verifies API Token
func (s *SchabloneServer) verifyAPIToken(apiToken string) (bool, error) {
	rows, err := s.queryDB("SELECT count(*) FROM ActiveAPIKey WHERE Content=?", apiToken)
	if err != nil {
		log.Printf("Error %s", err)
		return false, err
	}

	var value int
	for rows.Next() {
		err := rows.Scan(&value)
		if err != nil {
			log.Printf("Error %s", err)
			return false, err
		}
	}
	fmt.Println(value)
	if value >= 1 {
		return true, nil
	}
	return false, nil
}

func (s *SchabloneServer) userHasWriteAccessTo(userId int, userId2 int) (bool, error) {
}

func (s *SchabloneServer) userHasWriteAccessTo(groupId int) (bool, error) {
}

func (s *SchabloneServer) userHasReadAccessTo(userId int, userId2 int) (bool, error) {
}

func (s *SchabloneServer) userHasReadAccessTo(groupId int, userId2 int) (bool, error) {
}

func (s *SchabloneServer) getTemplateGroup(templateId int) int {
}

func (s *SchabloneServer) getMacroGroup(templateId int) int {
}
