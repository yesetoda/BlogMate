package controllers

import (
	"net/http"
	"strconv"

	"github.com/yesetoda/BlogMate/domain"
	"github.com/yesetoda/BlogMate/infrastructure"
	usecase "github.com/yesetoda/BlogMate/usecases"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	usecase usecase.BlogUsecase
}

func NewBlogController(uc usecase.BlogUsecase) *BlogController {
	return &BlogController{usecase: uc}
}

// HandleCreateBlog godoc
// @Summary Create a new blog post
// @Description Create a new blog post with the provided data.
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param blog body domain.Blog true "Blog post data"
// @Success 200 {object} domain.Blog "Successfully created blog post"
// @Failure 400 {object} map[string]interface{} "Invalid input data"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /blogs [post]
func (cont *BlogController) HandleCreateBlog(ctx *gin.Context) {
	var blog domain.Blog
	err := ctx.ShouldBindJSON(&blog)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	blog.AuthorId = claims.ID
	blog, err = cont.usecase.CreateBLog(blog)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blog)
	}
}

// HandleGetAllBlogs godoc
// @Summary Get all blog posts
// @Description Retrieve a paginated list of all blog posts. Defaults: pageNumber=1, pageSize=5.
// @Tags Blog
// @Produce json
// @Param pageNumber query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} interface{} "List of blog posts"
// @Failure 404 {object} map[string]interface{} "No blogs found"
// @Router /blogs [get]
func (cont *BlogController) HandleGetAllBlogs(ctx *gin.Context) {
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize
	blogs, err := cont.usecase.GetAllBlogs(x)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// HandleGetBlogById godoc
// @Summary Get a blog post by ID
// @Description Retrieve a single blog post by its unique ID.
// @Tags Blog
// @Produce json
// @Param blogId path string true "Blog post ID"
// @Success 200 {object} domain.Blog "Blog post details"
// @Failure 404 {object} map[string]interface{} "Blog not found"
// @Router /blogs/{blogId} [get]
func (cont *BlogController) HandleGetBlogById(ctx *gin.Context) {
	blogs, err := cont.usecase.GetBlogByBLogId(ctx.Param("blogId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// HandleGetPopularBlog godoc
// @Summary Get popular blog posts
// @Description Retrieve a list of blog posts that are popular.
// @Tags Blog
// @Produce json
// @Success 200 {object} interface{} "List of popular blog posts"
// @Failure 404 {object} map[string]interface{} "No popular blogs found"
// @Router /blogs/popular [get]
func (cont *BlogController) HandleGetPopularBlog(ctx *gin.Context) {
	blogs, err := cont.usecase.FindPopularBlog()
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// HandleFilterBlogs godoc
// @Summary Filter blog posts
// @Description Retrieve blog posts based on provided filter criteria with pagination.
// @Tags Blog
// @Accept json
// @Produce json
// @Param filter body domain.BlogFilterOption true "Filter criteria (include any fields you want to filter by)"
// @Param pageNumber query int false "Page number (default: 1)"
// @Param pageSize query int false "Number of items per page (default: 5)"
// @Success 200 {object} interface{} "Filtered blog posts"
// @Failure 400 {object} map[string]interface{} "Invalid filter data"
// @Failure 404 {object} map[string]interface{} "No matching blogs found"
// @Router /blogs/filter [post]
func (cont *BlogController) HandleFilterBlogs(ctx *gin.Context) {

	var blf domain.BlogFilterOption
	err := ctx.ShouldBindJSON(&blf)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize
	blf.Pagination = x

	blogs, err := cont.usecase.FilterBlogs(blf)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// HandleBlogUpdate godoc
// @Summary Update a blog post
// @Description Update an existing blog post using its ID and the new data.
// @Tags Blog
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param blogId path string true "Blog post ID"
// @Param blog body domain.Blog true "Updated blog post data"
// @Success 200 {object} domain.Blog "Updated blog post"
// @Failure 400 {object} map[string]interface{} "Invalid input data"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]interface{} "Blog not found"
// @Router /blogs/{blogId} [put]
func (cont *BlogController) HandleBlogUpdate(ctx *gin.Context) {
	var updateBlog domain.Blog
	err := ctx.ShouldBindJSON(&updateBlog)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	blog, err := cont.usecase.UpdateBLog(ctx.Param("blogId"), updateBlog)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blog)
	}
}

// HandleBlogDelete godoc
// @Summary Delete a blog post
// @Description Delete an existing blog post using its unique ID. Requires proper ownership.
// @Tags Blog
// @Produce json
// @Security BearerAuth
// @Param blogId path string true "Blog post ID"
// @Success 200 {object} map[string]string "Blog deleted"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]interface{} "Blog not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /blogs/{blogId} [delete]
func (cont *BlogController) HandleBlogDelete(ctx *gin.Context) {
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	err = cont.usecase.DeleteBLog(ctx.Param("blogId"), claims.ID)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted"})
	}
}

// HandleBlogLikeOrDislike godoc
// @Summary Interact with a blog post
// @Description Like, dislike, or view a blog post by specifying the interaction type.
// @Tags Blog Interactions
// @Produce json
// @Security BearerAuth
// @Param blogId path string true "Blog post ID"
// @Param type path string true "Interaction type (like/dislike/view)"
// @Success 200 {object} map[string]interface{} "Interaction result"
// @Failure 400 {object} map[string]string "Invalid interaction type"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]interface{} "Blog not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /blogs/{blogId}/interact/{type} [post]
func (cont *BlogController) HandleBlogLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	blogId := ctx.Param("blogId")
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	if interactionType == "like" {
		message, err := cont.usecase.LikeBlog(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": message, "error": err})
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": message, "error": err})
		}
	} else if interactionType == "dislike" {
		message, err := cont.usecase.DislikeBlog(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": message, "error": err})
		}
	} else if interactionType == "view" {
		message, err := cont.usecase.ViewBlogs(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": message, "error": err})
		}
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "allowed:[like,view,dislike]", "error": "unknown interaction type"})
	}
}

