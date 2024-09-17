package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mokhlesurr031/sre-project/backend/domain"
	"github.com/mokhlesurr031/sre-project/backend/internal/utils"
)

type ResourceHandler struct {
	ResourceUseCase domain.ResourceUseCase
}

func NewHTTPHandler(r *chi.Mux, resourceUseCase domain.ResourceUseCase) {
	handler := &ResourceHandler{
		ResourceUseCase: resourceUseCase,
	}
	r.Route("/api/v1/resource", func(r chi.Router) {
		r.Post("/create", handler.Post)
		r.Get("/list", handler.Get)
	})
}

type ReqResource struct {
	domain.Resource
}

func (resource *ResourceHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := ReqResource{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		utils.SendErrorResponse(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	resources := domain.Resource(req.Resource)

	res, err := resource.ResourceUseCase.Post(ctx, &resources)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (resource *ResourceHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	ctx := r.Context()
	resourceList, err := resource.ResourceUseCase.Get(ctx)
	if err != nil {
		log.Println(err)
	}
	er := json.NewEncoder(w).Encode(resourceList)
	if err != nil {
		log.Println(er)
	}
}
