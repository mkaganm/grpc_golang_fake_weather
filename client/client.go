package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_golang/weather"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)

	client := weather.NewWeatherClient(conn)

	req := &weather.WeatherRequest{
		City: "Istanbul",
	}

	// GetWeather
	res, err := client.GetWeather(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetWeather RPC: %v", err)
	}

	fmt.Printf("\n\n-*-*- GetWeather Response -*-*- \n\n")
	fmt.Printf("Weather in %s:\n", res.City)
	fmt.Printf("Temperature: %.2f°C\n", res.Temperature)
	fmt.Printf("Humidity: %.2f%%\n", res.Humidity)
	fmt.Printf("Condition: %s\n", res.Condition)

	fmt.Printf("\n--------------------------------------------\n")

	// StreamWeatherUpdates
	stream, err := client.StreamWeatherUpdates(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while opening stream: %v", err)
	}

	fmt.Printf("\n\n-*-*- Streaming Weather Updates -*-*-\n\n")
	for {
		update, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while receiving update: %v", err)
			break
		}

		fmt.Printf("Weather in %s:\n", update.City)
		fmt.Printf("Temperature: %.2f°C\n", update.Temperature)
		fmt.Printf("Humidity: %.2f%%\n", update.Humidity)
		fmt.Printf("Condition: %s\n", update.Condition)

		fmt.Printf("\n--------------------------------------------\n")
	}
}
