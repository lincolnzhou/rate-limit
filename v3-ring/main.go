package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("【rate-limit】v3-ring")

	resolution := time.Millisecond * 200
	fmt.Println(`resolution------->`, int64(resolution))

	now := time.Now().UnixNano()
	fmt.Println(`now------->`, now)

	slot := uint32(now / int64(resolution))
	fmt.Println(`slot------->`, slot)
}