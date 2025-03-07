package handler

import (
	"encoding/json"
	"queueAppV2/internal/config"
	"strconv"
)

const (
	serviceQuantity = 3
 	EmptyID = -1
)


// returning id int
// if bool = true => validation success, []byte is empty
// if bool = false => validation failed, []byte is error message json
func validateServiceID(idStr string) (int, []byte, bool) {
	if idStr == "" {
		data := ErrorMessage{Error: "service id must be not empty"}
		message, _ := json.Marshal(data)
		return EmptyID, message, false
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		data := ErrorMessage{Error: "service id must be a number"}
		message, _ := json.Marshal(data)
		return EmptyID, message, false
	}
	return id, nil, true
}

// returning id int
// if bool = true => validation success, []byte is empty
// if bool = false => validation failed, []byte is error message json
func validateSlotID(idStr string) (int, []byte, bool) {
	if idStr == "" {
		data := ErrorMessage{Error: "slot id must be not empty"}
		message, _ := json.Marshal(data)
		return EmptyID, message, false
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		data := ErrorMessage{Error: "slot id must be a number" }
		message, _ := json.Marshal(data)
		return EmptyID, message, false
	}
	return id, nil, true
}

// returning id int
// if bool = true => validation success, []byte is empty
// if bool = false => validation failed, []byte is error message json
func validateQueueID(idStr string) (int, []byte, bool) {
	if idStr == "" {
		data := ErrorMessage{Error: "queueID must be not empty"}
		message, _ := json.Marshal(data)
		return EmptyID, message, false
	}
	id ,err := strconv.Atoi(idStr)
	if err != nil {
		data := ErrorMessage{Error: "queueID must be a number"}
		message, _ := json.Marshal(data)
		return EmptyID, message, false
	}
	return id, nil, true
}

func validateCredentials(cfg *config.Config, data LoginData) bool {
	if cfg.Login == data.Login && cfg.Password == data.Password {
		return true
	}
	return false

}
