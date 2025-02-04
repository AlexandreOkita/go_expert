package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 1)
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	userID := ctx.Value("userID")
	fmt.Println(userID)
}
