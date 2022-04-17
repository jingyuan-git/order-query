package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/pkg/app"
	"server/pkg/e"
	"server/pkg/setting"
	"server/pkg/util"
	"server/service"
)

func GetOrders(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form service.Order
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	orderService := service.Order{
		SearchOrder:    form.SearchOrder,
		StartOrderDate: form.StartOrderDate,
		EndOrderDate:   form.EndOrderDate,
		PageNum:        util.GetPage(form.PageNum),    
		PageSize:       setting.ServerSetting.PageSize,
	}

	total, amount, err := orderService.CountAndTotalAmount()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	orders, err := orderService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = orders
	data["total"] = total
	data["totalAmount"] = amount

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
