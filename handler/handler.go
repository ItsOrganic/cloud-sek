package handler

import (
	"cloud-sek/constants"
	"cloud-sek/models"
	"cloud-sek/service"
	"cloud-sek/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var PostService *service.PostService

// @Summary		Create a new post
// @Description	Create a new post
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			post	body		models.Post	true	"Post object"
// @Success		200		{object}	map[string]interface{}
// @Failure		400		{object}	map[string]interface{}
// @Failure		500		{object}	map[string]interface{}
// @Router			/create [post]
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	post.Description = utils.ConvertToHTML(post.Description)

	if err := PostService.CreatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(200, gin.H{"message": "Post created successfully", "post": post})
}

// @Summary		Get a post by ID
// @Description	Retrieve a post by its ID
// @Tags			posts
// @Produce		json
// @Param			id	path		string	true	"Post ID"
// @Success		200	{object}	map[string]interface{}
// @Failure		404	{object}	map[string]interface{}
// @Failure		500	{object}	map[string]interface{}
// @Router			/post/{id} [get]
func GetPostById(c *gin.Context) {
	postId := c.Param("id")
	post, err := PostService.GetPostById(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		return
	}

	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, gin.H{"post": post})
}

// @Summary		Get comments by post ID
// @Description	Retrieve all comments for a specific post
// @Tags			comments
// @Produce		html
// @Param			id	path		string	true	"Post ID"
// @Success		200	{string}	html	"HTML content"
// @Failure		500	{object}	map[string]interface{}
// @Router			/post/{id}/comments [get]
func GetCommentsByPostID(c *gin.Context) {
	postId := c.Param("id")
	comments, err := PostService.GetCommentsByPostID(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	var html string
	html = fmt.Sprintf(constants.HTML_COMMENT, postId)

	for _, comment := range comments {
		html += `<div class="comment">
            <p class="author">` + comment.Author + `</p>
            <div>` + comment.Message + `</div>
        </div>`
	}

	html += `</body></html>`

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, html)
}

// @Summary		Create a new comment
// @Description	Create a new comment for a specific post
// @Tags			comments
// @Accept			json
// @Produce		json
// @Param			id		path		string			true	"Post ID"
// @Param			comment	body		models.Comment	true	"Comment object"
// @Success		200		{object}	map[string]interface{}
// @Failure		400		{object}	map[string]interface{}
// @Failure		500		{object}	map[string]interface{}
// @Router			/post/{id}/comment [post]
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	postId := c.Param("id")
	comment.PostID = postId

	comment.Message = utils.ConvertToHTML(comment.Message)
	if err := PostService.CreateComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment created successfully", "comment": comment})
}
