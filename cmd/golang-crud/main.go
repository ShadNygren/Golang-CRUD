package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

// Post represents a database entry
type Post struct {
    ID    int
    Title string
    Body  string
}

const (
    host     = "<YOUR_DB_HOST>"
    port     = 5432 // Default port
    user     = "<YOUR_DB_USER>"
    password = "<YOUR_DB_PASSWORD>"
    dbname   = "<YOUR_DB_NAME>"
)

func main() {
    // Initialize connection string.
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Connect to the database.
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check the connection.
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected!")

    // CRUD Operations
    createPost(db, Post{Title: "First Post", Body: "This is the first post"})
    readPost(db, 1)
    updatePost(db, 1, Post{Title: "Updated First Post", Body: "This is the updated first post"})
    deletePost(db, 1)
}

func createPost(db *sql.DB, post Post) {
    sqlStatement := `INSERT INTO posts (title, body) VALUES ($1, $2) RETURNING id`
    id := 0
    err := db.QueryRow(sqlStatement, post.Title, post.Body).Scan(&id)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("New record ID is:", id)
}

func readPost(db *sql.DB, id int) {
    var post Post
    sqlStatement := `SELECT id, title, body FROM posts WHERE id=$1;`
    row := db.QueryRow(sqlStatement, id)
    err := row.Scan(&post.ID, &post.Title, &post.Body)
    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
    case nil:
        fmt.Println(post)
    default:
        log.Fatal(err)
    }
}

func updatePost(db *sql.DB, id int, post Post) {
    sqlStatement := `UPDATE posts SET title = $2, body = $3 WHERE id = $1;`
    _, err := db.Exec(sqlStatement, id, post.Title, post.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Record updated successfully!")
}

func deletePost(db *sql.DB, id int) {
    sqlStatement := `DELETE FROM posts WHERE id = $1;`
    _, err := db.Exec(sqlStatement, id)
    if err != nil {
    log.Fatal(err)
    }
    fmt.Println("Record deleted successfully!")
}
