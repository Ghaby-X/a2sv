
package Infrastructure

import (
	"fmt"
	"os"
	"time"

	"github.com/Ghaby-X/task_manager/Domain"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateJWT(user Domain.User) (string, error)
	ValidateJWT(tokenString string) (*Claims, error)
}

// Claims defines the JWT claims structure
type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	jwtKey []byte
}

func NewJWTService() JWTService {
	return &jwtService{
		jwtKey: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (s *jwtService) GenerateJWT(user Domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID.Hex(),
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

func (s *jwtService) ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
