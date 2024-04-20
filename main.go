package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apiwatdev/go-gin-mongo-crud/bootstraps"
	"github.com/apiwatdev/go-gin-mongo-crud/routers"
)

func main() {

	bootstraps.InitEnv()

	ctx := context.Background()
	err := bootstraps.InitMongo(bootstraps.GetEnv().MONGO_CONNECTION_STR, bootstraps.GetEnv().MONGO_DATABASE)
	if err != nil {
		log.Fatalf("Error initializing MongoDB: %s", err)
	}
	defer bootstraps.GetMongoClient().Disconnect(ctx)

	// orderRepository := repositories.NewOrderRepository(bootstraps.GetDatabase(), ctx)
	// orderItemRepository := repositories.NewOrderItemRepository(bootstraps.GetDatabase(), ctx)

	// result, err := orderRepository.InsertOne(order)
	// orderID := result.InsertedID.(primitive.ObjectID)

	// orderItems := []models.OrderItem{
	// 	{OrderID: orderID, Product: "Product A", Quantity: 2, Price: 20},
	// 	{OrderID: orderID, Product: "Product B", Quantity: 1, Price: 10},
	// }

	// _, err = orderItemRepository.InsertMany(orderItems)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create a new Gin router
	router := routers.InitRouter(ctx)

	// Run the HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
