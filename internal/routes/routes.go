package routes

import (
    "blog-api/internal/handler"
    "blog-api/internal/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // روتهای عمومی
    v1 := r.Group("/api/v1")
    {
        // Auth
        auth := v1.Group("/auth")
        {
            auth.POST("/register", handler.Register)
            auth.POST("/login", handler.Login)
        }

        // پست‌های عمومی
        v1.GET("/posts", handler.GetPosts)
        v1.GET("/posts/:slug", handler.GetPost)
    }

    // روتهای محافظت‌شده (نیاز به JWT)
    protected := v1.Group("/")
    protected.Use(middleware.JWTAuthMiddleware())
    {
        // پروفایل
        protected.GET("/me", handler.GetMe)
        protected.PUT("/me", handler.UpdateProfile)

        // مدیریت پست
        protected.POST("/posts", handler.CreatePost)
        protected.PUT("/posts/:id", handler.UpdatePost)
        protected.DELETE("/posts/:id", handler.DeletePost)
        protected.PATCH("/posts/:id/publish", handler.PublishPost)
    }

    return r
}