package validate

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"tkh/models"
)

const (
	tagName        = "validate"
	tagValRequired = "required"
	tagValGreater  = "greater"
	tagValLesser   = "lesser"
)

type ValidatorFn func(v reflect.Value, tagdata string) (bool, string)

var validatorFunctions = map[string]ValidatorFn{}

func RegisterValidatorFn(name string, fn ValidatorFn) {
	if _, ok := validatorFunctions[name]; ok {
		panic(fmt.Sprintf("validator function %s already registered", name))
	}
	validatorFunctions[name] = fn
}

func init() {
	RegisterValidatorFn(tagValRequired, requireFn)
	RegisterValidatorFn(tagValGreater, greater)
	RegisterValidatorFn(tagValLesser, lesser)

}

func Validate(ctx context.Context, reader io.ReadCloser, obj models.IReqModel) *models.ErrorResponse {
	err := obj.Load(ctx, reader)
	if err != nil {
		fmt.Println(err)
		return &models.ErrorResponse{Code: http.StatusInternalServerError, Err: err}
	}
	//struct validation
	err = structValidation(ctx, obj)
	if err != nil {
		return &models.ErrorResponse{Code: http.StatusBadRequest, MessageType: "validation", Err: err}
	}
	err = obj.Validate(ctx)
	if err != nil {
		return &models.ErrorResponse{Code: http.StatusUnprocessableEntity, MessageType: "constraint", Err: err}
	}
	return nil
}

func structValidation(ctx context.Context, data interface{}) error {
	rt := reflect.TypeOf(data)
	ev := reflect.ValueOf(data)
	if rt.Kind() != reflect.Pointer {
		return errors.New("data must be a pointer")
	}

	return structv(ev.Elem())
}

func structv(ev reflect.Value) error {
	if ev.Kind() != reflect.Struct {
		return errors.New("data must be a struct")
	}
	rt := ev.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		if field.Type.Kind() == reflect.Slice {
			keyStr := field.Tag.Get(tagName)
			arr := strings.Split(keyStr, ",")
			for _, val := range arr {
				tagArr := strings.Split(val, "=")
				fn, ok := validatorFunctions[tagArr[0]]
				if !ok {
					continue
				}
				tagdata := ""
				if len(tagArr) > 1 {
					tagdata = tagArr[1]
				}
				ok, msg := fn(reflect.ValueOf(ev.Field(i).Len()), tagdata)
				if !ok {
					return errors.New(fmt.Sprintf("%s is %s", field.Tag.Get("json"), msg))
				}
			}

			for j := 0; j < ev.Field(i).Len(); j++ {
				err := structv(ev.Field(i).Index(j))
				if err != nil {
					return err
				}
			}
			continue
		}
		if field.Type.Kind() == reflect.Struct {
			err := structv(ev.Field(i))
			if err != nil {
				return err
			}
			continue
		}
		if field.Type.Kind() == reflect.Pointer && field.Type.Elem().Kind() == reflect.Struct {
			err := structv(ev.Field(i))
			if err != nil {
				return err
			}
			continue
		}

		keyStr := field.Tag.Get(tagName)
		arr := strings.Split(keyStr, ",")
		for _, val := range arr {
			tagArr := strings.Split(val, "=")
			fn, ok := validatorFunctions[tagArr[0]]
			if !ok {
				continue
			}
			tagdata := ""
			if len(tagArr) > 1 {
				tagdata = tagArr[1]
			}
			ok, msg := fn(ev.Field(i), tagdata)
			if !ok {
				return errors.New(fmt.Sprintf("%s is %s", field.Tag.Get("json"), msg))
			}
		}
	}
	return nil
}

func requireFn(v reflect.Value, tagData string) (bool, string) {
	msg := "required"
	if data := v.Interface(); data != nil {
		str, ok := data.(string)
		if !ok {
			return true, ""
		}
		if len(str) == 0 {
			return false, msg
		}

		return true, ""
	}
	return false, msg
}

func greater(v reflect.Value, tagData string) (bool, string) {
	msg := fmt.Sprintf("less than or equal %s", tagData)
	val, err := strconv.ParseInt(tagData, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false, msg
	}
	if v.CanInt() {
		if v.Int() > val {
			return true, ""
		}
	}
	return false, msg

}

func lesser(v reflect.Value, tagData string) (bool, string) {
	msg := fmt.Sprintf("greater than or equal %s", tagData)
	val, err := strconv.ParseInt(tagData, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false, msg
	}

	if v.CanInt() {
		if v.Int() < val {
			return true, ""
		}
	}
	return false, msg

}
