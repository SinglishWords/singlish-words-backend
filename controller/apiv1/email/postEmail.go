package email

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singlishwords/controller/apiv1"
	"singlishwords/model"
	"singlishwords/service"
	"strings"
)

type paramPostEmail struct {
	Email         string `json:"email" db:"email"`
	WantLuckyDraw string `json:"wantLuckyDraw" db:"want_lucky_draw"`
	WantUpdate    string `json:"wantUpdate" db:"want_update"`
	TimeOnPages   []int  `json:"timeOnPages" db:"time_on_pages"`
}

func tfToYs(s string) string {
	if s == "true" {
		return "yes"
	} else if s == "false" {
		return "no"
	}
	return s
}

func (p *paramPostEmail) toEmail() *model.Email {
	p.WantUpdate = tfToYs(p.WantUpdate)
	p.WantLuckyDraw = tfToYs(p.WantLuckyDraw)

	return &model.Email{
		Email:         p.Email,
		WantLuckyDraw: p.WantLuckyDraw,
		WantUpdate:    p.WantUpdate,
		TimeOnPages:   strings.Join(strings.Fields(fmt.Sprint(p.TimeOnPages)), ", "),
	}
}

// PostEmail godoc
// @Summary Post an email
// @Tags Email
// @Produce json
// @Param answer body paramPostEmail true "the email with two options"
// @Success 201 {object} model.Email
// @Failure 500 {object} apiv1.ErrorResponse
// @Router /email [post]
func PostEmail(c *gin.Context) (apiv1.HttpStatus, interface{}) {
	var param paramPostEmail
	err := c.BindJSON(&param)

	if err != nil {
		return apiv1.StatusPostParamError, err
	}

	email := param.toEmail()
	err = service.AddEmail(email)

	if err != nil {
		return apiv1.StatusFail(err.Error()), nil
	}

	return apiv1.StatusCreated, email
}
