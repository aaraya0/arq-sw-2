package repositories

import (
	"context"
	"fmt"

	"github.com/aaraya0/arq-software/arq-sw-2/dtos"
	model "github.com/aaraya0/arq-software/arq-sw-2/models"
	e "github.com/aaraya0/arq-software/arq-sw-2/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryMongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection string
}

func NewMongoDB(host string, port int, collection string) *RepositoryMongoDB {
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)).SetAuth(credential))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	titles, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Sprintf("[MongoDB] Available databases: %s", titles))

	return &RepositoryMongoDB{
		Client:     client,
		Database:   client.Database("publicaciones"),
		Collection: collection,
	}
}

func (repo *RepositoryMongoDB) Get(id string) (dtos.ItemDTO, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.ItemDTO{}, e.NewBadRequestApiError(fmt.Sprintf("error getting item %s invalid id", id))
	}
	result := repo.Database.Collection(repo.Collection).FindOne(context.TODO(), bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dtos.ItemDTO{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	var item model.Item
	if err := result.Decode(&item); err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}
	return dtos.ItemDTO{
		Id:          id,
		Titulo:      item.Titulo,
		Descripcion: item.Descripcion,
		Ciudad:      item.Ciudad,
		Estado:      item.Estado,
		Imagen:      item.Imagen,
		Vendedor:    item.Vendedor,
	}, nil
}

func (repo *RepositoryMongoDB) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	result, err := repo.Database.Collection(repo.Collection).InsertOne(context.TODO(), model.Item{
		Titulo: item.Titulo,
	})
	if err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
	}
	item.Id = fmt.Sprintf(result.InsertedID.(primitive.ObjectID).Hex())
	return item, nil
}
