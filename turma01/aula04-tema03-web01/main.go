package main

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
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
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongo://localhost:27018"))
	defer client.Disconnect(context.Background())

	collection := client.Database("weatherdb").Collection("weatherdata")

	var body = []byte(`{
		"id": 3477,
		"name": "São Paulo",
		"state": "SP",
		"country": "BR  ",
		"data": {
			"temperature": 15,
			"wind_direction": "SSE",
			"wind_velocity": 20,
			"humidity": 88,
			"condition": "Chuvisco",
			"pressure": 1025,
			"icon": "3",
			"sensation": 15,
			"date": "2023-07-29 08:14:29"
		}
	}`)

	var weatherData WeatherData
	json.Unmarshal(body, &weatherData)

	collection.InsertOne(context.Background(), weatherData)
}
