package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 結構體代表一個用戶
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 用一個切片來模擬數據庫
var users = []User{
	{ID: "1", Name: "哈利", Age: 18},
	{ID: "2", Name: "赫敏", Age: 18},
	{ID: "3", Name: "榮恩", Age: 17},
}

func main() {
	// 創建一個默認的路由引擎
	r := gin.Default()

	// GET API: 獲取所有用戶
	r.GET("/users", getAllUsers)

	// POST API: 更新用戶信息
	r.POST("/users/:id", updateUser)

	// 運行服務器在 8080 端口
	r.Run(":8080")
}

// getAllUsers 處理獲取所有用戶的請求
func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// updateUser 處理更新用戶信息的請求
func updateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.ID == id {
			// 更新用戶信息
			users[i].Name = updatedUser.Name
			users[i].Age = updatedUser.Age
			c.JSON(http.StatusOK, users[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "用戶未找到"})
}
