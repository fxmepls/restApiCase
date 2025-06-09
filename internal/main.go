package main

import (
	"fmt"
	"net/http"

	"restApiCase/internal/db"
	"restApiCase/internal/handlers"
)

func main() {

	database := db.PgConnect()
	defer database.Close()
	db.RunMigrations(database)

	redisClient := db.RedisConnect()
	defer redisClient.Close()

	mux := handlers.SetupRoutes(database, redisClient)

	fmt.Println("The server is running on port 8080")
	fmt.Printf(`
=== API ===

1. Create user:
   curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name":"Alice"}'

   or:
   Invoke-RestMethod -Uri http://localhost:8080/users 
     -Method Post 
     -ContentType "application/json" 
     -Body '{"name":"Alice"}'

2. Get user:
   curl http://localhost:8080/users/1

3. Postman:
   - Method: POST
   - URL: http://localhost:8080/users
   - Headers: Content-Type: application/json
   - Body (raw JSON): {"name":"Alice"}

   - Method: GET
   - URL: http://localhost:8080/users/1
   - Headers: Content-Type: application/json

   - Method: PUT
   - URL: http://localhost:8080/users/1
   - Headers: Content-Type: application/json
   - Body (raw JSON): {"name":"Alice"}

4. Sum:
	- POST
	- URL: http://localhost:8080/sum?a=5&b=9

5. Multiply:
	- POST
	- URL: http://localhost:8080/multiply?a=5&b=9
`)
	http.ListenAndServe(":8080", mux)
}
