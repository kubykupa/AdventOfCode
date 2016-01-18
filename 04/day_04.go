package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func MD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	const secret string = "ckczppom"
	//const secret string = "abcdef"

	var i int = 1
	for {
		res := MD5(secret + strconv.Itoa(i))
		if strings.HasPrefix(res, "00000") {
			//if strings.HasPrefix(res, "000000") {
			fmt.Println("FOUND: ", res, " i = ", i)
			break
		}
		i++
	}
}
