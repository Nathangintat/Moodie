package app

import (
	"github.com/Nathangintat/Moodie/config"
	"github.com/Nathangintat/Moodie/internal/adapter/handler"
	"github.com/Nathangintat/Moodie/internal/adapter/repository"
	"github.com/Nathangintat/Moodie/internal/core/service"
	"github.com/Nathangintat/Moodie/lib/auth"
	"github.com/Nathangintat/Moodie/lib/middleware"
	"github.com/Nathangintat/Moodie/lib/pagination"

	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	/*"github.com/gofiber/contrib/swagger"*/
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"os"
)

func RunServer() {
	cfg := config.NewConfig()
	db, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal("Error connecting to database: %v", err)
		return
	}

	err = os.MkdirAll("./temp/content", 0755)
	if err != nil {
		log.Fatal("Error creating temp directory: %v", err)
		return
	}

	jwt := auth.NewJwt(cfg)
	middlewareAuth := middleware.NewMiddleware(cfg)

	_ = pagination.NewPagination()

	// Repository
	authRepository := repository.NewAuthRepository(db.DB)
	userRepository := repository.NewUserRepository(db.DB)
	movieRepository := repository.NewMovieRepository(db.DB)
	reviewRepository := repository.NewReviewRepository(db.DB)
	voteRepository := repository.NewVoteRepository(db.DB)
	playlistRepository := repository.NewPlaylistRepository(db.DB)

	// Service
	authService := service.NewAuthService(authRepository, cfg, jwt)
	userService := service.NewUserService(userRepository)
	movieService := service.NewMovieService(movieRepository)
	reviewService := service.NewReviewService(reviewRepository)
	voteService := service.NewVoteService(voteRepository)
	playlistService := service.NewPlaylistService(playlistRepository)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	movieHandler := handler.NewMovieHandler(movieService)
	reviewHandler := handler.NewReviewHandler(reviewService)
	voteHandler := handler.NewVoteHandler(voteService)
	playlistHandler := handler.NewPlaylistHandler(playlistService)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip} ${status} - ${latency} ${method} ${path}\n",
	}))

	api := app.Group("/api")
	api.Post("/", movieHandler.GetMovies)
	api.Post("/login", authHandler.Login)
	api.Post("/register", userHandler.Register)

	adminApp := api.Group("/admin")
	adminApp.Use(middlewareAuth.CheckToken())

	// User
	userApp := adminApp.Group("/users")
	userApp.Get("/profile", userHandler.GetUserByID)
	userApp.Put("/update-password", userHandler.UpdatePassword)

	//review
	reviewApp := api.Group("/review")
	reviewApp.Use(middlewareAuth.CheckToken())
	reviewApp.Post("/create", reviewHandler.CreateReview)
	reviewApp.Post("/:reviewID/upvote", voteHandler.AddUpvote)
	reviewApp.Post("/:reviewID/downvote", voteHandler.AddDownvote)

	// movie
	movieApp := api.Group("/movie")
	movieApp.Use(middlewareAuth.CheckToken())
	movieApp.Get("/", movieHandler.GetMovies)
	movieApp.Get("/review/:movieID", reviewHandler.GetReviewByID)
	movieApp.Get("/:movieID", movieHandler.GetMovieByID)

	//playlist
	playlistApp := api.Group("/playlist")
	playlistApp.Use(middlewareAuth.CheckToken())
	playlistApp.Get("/", playlistHandler.GetPlaylistByID)
	playlistApp.Post("/create", playlistHandler.CreatePlaylist)
	playlistApp.Post("/insert", playlistHandler.InsertMovie)
	playlistApp.Get("/:playlistID/item", playlistHandler.GetPlaylistMovies)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting server on port:", port)

	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	<-quit

	log.Println("server shutdown of 5 seconds")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.ShutdownWithContext(ctx)
}