// HandleCommentOnBlog godoc
// @Summary Add a comment to a blog post
// @Description Add a new comment to the specified blog post.
// @Tags Blog Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param blogId path string true "Blog post ID"
// @Param comment body domain.Comment true "Comment data"
// @Success 200 {object} map[string]string "Comment added successfully"
// @Failure 400 {object} map[string]interface{} "Invalid comment data"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]interface{} "Blog not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /blogs/{blogId}/comments [post]
func (cont *BlogController) HandleCommentOnBlog(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	var newComment domain.Comment

	err := ctx.ShouldBindJSON(&newComment)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	newComment.AuthorId = claims.ID
	err = cont.usecase.AddComment(blogId, newComment)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
	}
}

// HandleGetAllComments godoc
// @Summary Get all comments for a blog post
// @Description Retrieve paginated comments for the specified blog post. Defaults: pageNumber=1, pageSize=5.
// @Tags Blog Comments
// @Produce json
// @Param blogId path string true "Blog post ID"
// @Param pageNumber query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} interface{} "List of comments"
// @Failure 404 {object} map[string]interface{} "No comments found"
// @Router /blogs/{blogId}/comments [get]
func (cont *BlogController) HandleGetAllComments(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize

	comments, err := cont.usecase.GetAllComments(blogId, x)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, comments)
	}
}

// HandleGetCommentById godoc
// @Summary Get a specific comment
// @Description Retrieve a single comment by its ID for the given blog post.
// @Tags Blog Comments
// @Produce json
// @Param blogId path string true "Blog post ID"
// @Param commentId path string true "Comment ID"
// @Success 200 {object} domain.Comment "Comment details"
// @Failure 404 {object} map[string]interface{} "Comment not found"
// @Router /blogs/{blogId}/comments/{commentId} [get]
func (cont *BlogController) HandleGetCommentById(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	comments, err := cont.usecase.GetCommentById(blogId, commentId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, comments)
	}
}

// HandleCommentLikeOrDislike godoc
// @Summary Interact with a comment
// @Description Like, dislike, or view a comment by specifying the interaction type.
// @Tags Blog Comments
// @Produce json
// @Security BearerAuth
// @Param blogId path string true "Blog post ID"
// @Param commentId path string true "Comment ID"
// @Param type path string true "Interaction type (like/dislike/view)"
// @Success 200 {object} map[string]string "Interaction result"
// @Failure 400 {object} map[string]string "Invalid interaction type"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]interface{} "Comment not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /blogs/{blogId}/comments/{commentId}/interact/{type} [post]
func (cont *BlogController) HandleCommentLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}

	if interactionType == "like" {
		err := cont.usecase.LikeComment(blogId, commentId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment liked successfully"})
		}
	} else if interactionType == "dislike" {
		err := cont.usecase.DislikeComment(blogId, commentId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment disliked successfully"})
		}
	} else if interactionType == "view" {
		err := cont.usecase.ViewComment(blogId, commentId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment viewed successfully"})
		}
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid interaction type"})
	}
}

