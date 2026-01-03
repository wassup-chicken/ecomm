package server

import (
	"log"
	"net/http"
	"time"
)

func (srv *JobServer) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"Method: %s | Path: %s | Duration: %v\n",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

func (srv *JobServer) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			log.Println("no auth to split")

			// return
		}

		// token := strings.Split(auth, " ")

		// if len(token) == 0 {
		// 	log.Println("no token provided.. either route them to register")

		// 	return
		// }
		// err := srv.Firebase.VerifyIDToken(r.Context(), token[1])

		// if err != nil {
		// 	log.Println("User unauthorized but that's ok for now.. will fix!:", err)
		// 	// w.WriteHeader(http.StatusUnauthorized)
		// 	// w.Write([]byte("User unauthorized..!"))
		// 	// return
		// }

		next.ServeHTTP(w, r)
	})
}
