<h1>API README<h1>

Introduction
This API is built using Go programming language, Gin framework, Gorm ORM, Swagger for documentation, and SQLite database. It provides a RESTful interface for interacting with the database.

<h2>Getting Started</h2>

<h2>Prerequisites</h2>
Go 1.17 or higher</br>
Gin 1.7.4 or higher</br>
Gorm 1.21.10 or higher</br>
SQLite 3.36.0 or higher</br>
Swagger 2.0 or higher</br>

<h2>Installation</h2>
Clone the repository: git clone https://github.com/Kayky18/GO.git</br>
Navigate to the project directory: cd your-repo-name</br>
Install dependencies: go get</br>
Run the API: go run main.go</br>

<h2>API Endpoints</h2>

Users
GET /openings: Retrieve a list of all openings</br>
GET /opening/:id: Retrieve a single user by ID</br>
POST /opening: Create a new user</br>
PUT /opening/:id: Update an existing user</br>
DELETE /opening/:id: Delete a user</br>

<h2>Database</h2>
The API uses SQLite as the database. The database file is located at ./db/main.db in the project root.

<h2>Docs</h2>
The API documentation is available at http://localhost:8080/swagger/index.html.

<h2>Contact</h2>
If you have any questions or issues, please contact kayky.marinho01@gmail.com.


