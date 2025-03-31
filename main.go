package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    //Load env vars
    err := godotenv.Load()
    if err != nil {
        log.Printf("No .env file found or couldn't load it. Proceeding with system env vars.")
    }

    //Conect to db
    db, err := ConnectDB()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
    defer db.Close()

    //Set up gin router
    r := gin.Default()

    // Test endpoin
    r.GET("/get-users", func(c *gin.Context) {
		//User struct
        type User struct {
            ID    int    `json:"id"`
            Name  string `json:"name"`
            Email string `json:"email"`

        }
		//Query to get users
        rows, err := db.Query("SELECT userID, username, email FROM users")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("DB query failed: %v", err)})
            return
        }
        defer rows.Close()

        var users []User
        for rows.Next() {
            var u User
            if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to scan user: %v", err)})
                return
            }
            users = append(users, u)
        }

        if err := rows.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Row iteration error: %v", err)})
            return
        }

        c.JSON(http.StatusOK, gin.H{"users": users})
    })

    //Start server on port 8080
    r.Run(":8080") 
}
