// Code generated by hertz generator.

package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"miniDouyin/biz/dal/pg"
	"miniDouyin/biz/model/miniDouyin/api"
)

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	fmt.Println("Feed 被调用")
	var err error
	var req api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		fmt.Println("参数绑定失败")
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FeedResponse)

	pg.DBVideoFeed(&req, resp)

	fmt.Printf("resp +v", resp)

	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserRegisterResponse)

	pg.DBUserRegister(&req, resp)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserLoginResponse)

	pg.DBUserLogin(&req, resp)

	// Debug
	fmt.Printf("resp = %+v\n", resp)

	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.UserResponse)

	pg.DBGetUserinfo(&req, resp)

	c.JSON(consts.StatusOK, resp)
}

// VideoPublishAction .
// @router /douyin/publish/action/ [POST]
func VideoPublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionRequest
	resp := new(api.PublishActionResponse)

	form, err := c.MultipartForm()

	//err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
	}

	pg.DBReceiveVideo(&req, resp, form, c)

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PublishListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionResponse
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/publish/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavoriteListResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionResponse
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentListResponse)

	c.JSON(consts.StatusOK, resp)
}

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RelationActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FollowList .
// @router /douyin/relatioin/follow/list/ [GET]
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RelationFollowListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FollowerList .
// @router /douyin/relatioin/follower/list/ [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RelationFollowerListResponse)

	c.JSON(consts.StatusOK, resp)
}

// FriendList .
// @router /douyin/relatioin/friend/list/ [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RelationFriendListResponse)

	c.JSON(consts.StatusOK, resp)
}

// ChatRec .
// @router /douyin/message/chat/ [GET]
func ChatRec(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.ChatRecordResponse
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.ChatRecordResponse)

	c.JSON(consts.StatusOK, resp)
}

// SendMsg .
// @router /douyin/message/action/ [POST]
func SendMsg(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SendMsgRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.SendMsgRequest)

	c.JSON(consts.StatusOK, resp)
}
