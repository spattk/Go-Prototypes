package main

import (
	"fmt"
	"sync"
	"crypto/sha256"
    "strconv"
	"math/big"
)

const (
	numBuckets = 5
)

func addUserDetails(userId int, userIdMap map[int][]int) {
	shardIndex := getShardIndex(userId)
	if _, exists := userIdMap[shardIndex]; !exists {
		userIdMap[shardIndex] = []int{}
	}
	userIdMap[shardIndex] = append(userIdMap[shardIndex], userId)
}

func contains(slice []int, target int) bool {
    for _, value := range slice {
        if value == target {
            return true
        }
    }
    return false
}

func getUserDetails(wg *sync.WaitGroup, userId int, userIdMap map[int][]int) {
	defer wg.Done()
	shardIndex := getShardIndex(userId)	
	validStr := "VALID"
	if !contains(userIdMap[shardIndex], userId) {
		validStr = "INVALID"
	}
	fmt.Printf("UserId: %v => Invoking shard %d :: %v\n", userId, shardIndex, validStr)
}

func getShardIndex(userId int) int {
	userIdStr := strconv.Itoa(userId)
	hasher := sha256.New()
	hasher.Write([]byte(userIdStr))
	hashBytes := hasher.Sum(nil)
	hashInt := new(big.Int).SetBytes(hashBytes)
	bucket := new(big.Int).Mod(hashInt, big.NewInt(int64(numBuckets)))
	return int(bucket.Int64())
}
