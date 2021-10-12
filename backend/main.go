package main

import (
	"context"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"onebounce_backend/controllers"
	"onebounce_backend/database"
)

var urlsCollection *mongo.Collection = database.OpenCollection(database.Client, "urls")
var validate = validator.New()

type urlEntry struct {
	Id string
	Url string
	Clicks int
}

type shortenEntry struct {
	Url string
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.LoadHTMLGlob("../dist/*.html")
	r.Static("/static", "../dist/static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/login", func(c *gin.Context) {
		if controllers.Authorized(c) {
			c.Redirect(http.StatusFound, "https://onebounce.me/panel")
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		}
	})

	r.GET("/panel", func(c *gin.Context) {
		if !controllers.Authorized(c) {
			c.Redirect(http.StatusFound, "https://onebounce.me/login")
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		}
	})

	r.POST("/api/auth/register", func(c *gin.Context) {
		controllers.Register(c)
	})

	r.POST("/api/auth/login", func(c *gin.Context) {
		controllers.Login(c)
	})

	r.GET("/api/auth/logout", func(c *gin.Context) {
		controllers.Logout(c)
	})

	r.POST("/api/shorten", func(c *gin.Context) {
		var data shortenEntry
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(data)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		id := GetMD5Hash(data.Url)

		opts := options.FindOneAndUpdate().SetUpsert(true)
		filter := bson.M{"id": id}
		update := bson.D{{"$setOnInsert", bson.M{"id": id, "url": data.Url, "clicks": 0}}}
		var result urlEntry
		insertErr := urlsCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)
		if insertErr != nil {
			msg := fmt.Sprintf("Url item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		shortUrlString := "onebounce.me/"+result.Id
		c.JSON(http.StatusOK, gin.H{"status": "success", "shorturl": shortUrlString})
	})

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var result urlEntry
		query := bson.D{{"id", id}}
		err := urlsCollection.FindOne(context.TODO(), query).Decode(&result)
		if err != nil {
			switch err {
			case mongo.ErrNoDocuments:
				c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
			default:
				panic(err)
			}
		}else {
			c.Redirect(http.StatusMovedPermanently, result.Url)
		}
	})

	r.Run("10.128.0.3:5000")
}

func GetMD5Hash(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])[0:5] //Return first 5 characters of the MD5 hash
}
