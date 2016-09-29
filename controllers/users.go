package controllers

import (
	"github.com/kataras/iris"
	"github.com/itsbalamurali/parse-server/database"
)

type UserAPI struct {
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
func (u UserAPI) Signup(ctx *iris.Context)  {

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
func (u UserAPI) Login(ctx *iris.Context) {

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
func (u UserAPI) Logout(ctx *iris.Context) {


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
func (u UserAPI) Get(ctx *iris.Context) {

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
func (u UserAPI) Me(ctx *iris.Context) {

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
func (u UserAPI) Update(ctx *iris.Context)  {


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
func (u *UserAPI) GetAll(ctx *iris.Context) {

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
func (u *UserAPI) Delete(ctx *iris.Context){

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
func (u *UserAPI) ResetPassword(ctx *iris.Context) {

	Db := database.MgoDb{}
	Db.Init()
	Db.Close()
}

/*
func (u *UserAPI) IsAuthenticated() bool {
}

func (u *UserAPI) GetSessionToken() string {
}

func (u *UserAPI) setSessionToken(token string) {
}

func (u *UserAPI) clearSession() {
}
*/