package repository

import (
    "context"
    "mytemplate/internal/model"
    "go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
    Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
    return &UserRepository{
        Collection: db.Collection("users"),
    }
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
    _, err := r.Collection.InsertOne(ctx, user)
    return err
}
