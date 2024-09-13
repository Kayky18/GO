#API README

Introduction
This API is built using Go programming language, Gin framework, Gorm ORM, and SQLite database. It provides a RESTful interface for interacting with the database.

#Getting Started

##Prerequisites
Go 1.17 or higher
Gin 1.7.4 or higher
Gorm 1.21.10 or higher
SQLite 3.36.0 or higher

##Installation
Clone the repository: git clone https://github.com/your-username/your-repo-name.git
Navigate to the project directory: cd your-repo-name
Install dependencies: go get
Run the API: go run main.go

##API Endpoints

Users
GET /openings: Retrieve a list of all openings
GET /opening/:id: Retrieve a single user by ID
POST /opening: Create a new user
PUT /opening/:id: Update an existing user
DELETE /opening/:id: Delete a user

##Database
The API uses SQLite as the database. The database file is located at ./db/main.db in the project root.

##Contact
If you have any questions or issues, please contact kayky.marinho01@gmail.com.


