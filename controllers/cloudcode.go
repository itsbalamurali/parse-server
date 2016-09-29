package controllers

import (
	"github.com/kataras/iris"
	"github.com/itsbalamurali/parse-server/database"
)

type FunctionAPI struct {
	*iris.Context
}

// @Title getOrdersByCustomer
// @Description retrieves orders for given customer defined by customer ID
// @Accept  json
// @Param   customer_id     path    int     true        "Customer ID"
// @Param   order_id        query   int     false        "Retrieve order with given ID only"
// @Param   order_nr        query   string  false        "Retrieve order with given number only"
// @Param   created_from    query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created starting from created_from"
// @Param   created_to      query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created before created_to"
// @Success 200 {array}  my_api.model.OrderRow
// @Failure 400 {object} my_api.ErrorResponse    "Customer ID must be specified"
// @Resource /order
// @Router /orders/by-customer/{customer_id} [get]
func (c *FunctionAPI) ExecuteFunc(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}

// @Title getOrdersByCustomer
// @Description retrieves orders for given customer defined by customer ID
// @Accept  json
// @Param   customer_id     path    int     true        "Customer ID"
// @Param   order_id        query   int     false        "Retrieve order with given ID only"
// @Param   order_nr        query   string  false        "Retrieve order with given number only"
// @Param   created_from    query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created starting from created_from"
// @Param   created_to      query   string  false        "Date-time string, MySQL format. If specified, API will retrieve orders that were created before created_to"
// @Success 200 {array}  my_api.model.OrderRow
// @Failure 400 {object} my_api.ErrorResponse    "Customer ID must be specified"
// @Resource /order
// @Router /orders/by-customer/{customer_id} [get]
func (c *FunctionAPI) TriggerJob(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()

	Db.Close()
}
