package data

import (
    "context"
    "go.mongodb.org/mongo-driver/x/bsonx"
    "time"

    log "github.com/sirupsen/logrus"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "portservice/data/models"
    "portservice/internal/grpc/domain"
)

const CollectionName string = "Airports"
const DbName string = "AirportService"

var (
    _ctx        context.Context
    _collection *mongo.Collection
)

type MongoDb struct {
    ConnString string
    client *mongo.Client
}

func newMongoDb() *MongoDb {
    return &MongoDb{}
}

func (db *MongoDb) Init(connStr string) {
    db.ConnString = connStr
    _ctx = context.Background()
}

func (db *MongoDb) OpenConnection() {
    client, err := mongo.NewClient(options.Client().ApplyURI(db.ConnString))
    if err != nil {
        log.Errorf("error on trying to create a new mongodb client: %v\n", err)
        return
    }

    db.client = client
    err = db.client.Connect(_ctx)
    if err != nil {
        log.Errorf("error on trying to connect to mongodb: %v\n", err)
    }

    err = db.client.Ping(_ctx, nil)
    if err != nil {
        log.Errorf("connection couldn't be established: %v\n", err)
    }

    initializeCollection(db)
}

func initializeCollection(db *MongoDb) {
    _collection = db.client.Database(DbName).Collection(CollectionName)
    _, _ = _collection.Indexes().CreateOne(
        context.Background(),
        mongo.IndexModel{
            Keys:    bsonx.Doc{{Key: "name", Value: bsonx.String("text")}},
            Options: options.Index().SetUnique(false),
        },
    )
}

func (db *MongoDb) CloseConnection() {
    err := db.client.Disconnect(_ctx)
    if err != nil {
        log.Errorf("error on trying to close mongodb connection: %v\n", err)
    }
}

//InsertAirport insert if new, update if is already present
func (db *MongoDb) InsertAirport(airport *domain.Airport) error {
    opts := options.Update().SetUpsert(true)
    filter := bson.M{"codename": bson.M{ "$eq":  airport.Codename}}
    update := bson.M{
     "$setOnInsert": bson.M{
         "name":         airport.Name,
         "city":         airport.City,
         "country":      airport.Country,
         "alias":        airport.Alias,
         "regions":      airport.Regions,
         "coordinates":  airport.Coordinates,
         "province":     airport.Province,
         "timezone":     airport.Timezone,
         "unlocs":       airport.Unlocs,
         "code":         airport.Code,
         "codename":     airport.Codename,
         "update_at":    time.Now(),
    }}

    _, err := _collection.UpdateOne(context.Background(), filter, update, opts)

    if err != nil {
        return err
    }

    return nil
}

func (db *MongoDb) SearchAirportByName(name string) (*[]*domain.Airport, error) {
    var results []*domain.Airport

    filter := bson.D{primitive.E{Key: "name", Value: name}}
    collection := db.client.Database(DbName).Collection(CollectionName)

    //err := collection.FindOne(_ctx, filter).Decode(&result)
    cursor, err := collection.Find(_ctx, filter)
    if err != nil {
        return nil, err
    }

    //Map result to slice
    for cursor.Next(_ctx) {
        t := models.AirportDB{}
        err := cursor.Decode(&t)
        if err != nil {
            return nil, err
        }
        results = append(results, mapAirportDBToDomainAirport(&t))
    }

    err = cursor.Close(_ctx)
    if err != nil {
        log.Errorf("error trying to close mongodb cursor: %v\n", err)
    }

    return &results, nil
}

func mapAirportDBToDomainAirport(airportdb *models.AirportDB) *domain.Airport {
    airport := domain.Airport {
        Name:        airportdb.Name,
        City:        airportdb.City,
        Country:     airportdb.Country,
        Alias:       airportdb.Alias,
        Regions:     airportdb.Regions,
        Coordinates: airportdb.Coordinates,
        Province:    airportdb.Province,
        Timezone:    airportdb.Timezone,
        Unlocs:      airportdb.Unlocs,
        Code:        airportdb.Code,
        Codename:    airportdb.Codename,
    }

    return &airport
}
