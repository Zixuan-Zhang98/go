package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("times")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "times",
			Value: "0",
			Path:  "/",
		}
	} else {

	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}

	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}

// Using cookies, track how many times a user has been to your website domain.
