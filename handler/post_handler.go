package handler

import (
	"blog1/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type AddPost struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type UpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ResponseJSON sends the reponse to the client.
func ResponseJSON(c *gin.Context, httpCode, errCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
}

// HandleAddPost is a handler on the add post api.
func HandleAddPost(c *gin.Context) {
	var addPost AddPost

	if err := c.BindJSON(&addPost); err != nil {
		ResponseJSON(c, http.StatusBadRequest, 400, "not able to parse the input data", nil)
		return
	}

	post, err := database.AddPost(addPost.Title, addPost.Content)
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return
	}

	response := AddPost{
		Id:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.String(),
	}

	ResponseJSON(c, http.StatusCreated, 201, "Successfully Created", response)

}

// HandleGetAllPosts is a handler on the getting All posts api.
func HandleGetAllPosts(c *gin.Context) {
	posts, err := database.GetAllPosts()
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
	}

	var response []AddPost

	for _, post := range posts {
		response = append(response, AddPost{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.String(),
		})
	}

	ResponseJSON(c, http.StatusOK, 200, "Successfully Fetched", response)

}

// HandleGetPostByID is a handler on the getting a post by id api.
func HandleGetPostByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ResponseJSON(c, http.StatusBadRequest, 400, err.Error(), nil)
		return
	}

	post, err := database.GetPostByID(id)
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return

	}

	response := AddPost{
		Id:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.String(),
	}

	ResponseJSON(c, http.StatusOK, 200, "Successfully Fetched", response)
}

// HandleDeletePost is a handler on the deleting a post by id api.
func HandleDeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return
	}

	if err := database.DeletePost(id); err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return
	}

	ResponseJSON(c, http.StatusOK, 200, "Successfully Deleted", nil)

}

// HandleUpdatePost is a handler on the updating a post by id api.
func HandleUpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ResponseJSON(c, http.StatusBadRequest, 400, err.Error(), nil)
		return
	}

	_, err = database.GetPostByID(id)
	if err != nil {
		ResponseJSON(c, http.StatusNotFound, 404, err.Error(), nil)
		return
	}

	var updatePost UpdatePost

	if err := c.BindJSON(&updatePost); err != nil {
		ResponseJSON(c, http.StatusBadRequest, 400, "not able to parse the input data", nil)
		return
	}

	post, err := database.UpdatePost(id, updatePost.Title, updatePost.Content)
	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, 500, err.Error(), nil)
		return
	}

	response := AddPost{
		Id:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UpdatedAt: post.UpdatedAt.String(),
	}

	ResponseJSON(c, http.StatusOK, 200, "Successfully updated", response)

}
