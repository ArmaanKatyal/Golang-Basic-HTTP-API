package main

// These are all the imported libraries
import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	Author      = "Armaan Katyal"
	Version     = "1.0.0"
	ReleaseDate = "2021-07-09"
)

// User is a struct that represents a user in out struct
// The names of the structs and values inside them have to be capital such that they can be exported and used anywhere in the file
type User struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Post is a struct that represents a single Post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

// we have just created a temporary datastructure to store data but we can use database to make it permanent
// We can use various kind of Databases(Sql or Non-Sql)
// MongoDb, MySql, Sqlite, Postgres etc.
var posts []Post = []Post{}

// This is the main function which runs the application
func main() {
	// Create the main Router
	router := mux.NewRouter()

	//Handlers
	router.HandleFunc("/posts", addPost).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{index}", getSinglePosts).Methods("GET")
	router.HandleFunc("/posts/{index}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{index}", deletePost).Methods("DELETE")
	router.HandleFunc("/posts/{index}", patchPost).Methods("PATCH")

	// Listener
	http.ListenAndServe(":5000", router)
}

// Write all the posts in the HTTP Response Body
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Write the post on the index that is passed in the parameter
func getSinglePosts(w http.ResponseWriter, r *http.Request) {
	// get index of the post from the route parameter
	var indexParam string = mux.Vars(r)["index"]
	w.Header().Set("Content-Type", "application/json")
	index, err := strconv.Atoi(indexParam)
	if err != nil {
		// there was an err
		w.WriteHeader(400)
		w.Write([]byte("Index could not be converted to integer"))
		return
	}

	// error checking
	if index >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified Index"))
		return
	}

	// get the post from the index
	post := posts[index]
	json.NewEncoder(w).Encode(post)
}

// Add post in the slice that we have created above
func addPost(w http.ResponseWriter, r *http.Request) {
	// get the item value from the JSON body

	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	// This add the new post in the slice.
	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")

	// Sends the message that the Data has been Added successfully
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{"Data Added Successfully"})
}

// Updates the Data we already in the slice
func updatePost(w http.ResponseWriter, r *http.Request) {
	// get the index of the post from the route parameters
	indexParam := mux.Vars(r)["index"]
	index, err := strconv.Atoi(indexParam)
	if err != nil {
		// there was an err
		w.WriteHeader(400)
		w.Write([]byte("Index could not be converted to integer"))
		return
	}

	// error checking
	if index >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified Index"))
		return
	}

	// get the value from the JSON body
	var updatedPost Post
	// Decode the JSON Body
	json.NewDecoder(r.Body).Decode(&updatedPost)
	posts[index] = updatedPost

	// Message after update is completed
	w.Header().Set("Content-Type", "application/json")

	// sends the response that we have updated the post
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{"Data Updated Successfully"})
}

// This func is for the internal usage in the deletePost func such that we can easily delete elements from the slice
func remove(slice []Post, s int) []Post {
	return append(slice[:s], slice[s+1:]...)
}

// Delete the post from the slice for which the index has been passed in the parameters
func deletePost(w http.ResponseWriter, r *http.Request) {
	indexParam := mux.Vars(r)["index"]
	index, err := strconv.Atoi(indexParam)
	if err != nil {
		// there was an err
		w.WriteHeader(400)
		w.Write([]byte("Index could not be converted to integer"))
		return
	}

	// error checking
	if index >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified Index"))
		return
	}

	// In this we pass the updated slice to variable we created initially in the code
	posts = remove(posts, index)

	// Message after post has been deleted from the slice
	w.Header().Set("Content-Type", "application/json")

	// Sends the Response that the post with the index passed as parameter has been deleted from the slice.
	json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{"Data Deleted Successfully"})

}

func patchPost(w http.ResponseWriter, r *http.Request) {
	// get the ID of the post from the route parameters
	var indexParam string = mux.Vars(r)["index"]
	index, err := strconv.Atoi(indexParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Index could not be converted to integer"))
		return
	}

	// error checking
	if index >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified Index"))
		return
	}

	// get the current value
	post := &posts[index]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
