<h1>API README<h1>

Introduction
This API is built using Go programming language, Gin framework, Gorm ORM, and SQLite database. It provides a RESTful interface for interacting with the database.

<h2>Getting Started</h2>

<h2>Prerequisites</h2>
Go 1.17 or higher
Gin 1.7.4 or higher
Gorm 1.21.10 or higher
SQLite 3.36.0 or higher

<h2>Installation</h2>
Clone the repository: git clone https://github.com/Kayky18/GO.git
Navigate to the project directory: cd your-repo-name
Install dependencies: go get
Run the API: go run main.go

<h2>API Endpoints</h2>

Users
GET /openings: Retrieve a list of all openings</br>
GET /opening/:id: Retrieve a single user by ID</br>
POST /opening: Create a new user</br>
PUT /opening/:id: Update an existing user</br>
DELETE /opening/:id: Delete a user</br>

<h2>Database</h2>
The API uses SQLite as the database. The database file is located at ./db/main.db in the project root.

<h2>Contact</h2>
If you have any questions or issues, please contact kayky.marinho01@gmail.com.


