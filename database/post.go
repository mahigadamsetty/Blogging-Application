package database

import (
	"blog1/ent"
	"context"
	"fmt"
	"time"
)

// AddPost creates post with the given tile and content.
func AddPost(title string, content string) (*ent.Post, error) {
	post, err := EntClient.Post.Create().
		SetTitle(title).
		SetContent(content).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Println(post.Title)

	return post, nil
}

// GetPostByID gets the post given by id.
func GetPostByID(id int) (*ent.Post, error) {
	post, err := EntClient.Post.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetAllPosts gets all the posts.
func GetAllPosts() ([]*ent.Post, error) {

	posts, err := EntClient.Post.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// DeletePost deletes post with the given id.
func DeletePost(id int) error {
	if err := EntClient.Post.DeleteOneID(id).Exec(context.Background()); err != nil {
		return err
	}

	return nil
}

// UpdatePost updates post with the given id title and content.
func UpdatePost(id int, title string, content string) (*ent.Post, error) {
	post, err := EntClient.Post.UpdateOneID(id).
		SetTitle(title).
		SetContent(content).
		SetUpdatedAt(time.Now()).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return post, err
}
