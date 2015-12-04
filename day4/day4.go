package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	text := ""
	for i := 0; i < 1000000000; i++ {
		text = "ckczppom" + strconv.Itoa(i)
		hasher := md5.New()
		hasher.Write([]byte(text))
		hash := hex.EncodeToString(hasher.Sum(nil))
		if hash[0:6] == "000000" {
			fmt.Println(hash)
			fmt.Println(text)
			break
		}
	}
}
