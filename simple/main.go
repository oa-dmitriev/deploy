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

	w.Write([]byte("Ян переспал с " + strconv.Itoa(magicNumber) + " женщинами"))
}

func main() {
	http.HandleFunc("/", OilPriceHandler)

	fmt.Println("starting service")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
