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
		// just for demonstration, this needs to be securely generated and unique
		sessionID := "thisIsNotSecureDoNotUseThis"

		// store this session on the server
		sessionStore[sessionID] = true

		// set a cookie with the sessionID on the client
		c.SetCookie("session_id", sessionID, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func logout(c *gin.Context) {
	// destroy session
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session not found"})
		return
	}

	// clear session on server
	delete(sessionStore, sessionID)

	// clear cookie on client
	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil || !sessionStore[sessionID] {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
