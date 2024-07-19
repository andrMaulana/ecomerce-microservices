package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/andrMaulana/ecomerce-microservices/api/products/models"

	"github.com/dgrijalva/jwt-go"
)

// ... (fungsi middleware lainnya)
var jwtKey = []byte("your_secret_key")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implementasikan logika autentikasi dan otorisasi di sini
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Claims.(jwt.Claims); ok {
				return token.Claims, nil
			}
			return nil, fmt.Errorf("Invalid JWT token")
		})
		if err != nil {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		// Verifikasi claims token (misalnya peran pengguna)
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		userID := claims["userID"].(string)
		userRole := claims["userRole"].(string)

		// Simpan informasi pengguna di context untuk handler berikutnya
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", userID)
		ctx = context.WithValue(ctx, "userRole", userRole)
		r = r.WithContext(ctx)

		// Panggil handler berikutnya
		next.ServeHTTP(w, r)
	})
}

func ValidateProductInputMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implementasikan validasi input produk di sini
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "Invalid product input", http.StatusBadRequest)
			return
		}

		if product.Name == "" || product.Price <= 0 || product.Inventory <= 0 {
			http.Error(w, "Invalid product data", http.StatusBadRequest)
			return
		}

		// Panggil handler berikutnya
		next.ServeHTTP(w, r)
	})
}

func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implementasikan logika logging request di sini
		userID := r.Context().Value("userID").(string)
		userRole := r.Context().Value("userRole").(string)

		logMessage := fmt.Sprintf("[Request] Method: %s, URL: %s, UserID: %s, UserRole: %s",
			r.Method, r.URL.String(), userID, userRole)
		fmt.Println(logMessage)

		// Simpan log ke file atau kirim ke layanan log eksternal (opsional)

		// Panggil handler berikutnya
		next.ServeHTTP(w, r)
	})
}
