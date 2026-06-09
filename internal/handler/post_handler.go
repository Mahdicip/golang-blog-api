package handler

import (
	"net/http"

	"blog-api/config"
	model "blog-api/internal/models"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
}

// Create Post
func CreatePost(c *gin.Context) {
	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user_id not found in token",
		})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user_id type",
		})
		return
	}

	post := model.Post{
		Title:   input.Title,
		Content: input.Content,
		Slug:    input.Slug,
		Status:  "draft",
		UserID:  userID,
	}

	err := config.DB.Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, post)
}
// Get All Posts
func GetPosts(c *gin.Context) {
	var posts []model.Post

	if err := config.DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// Get Single Post
func GetPost(c *gin.Context) {
	slug := c.Param("slug")

	var post model.Post

	if err := config.DB.Preload("User").
	Where("slug = ?", slug).
	First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// Update Post
func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var post model.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	post.Slug = input.Slug

	config.DB.Save(&post)

	c.JSON(http.StatusOK, post)
}

// Delete Post
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post model.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	config.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}

// Publish Post
func PublishPost(c *gin.Context) {
	id := c.Param("id")

	var post model.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	post.Status = "published"

	config.DB.Save(&post)

	c.JSON(http.StatusOK, post)
}