package main

import (
	"context"
	"log"
	"net/http"

	"fmt"
	pbAddress "github.com/dwiprastyoisworo/go-grpc-simple/proto/address"
	pbUser "github.com/dwiprastyoisworo/go-grpc-simple/proto/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserHandler struct {
	addressClient pbAddress.AddressServiceClient
}

func NewUserHandler() *UserHandler {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &UserHandler{
		addressClient: pbAddress.NewAddressServiceClient(conn),
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	// Get user data (contoh statis)
	user := &pbUser.UserResponse{
		UserId: userID,
		Name:   "John Doe",
		Email:  "john@example.com",
	}

	// Get address via gRPC
	address, err := h.addressClient.GetAddressByUserID(context.Background(),
		&pbAddress.AddressRequest{UserId: userID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get address",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"address": address,
	})
}

func main() {
	router := gin.Default()
	userHandler := NewUserHandler()

	router.GET("/users/:id", userHandler.GetUser)

	fmt.Println("User Service running on :8080")
	log.Fatal(router.Run(":8080"))
}
