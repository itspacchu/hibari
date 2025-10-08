package handlers

import (
	"fmt"
	"math/rand"
)

var threadRequests = []string{
	"%s, we recommend using threads for new topics.",
}

func GetRandomThreadRequest(username string) string {
	if len(threadRequests) == 0 {
		return "Error: No thread requests available."
	}
	randomIndex := rand.Intn(len(threadRequests))
	return fmt.Sprintf(threadRequests[randomIndex], username)
}
