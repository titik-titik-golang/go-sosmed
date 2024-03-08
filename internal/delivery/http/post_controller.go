package http

import (
	"encoding/json"
	"net/http"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
	"github.com/guregu/null"
)

type PostController struct {
	PostUseCase *use_case.PostUseCase
}

func NewPostController(useCase *use_case.PostUseCase) *PostController {
	return &PostController{
		PostUseCase: useCase,
	}
}

func (p *PostController) Create(w http.ResponseWriter, r *http.Request) {
	var req model_request.CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.PostUseCase.Create(&req); err != nil {
		http.Error(w, "Failed to create new post", http.StatusInternalServerError)
		return
	}

	response.NewResponse(w, http.StatusText(http.StatusOK), new(string), http.StatusOK)
}

func (p *PostController) Get(w http.ResponseWriter, r *http.Request) {
	var req model_request.GetPostRequest
	postId := mux.Vars(r)["id"]

	req.PostId = null.NewString(postId, true)

	post, errGet := p.PostUseCase.Get(&req)

	if errGet != nil {
		http.Error(w, "Failed to get post", http.StatusInternalServerError)
		return
	}

	response.NewResponse(w, http.StatusText(http.StatusOK), post, http.StatusOK)
}
