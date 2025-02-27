package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/newmohib/goLang-app-rss-feed-scraper/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// test feed list https://www.theguardian.com/world/rss
	// feed, err := urlToFeed("https://www.wagslane.dev/index.xml")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(feed)

	//
	myEnv, err := godotenv.Read()
	portString := myEnv["PROT"]
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// initialize database
	dbURL := myEnv["DB_URL"]
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// Open database connection
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}

	//defer conn.Close()

	fmt.Println("Successfully connected to the database!")
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	// strart go routine for feed scraper
	// wg := &sync.WaitGroup{}
	// wg.Add(1)
	go startScraping(db, 10, time.Minute)

	// added Server listening port
	fmt.Println(portString)
	router := chi.NewRouter()

	// added cours middleware
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeed)
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	v1Router.Post("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feeds_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feeds_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	// v1Router.Get("/liveness", handlerLiveness)
	router.Mount("/v1", v1Router)

	// added Server listening port
	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}
	log.Printf("Server starting on %v", srv.Addr)
	srv.ListenAndServe()
}
