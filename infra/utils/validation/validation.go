package validation

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)


type CustomValidator struct {
  validator *validator.Validate
}

func NewCustomValidator(v *validator.Validate) *CustomValidator {
  return &CustomValidator{v}
}


func (cv *CustomValidator) Validate(i interface{}) error {
    
  if err := cv.validator.Struct(i); err != nil {
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }

  return nil
}
