package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/kikils/golang-todo/interfaces/controllers"
	"google.golang.org/api/option"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opt := option.WithCredentialsFile("/.key/firebase-key.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Errorf("error initializing app: %v", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token\n"))
			return
		}

		log.Printf("Verified ID token. Access by user_id: %v\n", (*token).UID)
		ctx := controllers.SetUserID(r.Context(), (*token).UID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
