package handlers

import (
	"context"
	"log"

	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/models"
	authzpb "github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/proto"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	authzpb.UnimplementedAuthzServiceServer
	DB        *mongo.Collection
	jwtSecret string
	Jwt       utils.JwtWrapper
}

func (s *Server) VerifyRole(ctx context.Context, req *authzpb.VerifyRoleRequest) (*authzpb.VerifyRoleResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.JwtToken)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid or expired token")
	}

	var user models.User

	// Find the user in the database
	err = s.DB.FindOne(ctx, bson.M{"username": claims.Email}).Decode(user)
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
