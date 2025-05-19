package order

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"tkh/logger"
	"tkh/models"
	"tkh/storage"
	validate "tkh/validator"
)

type CreateRes struct {
	models.OrderReq
	Id       string           `json:"id"`
	Products []models.Product `json:"products"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	o := models.OrderReq{}
	verr := validate.Validate(ctx, r.Body, &o)
	if verr != nil {
		logger.Println("error while validating the order req", verr.Error())
		verr.Write(w)
		return
	}
	st := storage.GetStorage(ctx)
	p := models.Products{
		IDs: make([]string, 0),
	}
	for _, val := range o.Items {
		p.IDs = append(p.IDs, val.ProductId)
	}
	err := st.Get(ctx, &p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}

	byt, err := json.Marshal(CreateRes{OrderReq: o, Id: uuid.New().String(), Products: p.Products})
	if err != nil {
		logger.Println("error while marshalling the order res", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(byt)
}
