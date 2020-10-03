package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "Good day, today we have %s and it is %02d:%02d", t.Weekday().String(), t.Hour(), getMinute(t.Minute(), t.Second()))
}

func getMinute(minute int, second int) int {
	return minute + second/30
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
