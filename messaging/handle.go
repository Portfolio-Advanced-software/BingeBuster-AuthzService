package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-AuthzService/mongodb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Message struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	Action string `json:"action"`
}

func HandleMessage(body []byte) error {
	jsonStr := string(body)
	var msg Message
	err := json.Unmarshal([]byte(jsonStr), &msg)
	if err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return err
	}

	switch msg.Action {
	case "deleteAllRecords":
		_, err := mongodb.DeleteAuthzByUserId(context.Background(), msg.UserId)
		if err != nil {
			log.Println("Failed to delete all records:", err)
		}
	case "saveRecord":
		// Insert the data into the database, result contains the newly generated Object ID for the new document
		_, err := globals.AuthzDb.InsertOne(globals.MongoCtx, msg)
		// check for potential errors
		if err != nil {
			// return internal gRPC error to be handled later
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
		}
	default:
		fmt.Println("Unknown action:", msg.Action)
	}

	return nil
}
