package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Dbm *gorp.DbMap
)

func main() {
	// initialize the DbMap
	InitDB()
	port := os.Getenv("PORT")

	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("static", "static")

	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/moreJSON1/", func(c *gin.Context) { //여기 id던져줄수 있음 /:uid
	
		mainresults, _ := Dbm.Select(mainTable{}, "SELECT \"MainID\", \"MainName\" FROM main")
		
		var articles []*mainTable
		for _, r := range mainresults {
			b := r.(*mainTable)
			articles = append(articles, b)
		}

		c.JSON(http.StatusOK, articles)
	})
	
		router.GET("/moreJSON2/", func(c *gin.Context) { //여기 id던져줄수 있음 /:uid
		
		restaurantresults, _ := Dbm.Select(RestaurantTable{}, "SELECT \"RestaurantName\", \"RestaurantHours\", \"RestaurantPosition\" FROM restaurant")
		
		var articles []*RestaurantTable
		for _, r := range restaurantresults {
			b := r.(*RestaurantTable)
			articles = append(articles, b)
		}

		c.JSON(http.StatusOK, articles)
	})
	
		router.GET("/moreJSON3/", func(c *gin.Context) { //여기 id던져줄수 있음 /:uid
		
		menuresults, _ := Dbm.Select(MenuTable{}, "SELECT \"MenuName\", \"MenuPrice\" FROM menu")
		
		var articles []*MenuTable
		for _, r := range menuresults {
			b := r.(*MenuTable)
			articles = append(articles, b)
		}

		c.JSON(http.StatusOK, articles)
	})

	router.Run(":" + port)

}

type mainTable struct {
	MainID       int // 앞에 대문자 써줘야함
	MainName     string
}


type RestaurantTable struct {
	RestaurantID       int
	RestaurantName     string
	RestaurantHours    string
	RestaurantPosition string
}

type MenuTable struct {
	MenuID       int
	MenuName     string
	MenuPrice    string
}


func InitDB() {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	databseurl := "postgres://xlwavrftsukeol:RCuJ0I8i3srvr4-_GLiD05bn4c@ec2-54-225-194-162.compute-1.amazonaws.com:5432/dabjp8n883fbfa"
	db, err := sql.Open("postgres", databseurl)
	checkErr(err, "postgres.sql.Open failed")

	// construct a gorp DbMap
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	//dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	//err = dbmap.CreateTablesIfNotExists()
	//checkErr(err, "Create tables failed")

	//return dbmap
}

type Post struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id      int64 `db:"post_id"`
	Created int64
	Title   string `db:",size:50"`               // Column size set to 50
	Body    string `db:"article_body,size:1024"` // Set both column name and size
}

func newPost(title, body string) Post {
	return Post{
		Created: time.Now().UnixNano(),
		Title:   title,
		Body:    body,
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

