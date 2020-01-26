package controllers

import (
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
	app.Handle("GET", "/", "GetMemberList")
	app.Handle("GET", "/{id: string}", "GetMember")
	app.Router().Use(middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("POST", "/", "CreateMember")
	app.Handle("PUT", "/{id: string}", "UpdateMember")
	app.Handle("DELETE", "/{id: string}", "DeleteMember")
}

// GetMemberList GET /members
func (c *MemberController) GetMemberList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "members.created_at")
	if err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "MemberController::GetMemberList", errParams)
	}
	nameZh := c.Context.URLParamDefault("nameZh", "")
	nameEn := c.Context.URLParamDefault("nameEn", "")
	listParameters := utils.GetMemberListParameters{
		GetListParameters: listParams,
		NameZh:            nameZh,
		NameEn:            nameEn,
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
	var form memberCreateForm
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
	MemberID := c.Context.Params().Get("id")

	// PSQL - Looking for specified word viea the ID.
	member, err := c.MemberService.GetMember(MemberID)
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
	memberID := c.Context.Params().Get("id")
	var form memberUpdateForm

	// Read JSON from request and validate request
	if err := utils.ReadValidateForm(c.Context, &form); err != nil {
		utils.SetError(c.Context, iris.StatusBadRequest, "MemberController::UpdateMember", err)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the give ID
	if err := c.MemberService.UpdateMember(memberID, updateData); err != nil {
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
	memberID := c.Context.Params().Get("id")

	// PSQL = Soft delete of the given ID
	if err := c.MemberService.DeleteMember(memberID); err != nil {
		utils.SetError(c.Context, iris.StatusUnprocessableEntity, "MemberService::DeleteMember", errSQL)
		return
	}

	// Returns with 204 No Content Status.
	c.Context.StatusCode(iris.StatusNoContent)
}