// HandleReplyOnComment godoc
// @Summary Reply to a comment
// @Description Add a new reply to an existing comment on a blog post.
// @Tags Blog Comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param blogId path string true "Blog post ID"
// @Param commentId path string true "Comment ID"
// @Param reply body domain.Reply true "Reply data"
// @Success 200 {object} map[string]string "Reply added successfully"
// @Failure 400 {object} map[string]interface{} "Invalid reply data"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]interface{} "Comment not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /blogs/{blogId}/comments/{commentId}/replies [post]
func (cont *BlogController) HandleReplyOnComment(ctx *gin.Context) {
	var newReply domain.Reply
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	err := ctx.ShouldBindJSON(&newReply)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	newReply.AuthorId = claims.ID
	err = cont.usecase.ReplyToComment(blogId, commentId, newReply)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply added successfully"})
	}
}

// HandleGetAllRepliesForComment godoc
// @Summary Get all replies for a comment
// @Description Retrieve paginated replies for the specified comment. Defaults: pageNumber=1, pageSize=5.
// @Tags Blog Comments
// @Produce json
// @Param blogId path string true "Blog post ID"
// @Param commentId path string true "Comment ID"
// @Param pageNumber query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} interface{} "List of replies"
// @Failure 404 {object} map[string]interface{} "No replies found"
// @Router /blogs/{blogId}/comments/{commentId}/replies [get]
func (cont *BlogController) HandleGetAllRepliesForComment(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize

	replies, err := cont.usecase.GetAllReplies(blogId, commentId, x)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, replies)
	}
}

// HandleGetReplyById godoc
// @Summary Get a specific reply
// @Description Retrieve a single reply by its ID for the given comment and blog post.
// @Tags Blog Comments
// @Produce json
// @Param blogId path string true "Blog post ID"
// @Param commentId path string true "Comment ID"
// @Param replyId path string true "Reply ID"
// @Success 200 {object} domain.Reply "Reply details"
// @Failure 404 {object} map[string]interface{} "Reply not found"
// @Router /blogs/{blogId}/comments/{commentId}/replies/{replyId} [get]
func (cont *BlogController) HandleGetReplyById(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	replyId := ctx.Param("replyId")
	replies, err := cont.usecase.GetReplyById(blogId, commentId, replyId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, replies)
	}
}

// HandleReplyLikeOrDislike godoc
// @Summary Interact with a reply
// @Description Like, dislike, or view a reply by specifying the interaction type.
// @Tags Blog Comments
// @Produce json
// @Security BearerAuth
// @Param blogId path string true "Blog post ID"
// @Param commentId path string true "Comment ID"
// @Param replyId path string true "Reply ID"
// @Param type path string true "Interaction type (like/dislike/view)"
// @Success 200 {object} map[string]string "Interaction result"
// @Failure 400 {object} map[string]string "Invalid interaction type"
// @Failure 401 {string} string "Unauthorized - invalid or missing token"
// @Failure 404 {object} map[string]interface{} "Reply not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /blogs/{blogId}/comments/{commentId}/replies/{replyId}/interact/{type} [post]
func (cont *BlogController) HandleReplyLikeOrDislike(ctx *gin.Context) {
	like := ctx.Param("type")
	commentId := ctx.Param("commentId")
	blogId := ctx.Param("blogId")
	replyId := ctx.Param("replyId")
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	if like == "like" {
		err := cont.usecase.LikeReply(blogId, commentId, replyId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply liked successfully"})
		}
	} else if like == "dislike" {
		err := cont.usecase.DislikeReply(blogId, commentId, replyId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply disliked successfully"})
		}
	} else if like == "view" {
		err := cont.usecase.ViewReply(blogId, commentId, replyId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply viewed successfully"})
		}
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid interaction type"})
	}
}
