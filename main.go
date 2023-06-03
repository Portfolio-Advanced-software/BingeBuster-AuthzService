package main

import (
	"context"
	"log"
	"net"

	pb "your_package_path/authorization" // Update with the correct package path

	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/config"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) VerifyRole(ctx context.Context, req *pb.VerifyRoleRequest) (*pb.VerifyRoleResponse, error) {
	// Add your JWT token verification logic here
	token := req.Token
	hasRole := false

	// Perform the necessary checks to validate the user's role
	if tokenIsValid(token) {
		// Assume roles are stored in the token's claims
		claims := extractClaims(token)
		role := claims["role"].(string)

		// Check if the role is one of the allowed roles
		if role == "user" || role == "subtitle validator" || role == "movie maintainer" {
			hasRole = true
		}
	}

	return &pb.VerifyRoleResponse{HasRole: hasRole}, nil
}

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Println("gRPC server running on port", c.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Placeholder functions for JWT token verification and claims extraction
func tokenIsValid(token string) bool {
	// Implement your JWT token verification logic here
	// Return true if the token is valid, false otherwise
	return true
}

func extractClaims(token string) map[string]interface{} {
	// Implement your JWT claims extraction logic here
	// Return a map of claims extracted from the token
	return map[string]interface{}{}
}
