package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WeatherData struct {
	ID      int    `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	State   string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Data    Data   `json:"data" bson:"data"`
}

type Data struct {
	Temperature   int    `json:"temperature" bson:"temperature"`
	WindDirection string `json:"wind_direction" bson:"wind_direction"`
	WindVelocity  int    `json:"wind_velocity" bson:"wind_velocity"`
	Humidity      int    `json:"humidity" bson:"humidity"`
	Condition     string `json:"condition" bson:"condition"`
	Pressure      int    `json:"pressure" bson:"pressure"`
	Icon          string `json:"icon" bson:"icon"`
	Sensation     int    `json:"sensation" bson:"sensation"`
	Date          string `json:"date" bson:"date"`
}

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27018"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("weatherdb").Collection("weatherdata")

	var apiUrl = "http://apiadvisor.climatempo.com.br/api/v1/weather/locale/3477/current?token=89265a6623296fdf9a4cabf5ba0233a4"
	resp, _ := http.Get(apiUrl)
	body, _ := io.ReadAll(resp.Body)

	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	_, err = collection.InsertOne(context.Background(), weather)
	if err != nil {
		log.Fatal("Error inserting data into MongoDB:", err)
	}

	fmt.Println("Data inserted successfully.")
}
