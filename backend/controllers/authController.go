package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"onebounce_backend/database"
	"os"
	"time"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "auth")
var validate = validator.New()

type userEntry struct {
	UUID string
	Email string
	Password string
}

func hashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func generateToken(c *gin.Context, UUID string) (signedToken string) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: UUID,
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(), //Expires after 30 minutes
	})

	token, err := claims.SignedString([]byte(SECRET_KEY))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	return token
}

func Register(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var user userEntry

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while checking for the email"})
		log.Panic(err)
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This email already exists"})
		return
	}

	password := hashPassword(user.Password)
	user.Password = password
	user.UUID = uuid.New().String()

	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, resultInsertionNumber)
}

func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var user userEntry
	var foundUser userEntry

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		return
	}

	token := generateToken(c, foundUser.UUID)
	c.SetCookie("jwt-auth", token, 60*30, "/", "onebounce.me", true, true) //Change secure to true when on port 443
	c.JSON(http.StatusOK, gin.H{"message": "successfully signed in"})
}

func Authorized(c *gin.Context) bool {
	cookie, err := c.Request.Cookie("jwt-auth")
	if err != nil {
		return false
	}

	_, err = jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return false
	}
	return true
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt-auth", "", 60*30, "/", "onebounce.me", true, true) //Change secure to true when on port 443
	c.JSON(http.StatusOK, gin.H{"message": "successfully signed out"})
}