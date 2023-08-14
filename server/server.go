package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_golang/weather"
	"math/rand"
	"net"
	"time"
)

type WeatherServer struct {
	weather.UnimplementedWeatherServer
}

func (s *WeatherServer) GetWeather(ctx context.Context, req *weather.WeatherRequest) (*weather.WeatherResponse, error) {

	fmt.Println("GetWeather Request : ", req)

	temperature := rand.Float32() * 30.0
	humidity := rand.Float32() * 100.0

	conditions := []string{"Sunny", "Cloudy", "Rainy", "Snowy"}
	condition := conditions[rand.Intn(len(conditions))]

	resp := &weather.WeatherResponse{
		City:        req.City,
		Temperature: temperature,
		Humidity:    humidity,
		Condition:   condition,
	}
	return resp, nil
}

func (s *WeatherServer) StreamWeatherUpdates(req *weather.WeatherRequest, stream weather.Weather_StreamWeatherUpdatesServer) error {

	fmt.Println("StreamWeatherUpdates Request : ", req)

	for {
		temperature := rand.Float32() * 30.0
		humidity := rand.Float32() * 100.0

		conditions := []string{"Sunny", "Cloudy", "Rainy", "Snowy"}
		condition := conditions[rand.Intn(len(conditions))]

		resp := &weather.WeatherResponse{
			City:        req.City,
			Temperature: temperature,
			Humidity:    humidity,
			Condition:   condition,
		}

		if err := stream.Send(resp); err != nil {
			return err
		}

		time.Sleep(time.Second * 1) // update in 1 second
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	weather.RegisterWeatherServer(server, &WeatherServer{})

	fmt.Println("Weather server is listening on port 8080")
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
