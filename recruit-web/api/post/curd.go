package post

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/proto"
	"hr-saas-go/recruit-web/request"
	"hr-saas-go/recruit-web/utils"
	"net/http"
	"strconv"
)

func GetListOtherData(ctx *gin.Context, limit int32, list *proto.PostListResponse) ([]interface{}, error) {
	resumes := make([]interface{}, 0, limit)
	departmentIds := make([]int64, 0, limit)
	companyIds := make([]int64, 0, limit)
	creatorIds := make([]int64, 0, limit)
	for _, v := range list.Data {
		departmentIds = append(departmentIds, v.DepartmentId)
		companyIds = append(companyIds, v.CompanyId)
		creatorIds = append(creatorIds, v.CreatorId)
	}
	// 获取数据

	departments, err := global.DepartmentServCon.GetDepartmentListByIds(ctx, &proto.GetDepartmentListByIdsRequest{Ids: departmentIds})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return nil, err
	}

	departmentMap := make(map[int64]interface{})
	for _, d := range departments.Data {
		departmentMap[d.Id] = d
	}

	companies, err := global.CompanyServCon.GetCompanyListByIds(ctx, &proto.GetCompanyListByIdsRequest{Ids: companyIds})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return nil, err
	}
	companyMap := make(map[int64]interface{})
	for _, c := range companies.Data {
		companyMap[c.Id] = c
	}

	for _, v := range list.Data {
		resumes = append(resumes, map[string]interface{}{
			"id":            v.Id,
			"name":          v.Name,
			"type":          v.Type,
			"content":       v.Content,
			"experience":    v.Experience,
			"education":     v.Education,
			"desc":          v.Desc,
			"company_id":    v.CompanyId,
			"company":       companyMap[v.CompanyId],
			"department_id": v.DepartmentId,
			"department":    departmentMap[v.DepartmentId],
			"creator_id":    v.CreatorId,
			"status":        v.Status,
			"created_at":    v.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
			"updated_at":    v.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		})
	}
	return resumes, nil
}

func List(ctx *gin.Context) {
	search := make(map[string]string)

	userId := ctx.GetInt64("userId")
	search["name"] = ctx.Query("name")
	search["user_id"] = strconv.FormatInt(userId, 10)
	search["company_id"] = ctx.Query("company_id")

	zap.S().Info("cid", search["company_id"])

	page, limit := utils.GetPage(ctx)
	list, err := global.PostServCon.GetPostList(ctx, &proto.GetPostListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
		Sort: map[string]string{
			"updated_at": "desc",
		},
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	res, err := GetListOtherData(ctx, limit, list)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"data":  res,
		"total": list.Total,
	}))
}

func Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id不正确"))
		return
	}
	// police

	data, err := global.PostServCon.GetPostDetail(ctx, &proto.GetPostDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	// 获取company
	company, _ := global.CompanyServCon.GetCompanyDetail(ctx, &proto.GetCompanyDetailRequest{Id: data.CompanyId})

	// 获取department
	department, _ := global.DepartmentServCon.GetDepartmentDetail(ctx, &proto.GetDepartmentDetailRequest{Id: data.DepartmentId})
	// 获取creator_id

	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"id":            data.Id,
		"name":          data.Name,
		"type":          data.Type,
		"content":       data.Content,
		"experience":    data.Experience,
		"education":     data.Education,
		"desc":          data.Desc,
		"company_id":    data.CompanyId,
		"company":       company,
		"department_id": data.DepartmentId,
		"department":    department,
		"creator_id":    data.CreatorId,
		"status":        data.Status,
		"created_at":    data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		"updated_at":    data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
	}))
	return
}

func Create(ctx *gin.Context) {
	req, err := request.PostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")

	startAt, err := utils.FormatTimeToStampPb(req.StartAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}

	endAt, err := utils.FormatTimeToStampPb(req.EndAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}
	address, _ := json.Marshal(req.Address)

	res, err := global.PostServCon.CreatePost(
		context.Background(), &proto.CreatePostRequest{
			CompanyId:    req.CompanyId,
			DepartmentId: req.DepartmentId,
			CreatorId:    userId,
			Type:         req.Type,
			Name:         req.Name,
			Desc:         req.Desc,
			Content:      req.Content,
			Experience:   req.Experience,
			Education:    req.Education,
			Address:      string(address),
			StartAt:      startAt,
			EndAt:        endAt,
			Status:       req.Status,
		})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res.Id))
	return
}

func Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id错误"))
		return
	}
	req, err := request.PostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	startAt, err := utils.FormatTimeToStampPb(req.StartAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}

	endAt, err := utils.FormatTimeToStampPb(req.EndAt)
	if err != nil {
		ctx.JSON(http.StatusOK, utils.ErrorJson("时间错误"))
		return
	}
	address, _ := json.Marshal(req.Address)

	_, err = global.PostServCon.UpdatePost(ctx, &proto.UpdatePostRequest{
		Id:           int64(id),
		CompanyId:    req.CompanyId,
		DepartmentId: req.DepartmentId,
		CreatorId:    userId,
		Type:         req.Type,
		Name:         req.Name,
		Desc:         req.Desc,
		Content:      req.Content,
		Experience:   req.Experience,
		Education:    req.Education,
		Address:      string(address),
		StartAt:      startAt,
		EndAt:        endAt,
		Status:       req.Status,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

func Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("id不正确"))
		return
	}
	// 修改状态为取消投递 -1
	data, err := global.PostServCon.DeletePost(ctx, &proto.DeletePostRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return

}
