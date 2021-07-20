package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.price)
	http.HandleFunc("/read", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.price)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type dollars float32

type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) create(w http.ResponseWriter, req *http.Request) {}
func (db database) read(w http.ResponseWriter, req *http.Request)   {}
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "ParseFloat error")
		return
	}
	if price < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %v\n", price)
		return
	}
	db[item] = dollars(price)
}
func (db database) delete(w http.ResponseWriter, req *http.Request) {}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
