package middleware

import (
	"context"
	"fmt"
	"kino_backend/CSRF"
	"kino_backend/db"
	"kino_backend/logger"
	"kino_backend/repository"
	"log"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type key int

const (
	KeyIsAuthenticated key = iota
	KeySessionID
	KeyUserID
)

var (
	corsData = CorsData{
		AllowOrigins: []string{
			"localhost:8080",
		},
		AllowMethods:     []string{"GET", "DELETE", "POST", "PUT"},
		AllowHeaders:     []string{"Content-Type", "X-Content-Type-Options"},
		AllowCredentials: true,
	}
)

type CorsData struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}

func CorsMiddleware(h http.Handler) http.Handler {
	var mw http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("Request was accepted")
		val, ok := req.Header["Origin"]
		if ok {
			res.Header().Set("Access-Control-Allow-Origin", val[0])
			res.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(corsData.AllowCredentials))
		}

		if req.Method == "OPTIONS" {
			res.Header().Set("Access-Control-Allow-Methods", strings.Join(corsData.AllowMethods, ", "))
			res.Header().Set("Access-Control-Allow-Headers", strings.Join(corsData.AllowHeaders, ", "))
			return
		}

		h.ServeHTTP(res, req)
	}

	return mw
}

func RecoverMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("[PANIC]: ", err, " at ", string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func AccessLogMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		logger.Infow(r.URL.Path,
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
			"url", r.URL.Path,
			"work_time", time.Since(start).String(),
		)
	})
}

func SessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c, err := r.Cookie("session_id")
		if err == nil {
			s := repository.NewSessionsRepository(repository.Rd)
			uid, err := s.Get(r.Context(), c.Value)
			switch err {
			case nil:
				ctx = context.WithValue(ctx, KeyIsAuthenticated, true)
				ctx = context.WithValue(ctx, KeySessionID, c.Value)
				ctx = context.WithValue(ctx, KeyUserID, uid)
				tokenExpiration := time.Now().Add(24 * time.Hour)
				csrfToken, _ := CSRF.Tokens.Create(string(uid), c.Value, tokenExpiration.Unix())
				w.Header().Set("X-CSRF-Token", csrfToken)
			case db.ErrSessionNotFound:
				// delete unvalid cookie
				c.Expires = time.Now().AddDate(0, 0, -1)
				http.SetCookie(w, c)
				ctx = context.WithValue(ctx, KeyIsAuthenticated, false)
			default:
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else { // ErrNoCookie
			ctx = context.WithValue(ctx, KeyIsAuthenticated, false)
		}
		//fmt.Println("key  ", KeyIsAuthenticated)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}






