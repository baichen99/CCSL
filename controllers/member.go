package controllers

import (
	"ccsl/configs"
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// MemberController is for member CRUD
type MemberController struct {
	Context       iris.Context
	MemberService services.MemberInterface
}

// BeforeActivation will register routes for controllers
func (c *MemberController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle(iris.MethodGet, "/", "GetMemberList")
	app.Handle(iris.MethodGet, "/{id: string}", "GetMember")
	app.Router().Use(middlewares.CheckToken, middlewares.CheckUserRole([]string{configs.RoleAdminUser}))
	app.Handle(iris.MethodPost, "/", "CreateMember")
	app.Handle(iris.MethodPut, "/{id: string}", "UpdateMember")
	app.Handle(iris.MethodDelete, "/{id: string}", "DeleteMember")
}

// GetMemberList GET /members
func (c *MemberController) GetMemberList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "members.created_at")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "MemberController::GetMemberList", errParams)
	}
	listParameters := utils.GetMemberListParameters{
		GetListParameters: listParams,
		NameZh:            c.Context.URLParamDefault("nameZh", ""),
		NameEn:            c.Context.URLParamDefault("nameEn", ""),
	}
	members, count, err := c.MemberService.GetMemberList(listParameters)
	c.Context.JSON(iris.Map{
		message: success,
		data:    members,
		page:    listParams.Page,
		limit:   listParams.Limit,
		total:   count,
	})
}

// CreateMember POST /members
func (c *MemberController) CreateMember() {
	defer c.Context.Next()
	var form MemberCreateForm
	//   Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "MemberService::CreateMember", errParams)
		return
	}
	//    PSQL - Create  member in database.
	member := form.ConvertToModel()
	if err := c.MemberService.CreateMember(member); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "MemberService::CreateMember", errSQL)
		return
	}
	//    return 201 Created
	c.Context.StatusCode(iris.StatusCreated)
	c.Context.JSON(
		iris.Map{
			message: success,
		})
}

// GetMember GET /members/{id: string}
func (c *MemberController) GetMember() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL - Looking for specified word viea the ID.
	member, err := c.MemberService.GetMember(id)
	if err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "MemberServcie:GetMember", errSQL)
		return
	}
	// Returning word information in data key.
	c.Context.JSON(
		iris.Map{
			message: success,
			data:    member,
		})
}

// UpdateMember PUT /members/{id: string}
func (c *MemberController) UpdateMember() {
	defer c.Context.Next()

	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")
	var form MemberUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "MemberController::UpdateMember", err)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the give ID
	if err := c.MemberService.UpdateMember(id, updateData); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "MemberService::UpdateMember", errSQL)
		return
	}

	// Returns with 204 No Content Status
	c.Context.StatusCode(iris.StatusNoContent)
}

// DeleteMember DELETE /members/{id: string}
func (c *MemberController) DeleteMember() {
	defer c.Context.Next()
	// Getting ID from parameters in the URL
	id := c.Context.Params().Get("id")

	// PSQL = Soft delete of the given ID
	if err := c.MemberService.DeleteMember(id); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "MemberService::DeleteMember", errSQL)
		return
	}

	// Returns with 204 No Content Status.
	c.Context.StatusCode(iris.StatusNoContent)
}
