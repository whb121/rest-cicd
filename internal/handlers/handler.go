package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	{ID: 4, Name: "David", Email: "david@example.com"},
	{ID: 5, Name: "Eve", Email: "eve@example.com"},
	{ID: 6, Name: "Frank", Email: "frank@example.com"},
	{ID: 7, Name: "Grace", Email: "grace@example.com"},
	{ID: 8, Name: "Helen", Email: "helen@example.com"},
	{ID: 9, Name: "Irene", Email: "irene@example.com"},
	{ID: 10, Name: "James", Email: "james@example.com"},
	{ID: 11, Name: "Kevin", Email: "kevin@example.com"},
	{ID: 12, Name: "Lucas", Email: "lucas@example.com"},
	{ID: 13, Name: "Mary", Email: "mary@example.com"},
	{ID: 14, Name: "Nancy", Email: "nancy@example.com"},
	{ID: 15, Name: "Olivia", Email: "olivia@example.com"},
	{ID: 16, Name: "Peter", Email: "peter@example.com"},
	{ID: 17, Name: "Quinn", Email: "quinn@example.com"},
	{ID: 18, Name: "Rachel", Email: "rachel@example.com"},
	{ID: 19, Name: "Sophia", Email: "sophia@example.com"},
	{ID: 20, Name: "Thomas", Email: "thomas@example.com"},
	{ID: 21, Name: "William", Email: "william@example.com"},
	{ID: 22, Name: "Xavier", Email: "xavier@example.com"},
	{ID: 23, Name: "Yvonne", Email: "yvonne@example.com"},
	{ID: 24, Name: "Zoe", Email: "zoe@example.com"},
	{ID: 25, Name: "Andrew", Email: "andrew@example.com"},
	{ID: 26, Name: "Benjamin", Email: "benjamin@example.com"},
	{ID: 27, Name: "Charles", Email: "charles@example.com"},
	{ID: 28, Name: "Daniel", Email: "daniel@example.com"},
	{ID: 29, Name: "Edward", Email: "edward@example.com"},
	{ID: 30, Name: "Frank", Email: "frank@example.com"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": user})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成唯一ID：找到当前最大ID并加1
	maxID := 0
	for _, user := range users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	newUser.ID = maxID + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    newUser,
	})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			c.JSON(http.StatusOK, gin.H{
				"message": "User updated successfully",
				"data":    updatedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
