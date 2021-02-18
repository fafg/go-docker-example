package models

import (
    "time"
)

//Issue - struct to map with mongodb documents
type AirportDB struct {
    CreatedAt   time.Time          `bson:"created_at"`
    UpdatedAt   time.Time          `bson:"updated_at"`
    Name        string             `bson:"name"`
    City        string             `bson:"city"`
    Country     string             `bson:"country"`
    Alias       []string           `bson:"alias"`
    Regions     []string           `bson:"regions"`
    Coordinates []float64          `bson:"coordinates"`
    Province    string             `bson:"province"`
    Timezone    string             `bson:"timezone"`
    Unlocs      []string           `bson:"nlocs"`
    Code        string             `bson:"code"`
    Codename    string             `bson:"codename"`
}
