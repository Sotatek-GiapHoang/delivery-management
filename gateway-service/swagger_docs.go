package main

// @Summary Login user
// @Description Authenticate user and return token
// @Tags users
// @Accept json
// @Produce json
// @Param user body object true "Login information"
// @Success 200 {object} object
// @Router /users/login [post]
func swaggerLoginUser() {}

// @Summary Create a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body object true "User Registration Info"
// @Success 201 {object} object
// @Failure 400 {object} object
// @Router /users/register [post]
func swaggerRegisterUser() {}

// @Summary Get user information
// @Description Get detailed information of the current user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object
// @Router /users [get]
func swaggerGetUserInfo() {}

// @Summary Get deliveries by user
// @Description Get the list of deliveries for the current user
// @Tags deliveries
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} object
// @Router /deliveries [get]
func swaggerGetUserDeliveries() {}

// @Summary Get delivery information
// @Description Get detailed information of a delivery
// @Tags deliveries
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the delivery"
// @Success 200 {object} object
// @Router /deliveries/{id} [get]
func swaggerGetDelivery() {}

// @Summary Update delivery status
// @Description Update the status of a delivery
// @Tags deliveries
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the delivery"
// @Param status body object true "New status"
// @Success 200 {object} object
// @Router /deliveries/{id} [put]
func swaggerUpdateDeliveryStatus() {}

// @Summary Create an order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param order body object true "Order information"
// @Success 201 {object} object
// @Router /orders [post]
func swaggerCreateOrder() {}

// @Summary Get user orders
// @Description Get the list of orders for the current user
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} object
// @Router /orders [get]
func swaggerGetUserOrders() {}

// @Summary Get order information
// @Description Get detailed information of an order
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the order"
// @Success 200 {object} object
// @Router /orders/{id} [get]
func swaggerGetOrder() {}

// @Summary Update order status
// @Description Update the status of an order
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID of the order"
// @Param status body object true "New status"
// @Success 200 {object} object
// @Router /orders/{id} [put]
func swaggerUpdateOrderStatus() {}
