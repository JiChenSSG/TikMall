package utils

import "fmt"

func TokenKey(userID int64) string {
	return fmt.Sprintf("token:%v", userID)
}