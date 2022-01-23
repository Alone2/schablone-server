// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Group defines model for Group.
type Group struct {
	Id       *int    `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	ParentId *int    `json:"parentId,omitempty"`
}

// Macro defines model for Macro.
type Macro struct {
	Content         *string `json:"content,omitempty"`
	Id              *int    `json:"id,omitempty"`
	IsBeingEditedBy *int    `json:"isBeingEditedBy,omitempty"`
	Title           *string `json:"title,omitempty"`
}

// Template defines model for Template.
type Template struct {
	AttachementIds  *[]int  `json:"attachementIds,omitempty"`
	Content         *string `json:"content,omitempty"`
	Id              *int    `json:"id,omitempty"`
	IsBeingEditedBy *int    `json:"isBeingEditedBy,omitempty"`
	Subject         *string `json:"subject,omitempty"`
	Title           *string `json:"title,omitempty"`
}

// User defines model for User.
type User struct {
	Firstname *string `json:"firstname,omitempty"`
	GroupIds  *[]int  `json:"groupIds,omitempty"`
	Id        *int    `json:"id,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
	Username  *string `json:"username,omitempty"`
}

// PostGroupAddMacroParams defines parameters for PostGroupAddMacro.
type PostGroupAddMacroParams struct {
	// Macro ID
	MacroId int `json:"macroId"`

	// Group ID
	GroupId int `json:"groupId"`
}

// PostGroupAddTemplateParams defines parameters for PostGroupAddTemplate.
type PostGroupAddTemplateParams struct {
	// Template ID
	TemplateId int `json:"templateId"`

	// Group ID
	GroupId int `json:"groupId"`
}

// PostGroupAddUserParams defines parameters for PostGroupAddUser.
type PostGroupAddUserParams struct {
	// User ID
	UserId int `json:"userId"`

	// Group ID
	GroupId int `json:"groupId"`
}

// PostGroupChangeParentGroupParams defines parameters for PostGroupChangeParentGroup.
type PostGroupChangeParentGroupParams struct {
	// Parent Group ID
	ParentGroupId int `json:"parentGroupId"`

	// Group ID
	GroupId int `json:"groupId"`
}

// PostGroupCreateParams defines parameters for PostGroupCreate.
type PostGroupCreateParams struct {
	// Group title
	Name string `json:"name"`

	// Parent Group ID
	ParentGroupId int `json:"parentGroupId"`
}

// GetGroupListParams defines parameters for GetGroupList.
type GetGroupListParams struct {
	// The group ID
	GroupId int `json:"groupId"`
}

// PostGroupRemoveMacroParams defines parameters for PostGroupRemoveMacro.
type PostGroupRemoveMacroParams struct {
	// Macro ID
	MacroId int `json:"macroId"`

	// Group ID
	GroupId int `json:"groupId"`
}

// PostGroupRemoveTemplateParams defines parameters for PostGroupRemoveTemplate.
type PostGroupRemoveTemplateParams struct {
	// Template ID
	TemplateId int `json:"templateId"`

	// Group ID
	GroupId int `json:"groupId"`
}

// PostGroupRemoveUserParams defines parameters for PostGroupRemoveUser.
type PostGroupRemoveUserParams struct {
	// User ID
	UserId int `json:"userId"`

	// Group ID
	GroupId int `json:"groupId"`
}

// PostMacroCreateParams defines parameters for PostMacroCreate.
type PostMacroCreateParams struct {
	// Macro title
	Name string `json:"name"`

	// Macro content
	Content string `json:"content"`

	// Group ID of the templates parent group
	InitialGroup int `json:"initialGroup"`
}

// PostMacroEditCheckinParams defines parameters for PostMacroEditCheckin.
type PostMacroEditCheckinParams struct {
	// Macro ID
	MacroId int `json:"macroId"`

	// Macro title
	Name string `json:"name"`

	// Macro content
	Content string `json:"content"`
}

// PostMacroEditCheckoutParams defines parameters for PostMacroEditCheckout.
type PostMacroEditCheckoutParams struct {
	// Macro ID
	MacroId int `json:"macroId"`
}

// GetMacroListParams defines parameters for GetMacroList.
type GetMacroListParams struct {
	// Group ID
	GroupId int `json:"groupId"`
}

// PostTemplateCreateParams defines parameters for PostTemplateCreate.
type PostTemplateCreateParams struct {
	// Template title
	Name string `json:"name"`

	// Template subject
	Subject string `json:"subject"`

	// Template content
	Content string `json:"content"`

	// Group ID of the templates parent group
	InitialGroup int `json:"initialGroup"`
}

// PostTemplateEditCheckinParams defines parameters for PostTemplateEditCheckin.
type PostTemplateEditCheckinParams struct {
	// Template ID
	TemplateId int `json:"templateId"`

	// Template title
	Name string `json:"name"`

	// Template subject
	Subject string `json:"subject"`

	// Template content
	Content string `json:"content"`
}

// PostTemplateEditCheckoutParams defines parameters for PostTemplateEditCheckout.
type PostTemplateEditCheckoutParams struct {
	// Template ID
	TemplateId int `json:"templateId"`
}

// GetTemplateListParams defines parameters for GetTemplateList.
type GetTemplateListParams struct {
	// Group ID
	GroupId int `json:"groupId"`
}

// PostUserCreateParams defines parameters for PostUserCreate.
type PostUserCreateParams struct {
	// The user name
	Username string `json:"username"`

	// The first name of the user
	Firstname string `json:"firstname"`

	// The last name of the user
	Lastname string `json:"lastname"`

	// The password
	Password string `json:"password"`
}

// GetUserLoginParams defines parameters for GetUserLogin.
type GetUserLoginParams struct {
	// The user name
	Username string `json:"username"`

	// The user password
	Password string `json:"password"`
}

// PostUserModifyUserIdParams defines parameters for PostUserModifyUserId.
type PostUserModifyUserIdParams struct {
	// The user name
	Username string `json:"username"`

	// The first name of the user
	Firstname string `json:"firstname"`

	// The last name of the user
	Lastname string `json:"lastname"`

	// The password
	Password string `json:"password"`
}

