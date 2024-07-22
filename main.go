package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ? Yam yu jwt key 
var Key = []byte(primitive.NewObjectID().Hex())

// ? struct chido claimard
type Claims struct {
    UserID string `json:"userID"`
    jwt.StandardClaims
}

// ? Yam func JWT generaciya chidora
func GenerateJWT(userID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)

   // ? yam JWT standartne claim , aredta mash user Id sod
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    // ? Naw token chido
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // ? Wam key ar we token rebido
    tokenString, err := token.SignedString(Key)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func main() {
    token, err := GenerateJWT("12345")
    if err != nil {
        fmt.Println("Error generating token:", err)
        return
    }

    fmt.Println("Generated Token:", token)
}
