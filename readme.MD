# RESTAURANT MANAGEMENT API With GO

## Introduction

This is an advanced restaurant management project built with Go and mongo database. It allows a user to create a menu which has a name, category and an id. We could have different menu categories such as the sliders menu, smoothies menu etc.Also, each menu has multiple food items and the food item consist of a food name, food price and a food id. Then there is a table, where an order is placed which also consist of order items which are basically the food items ie the food quantity and the food price. Finally we have the invoice which is created for each order which has a payment method and a payment status. We also have authentication with JWT.

##### Do the stuff below to initialize your project

1. Clone the repository
   `git clone https://github.com/Shittu24/restaurant-management-api-go.git`
2. Change the directory
   `cd restaurant-management-api-go`
3. Get all the dependencies
   `go mod tidy`
4. Run your project
   `go run main.go`

   Then go to postman and start running the routes
