package handlers

import (
	authzpb "github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/proto"
)

type AuthzServiceServer struct {
	authzpb.UnimplementedAuthzServiceServer
}
