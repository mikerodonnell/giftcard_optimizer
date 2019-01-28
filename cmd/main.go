package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mikerodonnell/giftcard_optimizer/pkg/fileio"
	"github.com/mikerodonnell/giftcard_optimizer/pkg/gift"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 2 {
		log.Fatalf("price file and gift card balance are required. ex:  `go run cmd/main.go prices.txt 5000`")
	}

	lines, err := fileio.ReadLinesFromFile(arguments[0])
	if err != nil {
		log.Fatal(err)
	}

	if len(lines) < 2 {
		log.Fatalf("at least 2 items require in input file, found %d", len(lines))
	}

	list, err := gift.NewGiftList(lines)
	if err != nil {
		log.Fatal("failed to parse input file", err)
	}

	balance, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Fatalf("gift card balance must be numeric: %d", balance)
	}

	cheap, expensive := list.Optimize(balance)
	// GiftList.Optimize() returns (nil, nil) if there's no suitable combination
	if cheap == nil || expensive == nil {
		fmt.Println("Not possible")
	}

	fmt.Println(fmt.Sprintf("%s, %s", cheap, expensive))
}
