package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Agata-Liberty7/Agata-integrations/stats"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	// db           = make(map[string]interface{})
	// dbTimestemps = make(map[string]string)
	// dbClients = make(map[string]string)
	// HOST  = "178.62.240.212:8080"

	// // WORLD
	// HOST = "178.62.240.212:80"
	// // DB_HOST = "178.62.240.212"
	// DB_HOST = "127.0.0.1"

	// RUS
	HOST    = "80.249.144.41:80"
	DB_HOST = "80.249.144.41"

	debug = false
)

func init() {
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.Parse()
	// if debug {
	HOST = "0.0.0.0:8080"
	DB_HOST = "127.0.0.1"
	// DB_HOST = "178.62.240.212"
	// }
}

func index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/app")
}

func setupRouter(r *gin.Engine) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(gzip.Gzip(gzip.BestCompression))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Web routes
	r.GET("/", index)

	r.StaticFile("/favicon.ico", "web/dist/favicon.ico")
	r.StaticFile("/app", "web/dist/index.html")

	r.Static("/js", "web/dist/js")
	r.Static("/css", "web/dist/css")

	// Manual input template
	r.StaticFile("/template.xlsx", "template.xlsx")

	// Metadata
	r.GET("/meta", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.MetaDB)
	})

	r.GET("/meta/:section/:location/:period", func(c *gin.Context) {
		section := c.Params.ByName("section")
		location := c.Params.ByName("location")
		period := c.Params.ByName("period")

		for i, meta := range stats.MetaDB {
			if meta.Section == section && meta.Location == location {
				stats.MetaDB[i].Period = period
				break
			}
		}

		c.JSON(http.StatusOK, stats.MetaDB)
	})

	// // Clint routes
	// clients.InitRoutes(r)

	// // Fedstat routes
	// fedstat.InitRoutes(r)

	// // Eurostat routes
	// eurostat.InitRoutes(r)

	// All stats routes
	s := r.Group("/stats")
	stats.InitRoutes(s)
}

func main() {
	if !debug {
		go func() {
			log.Println("eurostat loading started")
			// eurostat.LoadAll()
			log.Println("eurostat loading finished")
			log.Println("fedstat loading started")
			// fedstat.LoadAll()
			log.Println("fedstat loading finished")
		}()
	}

	var (
		connStr = "host=" + DB_HOST + " port=5432 dbname=agata sslmode=disable"
		// connStr = os.Getenv("DB_CS")
		spn = fmt.Sprintf("user=%s password=admin", "postgres")
	)

	r := gin.Default()

	db, err := OpenPGorm(connStr, spn, debug)
	if err != nil {
		log.Println("failed to connect database")
	} else {
		if debug {
			db = db.Debug()
		}

		db.AutoMigrate(&types.Stat{})

		r.Use(
			WithDB(db),
		)
	}

	setupRouter(r)

	r.Run(HOST)
}

func WithDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ctxDatabase", db)
		c.Next()
	}
}

// OpenPGorm - открыть соединение с БД postgres
func OpenPGorm(
	connString, user string,
	isDebug bool,
) (db *gorm.DB, err error) {

	connString += " " + user
	if isDebug {
		log.Println("wanna open", connString)
	}

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
