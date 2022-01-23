package api

import (
	"database/sql"
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

func (s *SchabloneServer) editUser(username string, firstname string, lastname string, password string, userId int64) (int64, error) {
	// Hash password
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error %s", err)
		return 0, err
	}
	hashedPasswordString := string(hashedPassword)

	// Add user to DB
	_, err = s.executeOnDB("UPDATE User SET Firstname=?, Lastname=?, Username=?, Password=? WHERER Id = ?", firstname, lastname, username, hashedPasswordString, userId)
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
func (s *SchabloneServer) verifyAPIToken(apiToken string) (int64, error) {
	rows, err := s.queryDB("SELECT BelongsToUser FROM ActiveAPIKey WHERE Content=?", apiToken)
	if err != nil {
		log.Printf("Error %s", err)
		return 0, err
	}

	var userId int64
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			log.Printf("Error %s", err)
			return 0, err
		}
	}
	return userId, nil
}

func (s *SchabloneServer) userHasAccesssTo(groupId []int, sqlStatement string, args ...interface{}) (bool, error) {
	rows, err := s.queryDB(sqlStatement, args...)
	if err != nil {
		log.Printf("Error %s", err)
		log.Printf("Error Hello")
		return false, err
	}
	groups := []int{}
	for rows.Next() {
		var groupId int
		err := rows.Scan(&groupId)
		if err != nil {
			log.Printf("Error %s", err)
			return false, err
		}
		groups = append(groups, groupId)
	}

	for _, t := range groups {
		c, err := s.getAllGroupChildern(t)
		if err != nil {
			return false, err
		}
		for _, group := range c {
			for _, gId := range groupId {
				if group == gId {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

func (s *SchabloneServer) userHasUserModifyAccessTo(groupId []int, userId int) (bool, error) {
	return s.userHasAccesssTo(groupId, "SELECT TemplateGroup FROM User JOIN User_TemplateGroup ON User.Id = User_TemplateGroup.BelongsToUser WHERE Id=? AND UserHasUserModifyAccess=?", userId, true)
	// return true, nil
}
func (s *SchabloneServer) userHasWriteAccessTo(groupId []int, userId int) (bool, error) {
	return s.userHasAccesssTo(groupId, "SELECT TemplateGroup FROM User JOIN User_TemplateGroup ON User.Id = User_TemplateGroup.BelongsToUser WHERE Id=? AND UserHasWriteAccess=?", userId, true)
	// return true, nil
}

func (s *SchabloneServer) userHasReadAccessTo(groupId []int, userId int) (bool, error) {
	return s.userHasAccesssTo(groupId, "SELECT TemplateGroup FROM User JOIN User_TemplateGroup ON User.Id = User_TemplateGroup.BelongsToUser WHERE Id=?", userId)
	// return true, nil
}

func (s *SchabloneServer) getTemplateGroups(templateId int) ([]int, error) {
	groups := []int{}
	// Get TemplateGroup
	rows, err := s.queryDB("SELECT TemplateGroup FROM Template JOIN Template_TemplateGroup ON Template.Id = Template_TemplateGroup.BelongsToTemplate WHERE Id=?", templateId)
	if err != nil {
		log.Printf("Error %s", err)
		return []int{}, nil
	}
	for rows.Next() {
		var groupId int
		err := rows.Scan(&groupId)
		if err != nil {
			log.Printf("Error %s", err)
			return []int{}, nil
		}
		groups = append(groups, groupId)
	}
	return groups, nil
}

func (s *SchabloneServer) getMacroGroups(templateId int) ([]int, error) {
	groups := []int{}
	// Get TemplateGroup
	rows, err := s.queryDB("SELECT TemplateGroup FROM Macro JOIN Macro_TemplateGroup ON Macro.Id = Template_TemplateGroup.BelongsToTemplate WHERE Id=?", templateId)
	if err != nil {
		log.Printf("Error %s", err)
		return []int{}, err
	}
	for rows.Next() {
		var groupId int
		err := rows.Scan(&groupId)
		if err != nil {
			log.Printf("Error %s", err)
			return []int{}, err
		}
		groups = append(groups, groupId)
	}
	return groups, nil
}

func (s *SchabloneServer) getAllGroupChildern(template int) ([]int, error) {
	rows, err := s.queryDB("SELECT Id FROM TemplateGroup WHERE ParentTemplateGroup=?", template)
	if err != nil {
		log.Printf("Error %s", err)
		return []int{}, err
	}
	var templateIds []int
	for rows.Next() {
		var templateId int
		err := rows.Scan(&templateId)
		templateIds = append(templateIds, templateId)
		if err != nil {
			log.Printf("Error %s", err)
			return []int{}, err
		}
	}
	out := []int{template}
	for _, t := range templateIds {
		children, err := s.getAllGroupChildern(t)
		if err != nil {
			log.Printf("Error %s", err)
			return []int{}, err
		}
		out = append(out, children[:]...)
	}
	return out, nil
}
