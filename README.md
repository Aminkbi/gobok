# Gobok - Book Management Web Service

Gobok is a web service developed using Go Fiber and MongoDB to facilitate the management of books. It provides a simple and efficient CRUD (Create, Read, Update, Delete) interface for book-related operations. 
## Features

- **Create:** Add new books to your collection with ease.
- **Read:** Retrieve information about existing books.
- **Update:** Modify book details as needed.
- **Delete:** Remove books from your database when necessary.

## Technologies Used

- [Go Fiber](https://github.com/gofiber/fiber): A fast and lightweight web framework for Go.
- [MongoDB](https://www.mongodb.com/): A NoSQL database for storing book data.

## Prerequisites

Make sure you have the following installed on your system:

- Go
- MongoDB

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/aminkbi/Gobok.git
   cd Gobok
   ```

2. **Install dependencies:**

   ```bash
   go get -u
   ```

3. **Set up your MongoDB database and update the connection string in the `config/config.go` file.**

4. **Run the application:**

   ```bash
   go run main.go
   ```

## API Endpoints

- **GET /book:** Retrieve a list of all books.
- **GET /book/{id}:** Retrieve details of a specific book.
- **POST /book:** Add a new book to the database.
- **PUT /book/{id}:** Update details of a specific book.
- **DELETE /book/{id}:** Delete a book from the database.

