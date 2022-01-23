package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SchabloneServer struct {
	db *sql.DB
}

// Make sure we conform to ServerInterface
var _ ServerInterface = (*SchabloneServer)(nil)

// Create new server
func NewSchabloneServer(mariadbUsername string, mariadbPassword string, mariadbHost string) *SchabloneServer {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", mariadbUsername+":"+mariadbPassword+"@tcp("+mariadbHost+":3306)/schablone")
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	s := &SchabloneServer{
		db: db,
	}
	s.initializeDB()
	return s
}

// No permissions
func (s *SchabloneServer) noPermissions(w http.ResponseWriter) {
	w.WriteHeader(405)
}

// VS-Code: Generate Interface : s *SchabloneServer ServerInterface

// (POST /group/add_macro)
func (s *SchabloneServer) PostGroupAddMacro(w http.ResponseWriter, r *http.Request, params PostGroupAddMacroParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /group/add_template)
func (s *SchabloneServer) PostGroupAddTemplate(w http.ResponseWriter, r *http.Request, params PostGroupAddTemplateParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /group/add_user)
func (s *SchabloneServer) PostGroupAddUser(w http.ResponseWriter, r *http.Request, params PostGroupAddUserParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /group/create)
func (s *SchabloneServer) PostGroupCreate(w http.ResponseWriter, r *http.Request, params PostGroupCreateParams) {
	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if !access {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Execute request
	results, err := s.executeOnDB("INSERT INTO TemplateGroup(Name,ParentTemplateGroup) values (?,?)", params.Name, params.ParentGroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

// (GET /group/get/{groupId})
func (s *SchabloneServer) GetGroupGetGroupId(w http.ResponseWriter, r *http.Request, groupId int) {
	panic("not implemented") // TODO: Implement
}

// (GET /group/list)
func (s *SchabloneServer) GetGroupList(w http.ResponseWriter, r *http.Request, params GetGroupListParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /group/remove_macro)
func (s *SchabloneServer) PostGroupRemoveMacro(w http.ResponseWriter, r *http.Request, params PostGroupRemoveMacroParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /group/change_parent_group)
func (s *SchabloneServer) PostGroupChangeParentGroup(w http.ResponseWriter, r *http.Request, params PostGroupChangeParentGroupParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /group/remove_template)
func (s *SchabloneServer) PostGroupRemoveTemplate(w http.ResponseWriter, r *http.Request, params PostGroupRemoveTemplateParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /group/remove_user)
func (s *SchabloneServer) PostGroupRemoveUser(w http.ResponseWriter, r *http.Request, params PostGroupRemoveUserParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /macro/create)
func (s *SchabloneServer) PostMacroCreate(w http.ResponseWriter, r *http.Request, params PostMacroCreateParams) {

	// Execute request
	templateId, err := s.executeOnDB("INSERT INTO Macro(Name,Content) values (?,?)", params.Name, params.Content)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	_, err = s.executeOnDB("INSERT INTO Macro_TemplateGroup(BelongsToMacro,TemplateGroup) values (?,?)", templateId, params.InitialGroup)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(templateId)
}

// (POST /macro/edit/checkin)
func (s *SchabloneServer) PostMacroEditCheckin(w http.ResponseWriter, r *http.Request, params PostMacroEditCheckinParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /macro/edit/checkout)
func (s *SchabloneServer) PostMacroEditCheckout(w http.ResponseWriter, r *http.Request, params PostMacroEditCheckoutParams) {
	panic("not implemented") // TODO: Implement
}

// (GET /macro/get/{macroId})
func (s *SchabloneServer) GetMacroGetMacroId(w http.ResponseWriter, r *http.Request, macroId int) {
	panic("not implemented") // TODO: Implement
}

// (GET /macro/list)
func (s *SchabloneServer) GetMacroList(w http.ResponseWriter, r *http.Request, params GetMacroListParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /template/create)
func (s *SchabloneServer) PostTemplateCreate(w http.ResponseWriter, r *http.Request, params PostTemplateCreateParams) {

	// Execute request
	templateId, err := s.executeOnDB("INSERT INTO Template(Name,Subject,Content) values (?,?,?)", params.Name, params.Subject, params.Content)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	_, err = s.executeOnDB("INSERT INTO Template_TemplateGroup(BelongsToTemplate,TemplateGroup) values (?,?)", templateId, params.InitialGroup)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(templateId)
}

// (POST /template/edit/checkin)
func (s *SchabloneServer) PostTemplateEditCheckin(w http.ResponseWriter, r *http.Request, params PostTemplateEditCheckinParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /template/edit/checkout)
func (s *SchabloneServer) PostTemplateEditCheckout(w http.ResponseWriter, r *http.Request, params PostTemplateEditCheckoutParams) {
	panic("not implemented") // TODO: Implement
}

// (GET /template/get/{templateId})
func (s *SchabloneServer) GetTemplateGetTemplateId(w http.ResponseWriter, r *http.Request, templateId int) {
	panic("not implemented") // TODO: Implement
}

// (GET /template/list)
func (s *SchabloneServer) GetTemplateList(w http.ResponseWriter, r *http.Request, params GetTemplateListParams) {
	panic("not implemented") // TODO: Implement
}

// (POST /user/create)
func (s *SchabloneServer) PostUserCreate(w http.ResponseWriter, r *http.Request, params PostUserCreateParams) {
	userId, err := s.createUser(params.Username, params.Firstname, params.Lastname, params.Password)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userId)
}

// (GET /user/get/{userId})
func (s *SchabloneServer) GetUserGetUserId(w http.ResponseWriter, r *http.Request, userId int) {
	panic("not implemented") // TODO: Implement
}

// (GET /user/list)
func (s *SchabloneServer) GetUserList(w http.ResponseWriter, r *http.Request) {
	panic("not implemented") // TODO: Implement
}

// (GET /user/login)
func (s *SchabloneServer) GetUserLogin(w http.ResponseWriter, r *http.Request, params GetUserLoginParams) {
	apiKey, err := s.verifyUser(params.Username, params.Password)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiKey)
}

// (POST /user/modify/{userId})
func (s *SchabloneServer) PostUserModifyUserId(w http.ResponseWriter, r *http.Request, userId int, params PostUserModifyUserIdParams) {
	panic("not implemented") // TODO: Implement
}
