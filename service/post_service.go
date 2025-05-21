package service

import (
	"cloud-sek/cache"
	"cloud-sek/constants"
	"cloud-sek/database"
	"cloud-sek/models"
	"fmt"
)

type PostService struct {
	Repo *database.PostRepository
}

func NewPostService(repo *database.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (ps *PostService) CreatePost(post models.Post) error {
	if err := ps.Repo.InsertPost(constants.INSERT_POST, post); err != nil {
		fmt.Print("Error inserting post into database: ", err)
		return err
	}
	cache.SetPostCacheById(post.ID, post)
	return nil
}

func (ps *PostService) GetPostById(postId string) (*models.Post, error) {
	cachedPost := cache.GetCouponCacheById(postId)
	if cachedPost != nil {
		return cachedPost, nil
	}

	post, err := ps.Repo.GetPostById(constants.GET_POST_BY_ID, postId)
	if err != nil {
		return nil, err
	}

	cache.SetPostCacheById(post.ID, post)
	return &post, nil
}

func (ps *PostService) GetCommentsByPostID(postId string) ([]models.Comment, error) {
	comments, err := ps.Repo.GetCommentsByPostID(constants.GET_COMMENTS_BY_POST_ID, postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (ps *PostService) CreateComment(comm models.Comment) error {
	_, err := ps.Repo.InsertComment(constants.INSERT_COMMENTS, comm)
	if err != nil {
		return err
	}
	return nil
}
