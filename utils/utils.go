package utils

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("You must set your '%s' environmental variable.", key)
	}
	return val
}
