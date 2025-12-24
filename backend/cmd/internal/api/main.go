package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sraraa/cors"
	"sraraa/db"
	access_auth_controller "sraraa/reciever_src/controllers/auth/access_auth"
	"sraraa/reciever_src/routes/auth/access_auth_routes"
	login_routes "sraraa/reciever_src/routes/auth/login"
	forgot_password_routes "sraraa/reciever_src/routes/auth/password"
	signup_routes "sraraa/reciever_src/routes/auth/signup"
	"sraraa/reciever_src/routes/auth/signup/onboarding_routes"
	verify_session_routes "sraraa/reciever_src/routes/auth/verify_session"
	user_assets_routes "sraraa/reciever_src/routes/main/user"
	user_info_sender_routes "sraraa/sender_src/sender_routes/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// ✅ IMPORTANT: Initialize ALL tables, not just the connection
	_, err := db.InitializeDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database with tables:", err)
	}

	// Get the database connection
	dbConn := db.GetDB()
	if dbConn == nil {
		log.Fatal("Failed to get DB connection")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ginRouter := gin.New()
	ginRouter.Use(gin.Logger(), gin.Recovery())

	signup_routes.RegisterSignupRoutes()
	onboarding_routes.RegisterOnboardingRoutes()
	login_routes.LoginRoutes()

	access_auth_routes.RegisterAccessAuthRoutes(dbConn)
	verify_session_routes.VerifySessionRoutes()
	user_info_sender_routes.RegisterUserSenderRoutes()
	user_assets_routes.RegisterUserAssetsRoutes(ginRouter)
	forgot_password_routes.RegisterForgotPasswordRoutes()

	http.Handle("/", ginRouter)

	coreHandler := cors.EnableCORS(http.DefaultServeMux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: coreHandler,
	}

	go func() {
		for {
			access_auth_controller.AutoDeleteUnverifiedUsers(dbConn)
			time.Sleep(1 * time.Hour)
		}
	}()

	go func() {
		fmt.Printf("Server running on port %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Close database connection
	db.CloseDB()

	fmt.Println("Server exiting")
}
