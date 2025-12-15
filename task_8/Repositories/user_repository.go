
package Repositories

import (
	"context"
	"fmt"

	"github.com/Ghaby-X/task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	IsDatabaseEmpty() (bool, error)
	CreateUser(user Domain.User) error
	GetUserByUsername(username string) (Domain.User, error)
	PromoteUser(id primitive.ObjectID) error
}

type userRepository struct {
	usersCollection *mongo.Collection
}

func NewUserRepository(usersCollection *mongo.Collection) UserRepository {
	return &userRepository{usersCollection: usersCollection}
}

func (r *userRepository) IsDatabaseEmpty() (bool, error) {
	count, err := r.usersCollection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return false, fmt.Errorf("failed to count documents: %w", err)
	}
	return count == 0, nil
}

func (r *userRepository) CreateUser(user Domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userRepository) GetUserByUsername(username string) (Domain.User, error) {
	var user Domain.User
	filter := bson.M{"username": username}
	err := r.usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return Domain.User{}, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

func (r *userRepository) PromoteUser(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := r.usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to promote user: %w", err)
	}
	return nil
}
