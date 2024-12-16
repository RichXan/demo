package service

import (
	"context"
	"x-micro-blog/internal/access/model"
)

// UserService 用户服务接口
type UserService interface {
	// 用户基本功能
	Register(ctx context.Context, username, password, email string) (*model.User, error)
	Login(ctx context.Context, username, password string) (string, error)
	GetUser(ctx context.Context, userID int64) (*model.User, error)
	UpdateUser(ctx context.Context, userID int64, nickname, avatar, bio string) (*model.User, error)
	ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error
	ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int64, error)

	// 社交账号功能
	ListSocialAccounts(ctx context.Context, userID int64) ([]*model.SocialAccount, error)
	BindSocialAccount(ctx context.Context, userID int64, provider, code, state string) error
	UnbindSocialAccount(ctx context.Context, userID int64, provider string) error

	// 关注功能
	FollowUser(ctx context.Context, userID, targetID int64) error
	UnfollowUser(ctx context.Context, userID, targetID int64) error
	ListFollowers(ctx context.Context, userID int64, page, pageSize int) ([]*model.User, int64, error)
	ListFollowing(ctx context.Context, userID int64, page, pageSize int) ([]*model.User, int64, error)
}

// PostService 文章服务接口
type PostService interface {
	// TODO: 添加文章服务接口定义
	CreatePost(ctx context.Context, userID int64, title, content string) (*model.Post, error)
	GetPost(ctx context.Context, postID int64) (*model.Post, error)
	ListPosts(ctx context.Context, page, pageSize int) ([]*model.Post, int64, error)
	UpdatePost(ctx context.Context, postID int64, title, content string) (*model.Post, error)
	DeletePost(ctx context.Context, postID int64) error
	CreateComment(ctx context.Context, postID, userID int64, content string) (*model.Comment, error)
	ListComments(ctx context.Context, postID int64, page, pageSize int) ([]*model.Comment, int64, error)
	DeleteComment(ctx context.Context, commentID int64) error
	GetComment(ctx context.Context, commentID int64) (*model.Comment, error)
	ListPostLikes(ctx context.Context, postID int64, page, pageSize int) ([]*model.User, int64, error)
	LikePost(ctx context.Context, postID, userID int64) error
	UnlikePost(ctx context.Context, postID, userID int64) error
}

// CommentService 评论服务接口
type CommentService interface {
	// TODO: 添加评论服务接口定义
}

// LikeService 点赞服务接口
type LikeService interface {
	// TODO: 添加点赞服务接口定义
}
