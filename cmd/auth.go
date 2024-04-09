package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Password string `json:"password"`
}

func login(c *gin.Context) {
	var req Credentials
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if req.Password == "angularstart" {
		sm.Put(c.Request.Context(), "authenticated", true)
		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func logout(c *gin.Context) {
	sm.Destroy(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func Adapt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the SCS LoadAndSave middleware
		sm.LoadAndSave(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Let Gin handle the request
			c.Request = r
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	}
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if auth, ok := sm.Get(c.Request.Context(), "authenticated").(bool); !ok || !auth {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
