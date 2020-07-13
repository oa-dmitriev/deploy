package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func OilPriceHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	magicNumber := rand.Int()

	w.Write([]byte(strconv.Itoa(magicNumber)))
}

func main() {
	http.HandleFunc("/", OilPriceHandler)

	fmt.Println("starting service")
	log.Fatal(http.ListenAndServe("85.89.126.28:8080", nil))
}
