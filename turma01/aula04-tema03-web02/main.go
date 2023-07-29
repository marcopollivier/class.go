package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
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
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27018"))
	defer client.Disconnect(context.Background())

	app := fiber.New()

	app.Get("/weather", func(c *fiber.Ctx) error {

		collection := client.Database("weatherdb").Collection("weatherdata")
		var weather WeatherData
		collection.FindOne(context.Background(), bson.M{}).Decode(&weather)

		return c.JSON(weather)
	})

	log.Fatal(app.Listen(":8080"))
}
