package controllers

import (
	"ccsl/middlewares"
	"ccsl/services"
	"ccsl/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// MemberController is for member CURD
type MemberController struct {
	Context       iris.Context
	MemberService services.MemberInterface
}

func (c *MemberController) BeforeActivation(app mvc.BeforeActivation) {
	app.Handle("GET", "/", "GetMemberList")
	app.Handle("GET", "/{id: string}", "GetMember")
	app.Handle("POST", "/", "CreateMember", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("PUT", "/{id: string}", "UpdateMember", middlewares.CheckJWTToken, middlewares.CheckAdmin)
	app.Handle("DELETE", "/{id string", "DeleteMember", middlewares.CheckJWTToken, middlewares.CheckAdmin)
}

func (c *MemberController) GetMemberList() {
	defer c.Context.Next()
	listParams, err := utils.GetListParamsFromContext(c.Context, "created_at")
	if err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "order only accepts 'asc' or 'desc'", err)
	}
	nameZh := c.Context.URLParamDefault("name_zh", "")
	nameEn := c.Context.URLParamDefault("name_en", "")
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
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "MemberService::CreateMember", err)
		return
	}
	//    PSQL - Create  member in database.
	member := form.ConvertToModel()
	if err := c.MemberService.CreateMember(member); err != nil {
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "MemberService::CreateMember", err)
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
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "MemberServcie:GetMember", err)
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
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "MemberController::ParamsError", err)
		return
	}
	updateData := utils.MakeUpdateData(form)

	// PSQL - Update of the give ID
	if err := c.MemberService.UpdateMember(memberID, updateData); err != nil {
		utils.SetResponseError(c.Context, iris.StatusBadRequest, "MemberService::UpdateMember", err)
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
		utils.SetResponseError(c.Context, iris.StatusUnprocessableEntity, "MemberService::DeleteMember", err)
		return
	}

	// Returns with 204 No Content Status.
	c.Context.StatusCode(iris.StatusNoContent)
}
