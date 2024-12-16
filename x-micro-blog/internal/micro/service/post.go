package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	pb "x-micro-blog/internal/micro/proto/post"

	"github.com/richxan/xcommon/xlog"

	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-micro.dev/v4"
)

const postServiceName = "post-service"

// StartPostService 启动文章服务
func StartPostService(srv micro.Service, logger *xlog.Logger) error {
	// 注册处理器
	if err := pb.RegisterPostServiceHandler(srv.Server(), NewPostHandler(logger)); err != nil {
		return fmt.Errorf("register handler error: %v", err)
	}

	// 启动 metrics 服务器
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9090", nil); err != nil {
			logger.Error().Err(err).Msg("metrics server error")
		}
	}()

	return nil
}

// PostHandler 文章服务处理器
type PostHandler struct {
	logger *xlog.Logger
}

// NewPostHandler 创建文章服务处理器
func NewPostHandler(logger *xlog.Logger) *PostHandler {
	return &PostHandler{logger: logger}
}

// CreatePost 创建文章
func (h *PostHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest, rsp *pb.CreatePostResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreatePost")
	defer span.Finish()

	h.logger.Info().
		Str("title", req.Title).
		Msg("Creating new post")

	// TODO: 实现创建文章逻辑

	return nil
}

// GetPost 获取文章
func (h *PostHandler) GetPost(ctx context.Context, req *pb.GetPostRequest, rsp *pb.GetPostResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetPost")
	defer span.Finish()

	postID, err := strconv.ParseUint(req.PostId, 10, 64)
	if err != nil {
		h.logger.Error().Err(err).Str("post_id", req.PostId).Msg("Invalid post ID format")
		return fmt.Errorf("invalid post ID format: %v", err)
	}

	h.logger.Info().
		Uint64("post_id", postID).
		Msg("Getting post")

	// TODO: 实现获取文章逻辑

	return nil
}

// ListPosts 获取文章列表
func (h *PostHandler) ListPosts(ctx context.Context, req *pb.ListPostsRequest, rsp *pb.ListPostsResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ListPosts")
	defer span.Finish()

	h.logger.Info().
		Int32("page", req.Page).
		Int32("page_size", req.PageSize).
		Msg("Listing posts")

	// TODO: 实现获取文章列表逻辑

	return nil
}

// UpdatePost 更新文章
func (h *PostHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest, rsp *pb.UpdatePostResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdatePost")
	defer span.Finish()

	postID, err := strconv.ParseUint(req.PostId, 10, 64)
	if err != nil {
		h.logger.Error().Err(err).Str("post_id", req.PostId).Msg("Invalid post ID format")
		return fmt.Errorf("invalid post ID format: %v", err)
	}

	h.logger.Info().
		Uint64("post_id", postID).
		Str("title", req.Title).
		Msg("Updating post")

	// TODO: 实现更新文章逻辑

	return nil
}

// DeletePost 删除文章
func (h *PostHandler) DeletePost(ctx context.Context, req *pb.DeletePostRequest, rsp *pb.DeletePostResponse) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DeletePost")
	defer span.Finish()

	postID, err := strconv.ParseUint(req.PostId, 10, 64)
	if err != nil {
		h.logger.Error().Err(err).Str("post_id", req.PostId).Msg("Invalid post ID format")
		return fmt.Errorf("invalid post ID format: %v", err)
	}

	h.logger.Info().
		Uint64("post_id", postID).
		Msg("Deleting post")

	// TODO: 实现���除文章逻辑

	return nil
}
