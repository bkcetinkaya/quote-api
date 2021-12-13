package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomInt(documentCount int64) int {

	//var documentCount = db.DocumentCount

	rand.Seed(time.Now().UnixNano())

	randomInt := rand.Intn(int(documentCount)) + 1

	fmt.Println(randomInt)
	return randomInt
}
