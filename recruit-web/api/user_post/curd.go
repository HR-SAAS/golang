package user_post

import (
	"context"
	"github.com/gin-gonic/gin"
	"hr-saas-go/recruit-web/api/post"
	"hr-saas-go/recruit-web/global"
	"hr-saas-go/recruit-web/proto"
	"hr-saas-go/recruit-web/request"
	"hr-saas-go/recruit-web/utils"
	"net/http"
	"strconv"
)

// List 投递列表,投递人信息,企业信息,简历信息
func List(field string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var sort map[string]string
		search := make(map[string]string)
		if field == "user" {
			userId := ctx.GetInt64("userId")
			search["user_id"] = strconv.FormatInt(userId, 10)
			sort = map[string]string{
				"updated_at": "desc",
				"status":     "desc",
			}
		}
		if field == "company" {
			companyId := ctx.Query("company_id")
			search["company_id"] = companyId
			if status := ctx.Query("status"); status != "" {
				search["status"] = status
			} else {
				search["status"] = "1"
			}
			sort = map[string]string{
				"created_at": "asc",
				"status":     "desc",
			}
		}
		if field == "post" {
			companyId := ctx.Query("post_id")
			search["post_id"] = companyId
		}
		page, limit := utils.GetPage(ctx)
		var list, err = global.UserPostServCon.GetUserPostList(ctx, &proto.GetUserPostListRequest{
			Page:   page,
			Limit:  limit,
			Sort:   sort,
			Search: search,
		})
		if err != nil {
			utils.HandleGrpcError(err, ctx)
			return
		}
		res := make([]interface{}, 0, limit)

		// 获取其他信息
		userIds := make([]int64, 0, limit)
		postIds := make([]int64, 0, limit)
		for _, v := range list.Data {
			postIds = append(postIds, v.PostId)
			userIds = append(userIds, v.UserId)
		}
		users, err := global.UserServCon.GetUserListByIds(ctx, &proto.GetUserListByIdsRequest{Ids: userIds})
		if err != nil {
			utils.HandleGrpcError(err, ctx)
			return
		}
		posts, err := global.PostServCon.GetPostListByIds(ctx, &proto.GetPostListByIdsRequest{Ids: postIds})
		if err != nil {
			utils.HandleGrpcError(err, ctx)
			return
		}
		dataTemp, err := post.GetListOtherData(ctx, limit, posts)
		if err != nil {
			return
		}
		data := make(map[int64]interface{})
		for _, v := range dataTemp {
			temp := v.(map[string]interface{})
			data[temp["id"].(int64)] = v
		}
		userMap := make(map[int64]interface{})
		for _, v := range users.Data {
			userMap[v.Id] = v
		}
		for _, v := range list.Data {
			res = append(res, map[string]interface{}{
				"id":          v.Id,
				"post_id":     v.PostId,
				"post":        data[v.PostId],
				"user_id":     v.UserId,
				"resume_id":   v.ResumeId,
				"resume_type": v.ResumeType,
				"resume_name": v.ResumeName,
				"user":        userMap[v.UserId],
				"resume":      v.Resume,
				"review_id":   v.ReviewId,
				"status":      v.Status,
				"company_id":  v.CompanyId,
				"remark":      v.Remark,
			})
		}

		ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
			"data":  res,
			"total": list.Total,
		}))
	}

}

// Show 查看投递详情
func Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id不正确"))
		return
	}
	// police

	data, err := global.UserPostServCon.GetUserPostDetail(ctx, &proto.GetUserPostDetailRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	// TODO 获取其他信息 联表?

	ctx.JSON(http.StatusOK, utils.SuccessJson(map[string]interface{}{
		"id": data.Id,
		// 信息
		"created_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
		"updated_at": data.CreatedAt.AsTime().Format("2006-01-02 15:01:05"),
	}))
	return
}

// Create 投递简历
func Create(ctx *gin.Context) {
	req, err := request.UserPostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	postDetail, err := global.PostServCon.GetPostDetail(ctx, &proto.GetPostDetailRequest{Id: req.PostId})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	if postDetail.Status != 1 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("岗位未开放或者已招满"))
		return
	}

	res, err := global.UserPostServCon.CreateUserPost(
		context.Background(), &proto.CreateUserPostRequest{
			PostId:     req.PostId,
			UserId:     userId,
			ResumeId:   req.ResumeId,
			ResumeType: req.ResumeType,
			ResumeName: req.ResumeName,
			Resume:     req.Resume,
			CompanyId:  postDetail.CompanyId,
			Status:     1,
		})

	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(res.Id))
	return
}

func CheckCanPost(ctx *gin.Context) {
	// 判断是否能够提交简历
	search := make(map[string]string)
	companyId, _ := strconv.Atoi(ctx.Query("company_id"))
	postId, _ := strconv.Atoi(ctx.Query("post_id"))

	search["user_id"] = strconv.FormatInt(ctx.GetInt64("userId"), 10)

	res, err := global.RecruitCounterServiceServCon.CountUserPost(ctx, &proto.CountUserPostRequest{
		CompanyId: int64(companyId),
		PostId:    int64(postId),
		Search:    nil,
	})

	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessJson(res.Count))
}

// Update 更新状态
func Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id错误"))
		return
	}
	req, err := request.UserPostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.UserPostServCon.UpdateUserPost(ctx, &proto.UpdateUserPostRequest{
		Id:        int64(id),
		ReviewId:  userId,
		Status:    req.Status,
		Remark:    req.Remark,
		CompanyId: -1,
		PostId:    -1,
		ResumeId:  -1,
		UserId:    -1,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

// BatchHandle 批量处理
func BatchHandle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusOK, utils.ErrorJson("id错误"))
		return
	}
	req, err := request.UserPostSaveRequestGet(ctx)
	if err != nil {
		return
	}
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("系统错误"))
		return
	}
	_, err = global.UserPostServCon.BatchUpdateUserPost(ctx, &proto.BatchUpdateUserPostRequest{
		Ids:      []int64{1},
		ReviewId: userId,
		Status:   req.Status,
		Remark:   req.Remark,
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(nil))
	return
}

// Delete 删除投递记录
func Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorJson("id不正确"))
		return
	}

	data, err := global.UserPostServCon.DeleteUserPost(ctx, &proto.DeleteUserPostRequest{
		Id: int64(id),
	})
	if err != nil {
		utils.HandleGrpcError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessJson(data))
	return

}
