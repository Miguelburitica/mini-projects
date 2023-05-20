package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Miguelburitica/go-rest-server/helpers"
	"github.com/Miguelburitica/go-rest-server/models"
	"github.com/Miguelburitica/go-rest-server/repository"
	"github.com/Miguelburitica/go-rest-server/server"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type UpsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id  				string `json:"id"`
	PostContent string `json:"post_content"`
}

type PostUpdateResponse struct {
	Message string `json:"message"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if claims, ok, err := helpers.SecurityBasic(s, w, r); ok {
			var postRequest = UpsertPostRequest{}

			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			post := models.Post{
				Id: 				 id.String(),
				PostContent: postRequest.PostContent,
				UserId: 		 claims.UserId,
			}

			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var postMessage = models.WebsocketMessage{
				Type:   "Post_Created",
				Payload: post,
			}
			s.Hub().Broadcast(postMessage, nil)
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				Id: 				 post.Id,
				PostContent: post.PostContent,
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok, err := helpers.SecurityBasic(s, w, r); ok {
			params := mux.Vars(r)
			fmt.Println("params query id: ", params)
			post, err := repository.GetPostById(r.Context(), params["id"])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println("post: ", post)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if claims, ok, err := helpers.SecurityBasic(s, w, r); ok {
			params := mux.Vars(r)
			var postRequest = UpsertPostRequest{}

			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			post, err := repository.GetPostById(r.Context(), params["id"])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if post.UserId != claims.UserId {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			post.PostContent = postRequest.PostContent

			err = repository.UpdatePost(r.Context(), post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostUpdateResponse{
				Message: "Post updated successfully",
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if claims, ok, err := helpers.SecurityBasic(s, w, r); ok {
			params := mux.Vars(r)

			post, err := repository.GetPostById(r.Context(), params["id"])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if post.UserId != claims.UserId {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			err = repository.DeletePost(r.Context(), params["id"])
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostUpdateResponse{
				Message: "Post deleted successfully",
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func ListPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, ok, err := helpers.SecurityBasic(s, w, r); ok {
			var err error
			pageStr := r.URL.Query().Get("page")
			var page = uint64(0)

			if pageStr != "" {
				page, err = strconv.ParseUint(pageStr, 10, 64)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			posts, err := repository.ListPost(r.Context(), page)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}