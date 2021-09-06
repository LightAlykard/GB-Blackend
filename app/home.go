package app

import (
	"fmt"
	"net/http"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "This is simple shortener service!\n")
	}
}
