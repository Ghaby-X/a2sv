package data

import (
	"context"
	"fmt"

	"github.com/Ghaby-X/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo" // Import mongo package
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the user's password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares provided password with hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsDatabaseEmpty checks if the users collection is empty
func IsDatabaseEmpty(usersCollection *mongo.Collection) (bool, error) {
	count, err := usersCollection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		return false, fmt.Errorf("failed to count documents: %w", err)
	}
	return count == 0, nil
}

// CreateUser creates a new user in the database
func CreateUser(user models.User, usersCollection *mongo.Collection) error {
	isDbEmpty, err := IsDatabaseEmpty(usersCollection)
	if err != nil {
		return err
	}

	if isDbEmpty {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()

	_, err = usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string, usersCollection *mongo.Collection) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

// PromoteUser promotes a user to admin role
func PromoteUser(id primitive.ObjectID, usersCollection *mongo.Collection) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to promote user: %w", err)
	}
	return nil
}