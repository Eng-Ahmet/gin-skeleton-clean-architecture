// package http
package http

import (
	"hwai-api/internal/delivery/dto"
	"hwai-api/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler similar to controller in MVC, it's responsible for handling the incoming HTTP requests
// and returning the response to the client
type UserHandler struct {
	usecase usecase.UserUsecase // UserUsecase interface from usecase
}

// NewUserHandler creates a new user handler with the necessary dependencies
func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// create new user handler
func (h *UserHandler) CreateUser(c *gin.Context) {

	//user dto CreateUserRequest
	var request dto.CreateUserRequest

	// binding the request body to the request struct
	// c.BindJSON is a helper function that does the binding and returns an error if the request body is not in the correct format
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if all the fields are provided
	// if any of the fields are empty, return an error to the client
	if request.FirstName == "" || request.LastName == "" || request.Password == "" || request.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// create a new user using the usecase
	// if there is an error in creating the user, return a response to the client
	// if there is no error, return the user that was created
	user, err := h.usecase.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return the user that was created
	c.JSON(http.StatusCreated, user)
}

// GetUserByID gets a user by the given ID
// it takes the ID as a parameter and returns the user or an error
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id") // get the ID from the URL parameter

	idUint64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the User ID "}) // if the ID is not correct, return an error
		return
	}
	idUint := uint(idUint64)

	user, err := h.usecase.GetUserByID(idUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // if the user is not found, return an error
		return
	}

	// return the user
	c.JSON(http.StatusOK, user)
}
