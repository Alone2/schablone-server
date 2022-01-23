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

// func setHeaders(w http.ResponseWriter) {
// 	w.Header().Add("Access-Control-Allow-Origin", "*")
// 	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 	w.Header().Add("Access-Control-Allow-Headers", "origin, content-type, accept, x-requested-with")
// 	w.Header().Add("Access-Control-Max-Age", "3600")
// }

// VS-Code: Generate Interface : s *SchabloneServer ServerInterface

// (POST /group/add_macro)
func (s *SchabloneServer) PostGroupAddMacro(w http.ResponseWriter, r *http.Request, params PostGroupAddMacroParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Execute request
	_, err = s.executeOnDB("INSERT INTO Macro_TemplateGroup(BelongsToMacro,TemplateGroup) values (?,?)", params.MacroId, params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// (POST /group/add_template)
func (s *SchabloneServer) PostGroupAddTemplate(w http.ResponseWriter, r *http.Request, params PostGroupAddTemplateParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Execute request
	_, err = s.executeOnDB("INSERT INTO Template_TemplateGroup(BelongsToTemplate,TemplateGroup) values (?,?)", params.TemplateId, params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// (POST /group/add_user)
func (s *SchabloneServer) PostGroupAddUser(w http.ResponseWriter, r *http.Request, params PostGroupAddUserParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasUserModifyAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Execute request
	_, err = s.executeOnDB("INSERT INTO User_TemplateGroup(BelongsToUser, TemplateGroup, UserHasWriteAccess, UserHasUserModifyAccess) values (?,?,?,?)", params.UserId, params.GroupId, 1, 0)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// (POST /group/create)
func (s *SchabloneServer) PostGroupCreate(w http.ResponseWriter, r *http.Request, params PostGroupCreateParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.ParentGroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
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
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasReadAccessTo([]int{groupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	var group Group
	// Get Group
	rows, err := s.queryDB("SELECT Id, Name, ParentTemplateGroup FROM TemplateGroup WHERE Id=?", groupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		err := rows.Scan(&group.Id, &group.Name, group.ParentId)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(group)
}

// (GET /group/list)
func (s *SchabloneServer) GetGroupList(w http.ResponseWriter, r *http.Request, params GetGroupListParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	var groupList []Group
	// Get Groups
	rows, err := s.queryDB("SELECT Id, Name, ParentTemplateGroup FROM TemplateGroup")
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.Id, &group.Name, group.ParentId)
		groupList = append(groupList, group)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groupList)
}

// (POST /group/remove_macro)
func (s *SchabloneServer) PostGroupRemoveMacro(w http.ResponseWriter, r *http.Request, params PostGroupRemoveMacroParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Execute request
	_, err = s.executeOnDB("DELETE FROM Macro_TemplateGroup WHERE BelongsToMacro=? AND TemplateGroup=?)", params.MacroId, params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// (POST /group/change_parent_group)
func (s *SchabloneServer) PostGroupChangeParentGroup(w http.ResponseWriter, r *http.Request, params PostGroupChangeParentGroupParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.ParentGroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	_, err = s.executeOnDB("UPDATE Template SET ParentTemplateGroup=? WHERE Id=?", params.ParentGroupId, params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// (POST /group/remove_template)
func (s *SchabloneServer) PostGroupRemoveTemplate(w http.ResponseWriter, r *http.Request, params PostGroupRemoveTemplateParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Execute request
	_, err = s.executeOnDB("DELETE FROM Template_TemplateGroup WHERE BelongsToTemplate=? AND TemplateGroup=?)", params.TemplateId, params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// (POST /group/remove_user)
func (s *SchabloneServer) PostGroupRemoveUser(w http.ResponseWriter, r *http.Request, params PostGroupRemoveUserParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasUserModifyAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Execute request
	_, err = s.executeOnDB("DELETE FROM User_TemplateGroup WHERE BelongsToUser=? AND TemplateGroup=?)", params.UserId, params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// (POST /macro/create)
func (s *SchabloneServer) PostMacroCreate(w http.ResponseWriter, r *http.Request, params PostMacroCreateParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.InitialGroup}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

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
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	groups, _ := s.getMacroGroups(params.MacroId)
	perms, err := s.userHasWriteAccessTo(groups, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Update checkin
	_, err = s.executeOnDB("Update Macro SET IsBeingEditedBy=?, Name=?, Content=? WHERE Id=?", nil, params.Name, params.Content, params.MacroId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// (POST /macro/edit/checkout)
func (s *SchabloneServer) PostMacroEditCheckout(w http.ResponseWriter, r *http.Request, params PostMacroEditCheckoutParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	groups, _ := s.getMacroGroups(params.MacroId)
	perms, err := s.userHasWriteAccessTo(groups, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Update checkout
	_, err = s.executeOnDB("Update Macro SET IsBeingEditedBy=? WHERE Id=?", access, params.MacroId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// (GET /macro/get/{macroId})
func (s *SchabloneServer) GetMacroGetMacroId(w http.ResponseWriter, r *http.Request, macroId int) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	groups, _ := s.getMacroGroups(macroId)
	perms, err := s.userHasReadAccessTo(groups, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	var macro Macro
	// Get Macro
	rows, err := s.queryDB("SELECT Id, Name, Content, IsBeingEditedBy FROM Macro WHERE Id=?", macroId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		err := rows.Scan(&macro.Id, &macro.Title, &macro.Content, &macro.IsBeingEditedBy)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(macro)
}

// (GET /macro/list)
func (s *SchabloneServer) GetMacroList(w http.ResponseWriter, r *http.Request, params GetMacroListParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasReadAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	var macroList []Macro
	// Get Macro
	rows, err := s.queryDB("SELECT Id, Name, Content, IsBeingEditedBy FROM Macro JOIN Macro_TemplateGroup ON Macro.Id = Macro_TemplateGroup.BelongsToMacro WHERE TemplateGroup=?", params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		var macro Macro
		err := rows.Scan(&macro.Id, &macro.Title, &macro.Content, &macro.IsBeingEditedBy)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
		macroList = append(macroList, macro)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(macroList)
}

// (POST /template/create)
func (s *SchabloneServer) PostTemplateCreate(w http.ResponseWriter, r *http.Request, params PostTemplateCreateParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{params.InitialGroup}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

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
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	groups, _ := s.getTemplateGroups(params.TemplateId)
	perms, err := s.userHasWriteAccessTo(groups, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Update checkin
	_, err = s.executeOnDB("UPDATE Template SET IsBeingEditedBy=?, Content=?, Subject=? WHERE Id=?", nil, params.Content, params.Name, params.Subject, params.TemplateId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// (POST /template/edit/checkout)
func (s *SchabloneServer) PostTemplateEditCheckout(w http.ResponseWriter, r *http.Request, params PostTemplateEditCheckoutParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	groups, _ := s.getTemplateGroups(params.TemplateId)
	perms, err := s.userHasWriteAccessTo(groups, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Update checkout
	_, err = s.executeOnDB("Update Template SET IsBeingEditedBy=? WHERE Id=?", access, params.TemplateId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// (GET /template/get/{templateId})
func (s *SchabloneServer) GetTemplateGetTemplateId(w http.ResponseWriter, r *http.Request, templateId int) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	groups, _ := s.getTemplateGroups(templateId)
	perms, err := s.userHasReadAccessTo(groups, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	var template Template
	template.AttachementIds = &[]int{}
	// Get Template
	rows, err := s.queryDB("SELECT Id, Name, Content, Subject, IsBeingEditedBy FROM Template WHERE Id=?", templateId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		err := rows.Scan(&template.Id, &template.Title, &template.Content, &template.Subject, &template.IsBeingEditedBy)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(template)
}

// (GET /template/list)
func (s *SchabloneServer) GetTemplateList(w http.ResponseWriter, r *http.Request, params GetTemplateListParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasReadAccessTo([]int{params.GroupId}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	var templateList []Template
	// Get Macro
	rows, err := s.queryDB("SELECT Id, Name, Content, Subject, IsBeingEditedBy FROM Template JOIN Template_TemplateGroup ON Template.Id = Template_TemplateGroup.BelongsToTemplate WHERE TemplateGroup=?", params.GroupId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		var template Template
		template.AttachementIds = &[]int{}
		err := rows.Scan(&template.Id, &template.Title, &template.Content, &template.IsBeingEditedBy)
		templateList = append(templateList, template)

		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(templateList)
}

// (POST /user/create)
func (s *SchabloneServer) PostUserCreate(w http.ResponseWriter, r *http.Request, params PostUserCreateParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	// Permissions
	perms, err := s.userHasWriteAccessTo([]int{1}, int(access))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}
	if !perms {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

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
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	var user User
	// Get User
	rows, err := s.queryDB("SELECT Id, Firstname, Lastname, Username FROM User WHERE Id=?", userId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
	}

	user.GroupIds = &[]int{}
	// Get Groups
	rows, err = s.queryDB("SELECT TemplateGroup FROM User_TemplateGroup WHERE BelongsToUser=?", userId)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		var templateGroupId int
		err := rows.Scan(&templateGroupId)
		combination := append(*user.GroupIds, templateGroupId)
		user.GroupIds = &combination
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}

// (GET /user/list)
func (s *SchabloneServer) GetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	var userList []User
	// Get User
	rows, err := s.queryDB("SELECT Id, Firstname, Lastname, Username FROM User")
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	for rows.Next() {
		var user User
		user.GroupIds = &[]int{}
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
		// Get Groups
		rows, err = s.queryDB("SELECT TemplateGroup FROM User_TemplateGroup WHERE BelongsToUser=?", user.Id)
		if err != nil {
			log.Printf("Error %s", err)
			w.WriteHeader(400)
			return
		}
		for rows.Next() {
			var templateGroupId int
			err := rows.Scan(&templateGroupId)
			combination := append(*user.GroupIds, templateGroupId)
			user.GroupIds = &combination
			if err != nil {
				log.Printf("Error %s", err)
				w.WriteHeader(400)
				return
			}
		}
		userList = append(userList, user)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(userList)
}

// (GET /user/login)
func (s *SchabloneServer) GetUserLogin(w http.ResponseWriter, r *http.Request, params GetUserLoginParams) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
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
	w.Header().Add("Access-Control-Allow-Origin", "*")

	// Check permission
	api_key := r.Header.Get("X-API-Key")
	access, err := s.verifyAPIToken(api_key)
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	// API token Invalid
	if access <= 0 {
		log.Printf("API Token invalid %s", api_key)
		w.WriteHeader(405)
		return
	}

	if !(access == 0 || access == int64(userId)) {
		log.Printf("Error %s", err)
		w.WriteHeader(405)
		return
	}

	// Modify User
	_, err = s.editUser(params.Username, params.Firstname, params.Lastname, params.Password, int64(userId))
	if err != nil {
		log.Printf("Error %s", err)
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
