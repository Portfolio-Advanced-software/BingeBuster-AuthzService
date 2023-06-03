package handlers

import (
	"context"
	"log"

	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/models"
	authzpb "github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AuthzServiceServer) VerifyRole(ctx context.Context, req *authzpb.VerifyRoleRequest) (*authzpb.VerifyRoleResponse, error) {
	// Extract the user ID from the request
	userID := req.UserId

	// Find the user in the database by user ID
	var user models.User
	err := globals.AuthzDb.FindOne(ctx, bson.M{"userid": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "User not found")
		}
		log.Printf("Failed to find user: %v", err)
		return nil, status.Error(codes.Internal, "Failed to check authorization")
	}

	// Check if the user has the required role
	if user.Role != req.Role {
		return &authzpb.VerifyRoleResponse{IsAuthorized: false}, nil
	}

	return &authzpb.VerifyRoleResponse{IsAuthorized: true}, nil
}
