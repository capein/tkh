package product

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"tkh/models"
	"tkh/storage"
)

func List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	st := storage.GetStorage(ctx)
	p := models.Products{}
	err := st.Get(ctx, &p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}
	s := models.SuccessResponse{
		Code: http.StatusOK,
		Data: p.Products,
	}
	s.Write(w)
}

func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	param := mux.Vars(r)
	st := storage.GetStorage(ctx)
	p := models.Product{Id: param["id"]}
	err := st.Get(ctx, &p)
	if err != nil && !errors.Is(err, models.ErrNoData) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}
	if errors.Is(err, models.ErrNoData) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s := models.SuccessResponse{
		Code: http.StatusOK,
		Data: p,
	}
	s.Write(w)
}
