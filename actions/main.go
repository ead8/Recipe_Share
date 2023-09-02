package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gin-cloudinary-api/controllers"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	HASURA_URL            = "http://localhost:8080/v1/graphql"
	HASURA_HEADERS_SECRET = "myadminsecretkey" 
	HASURA_JWT_SECRET     = "my-long-and-secure-jwt-secret-key"
)

func MapToJSON(data map[string]interface{}) ([]byte, error) {
	return json.Marshal(data)
}

var (
	hclient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
)

// Client represents the GraphQL client
type Client struct {
	URL     string
	Headers map[string]string
}

func (c *Client) runQuery(query string, variables map[string]interface{}) ([]byte, error) {
	reqBody := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}
	reqJSON, err := MapToJSON(reqBody)
	if err != nil {
		return nil, err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}

	// Add the X-Hasura-Admin-Secret header
	req.Header.Set("X-Hasura-Admin-Secret", HASURA_HEADERS_SECRET)
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	resp, err := hclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}


func findUserByEmail(email string) ([]byte, error) {
	client := &Client{URL: HASURA_URL, Headers: map[string]string{"X-Hasura-Admin-Secret": HASURA_HEADERS_SECRET}}
	query := `
		query UserByEmail($email: String!) {
			users(where: {email: {_eq: $email}}, limit: 1) {
				id
				email
				password
			}
		}
	`
	variables := map[string]interface{}{
		"email": email,
	}
	return client.runQuery(query, variables)
}

func findUserByID(id int) ([]byte,error){
	client:=&Client{URL: HASURA_URL,Headers: map[string]string{"X-Hasura-Admin-Secret":HASURA_HEADERS_SECRET}}
	query := `
	query UserByID($id: Int!) {
		users_by_pk(id: $id) {
			id
			email
			password
		}
	}
`
	variables:=map[string]interface{}{
		"id":id,
	}
	return client.runQuery(query,variables)
}
// CreateUserOutput represents the output of the create user mutation
type CreateUserOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createUser(username string, email string, password string) (*CreateUserOutput, error) {
	client := &Client{URL: HASURA_URL, Headers: map[string]string{"X-Hasura-Admin-Secret": HASURA_HEADERS_SECRET}}
	query := `
		mutation CreateUser($username:String!,$email: String!, $password: String!) {
			insert_users_one(object: {username:$username, email: $email, password: $password}) {
				id
				username
				email
				password
			}
		}
	`
	variables := map[string]interface{}{
		"username": username,
		"email":    email,
		"password": password,
	}

	responseData, err := client.runQuery(query, variables)

	fmt.Println("Response Data:", string(responseData)) // Add this line for debugging

	if err != nil {
		return nil, err
	}

	// Define a struct to match the expected response
	type createUserResponseStruct struct {
		InsertUsersOne CreateUserOutput `json:"insert_users_one"`
	}

	// Unmarshal the response into the struct
	var createUserResponse createUserResponseStruct

	err = json.Unmarshal(responseData, &createUserResponse)
	
	if err != nil {
		return nil, err
	}

	return &createUserResponse.InsertUsersOne, nil

}


func updatePassword(id int, password string) ([]byte, error) {
	client := &Client{URL: HASURA_URL, Headers: map[string]string{"X-Hasura-Admin-Secret": HASURA_HEADERS_SECRET}}
	query := `
		mutation UpdatePassword($id: Int!, $password: String!) {
			update_users_by_pk(pk_columns: {id: $id}, _set: {password: $password}) {
				password
			}
		}
	`
	variables := map[string]interface{}{
		"id":       id,
		"password": password,
	}
	return client.runQuery(query, variables)
}

func generateToken(user map[string]interface{}) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["https://hasura.io/jwt/claims"] = map[string]interface{}{
		"x-hasura-allowed-roles": []string{"user"},
		"x-hasura-default-role": "user",
		"x-hasura-user-id":      user["id"],
	}

	// Set token expiration time
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and return it
	tokenString, err := token.SignedString([]byte(HASURA_JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func rehashAndSavePasswordIfNeeded(user map[string]interface{}, plaintextPassword string) {
	if needsRehash(user["password"].(string)) {
		updatePassword(int(user["id"].(float64)), hashPassword(plaintextPassword))
	}
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func needsRehash(hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte{})
	return err == bcrypt.ErrMismatchedHashAndPassword
}

// AuthArgs represents the input arguments for signup/login
type AuthArgs struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signupHandler(c *gin.Context) {
	
	var args AuthArgs
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	hashedPassword := hashPassword(args.Password)
	userData, err := createUser(args.Username, args.Email, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, userData)
}

func loginHandler(c *gin.Context) {
	var args AuthArgs
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	userData, err := findUserByEmail(args.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch user"})
		return
	}

	// Check if the user map and "data" key exist
	var userResponse map[string]interface{}
	if err := json.Unmarshal(userData, &userResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user data"})
		return
	}

	data, ok := userResponse["data"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user data"})
		return
	}

	users, ok := data["users"].([]interface{})
	if !ok || len(users) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user data"})
		return
	}

	userObj, ok := users[0].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user data"})
		return
	}

	passwordValue, ok := userObj["password"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid user data"})
		return
	}

	fmt.Println("Received login request with args:", passwordValue)

	if err := bcrypt.CompareHashAndPassword([]byte(passwordValue), []byte(args.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// Generate JWT token and return it
	token, err := generateToken(userObj)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func refreshTokenHandler(c *gin.Context) {
	// Get the token from the request header
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authorization header is missing"})
		return
	}

	// Extract the token from the "Bearer <token>" format
	tokenString := strings.Split(authorizationHeader, " ")[1]

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		// Return the secret key used for signing the token
		return []byte(HASURA_JWT_SECRET), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	// Check if the token is valid
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	// Extract the user ID from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse token claims"})
		return
	}
	userID := int(claims["https://hasura.io/jwt/claims"].(map[string]interface{})["x-hasura-user-id"].(float64))

	// Fetch the user from the database based on the user ID
	userData, err := findUserByID(userID) // Implement your own function to fetch the user by ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch user"})
		return
	}

	// Check if the user exists
	if userData == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}

	// Unmarshal the userData into a map
	var userMap map[string]interface{}
	if err := json.Unmarshal(userData, &userMap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse user data"})
		return
	}

	// Generate a new access token with an extended expiration time using the userMap
	accessToken, err := generateToken(userMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}



func main() {
	router := gin.Default()
    router.POST("/refresh-token", refreshTokenHandler)
	router.Use(cors.Default())

	router.POST("/signup", signupHandler)
	router.POST("/login", loginHandler)
	
	//image upload
	router.POST("/file", controllers.MultipleFileUpload())
	router.POST("/remote", controllers.RemoteUpload())
	
	err := router.Run(":3000")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}