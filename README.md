# Ecommerce_ste_OCG_final_project
Ecommerce_ste_OCG_final_project

This project include develop front end and backend(include server and database) for an ecommerce site


To reproduce the database, use your mysql CLI to run the sql files in path./database/database_backup in this order

ecommerce_category.sql

ecommerce_user.sql

ecommerce_product.sql

ecommerce_cart.sql

The backend include 2 subrouter, one for general access and one for personal access that require authorization.

Backend is developed in Golang divided in 3 package.

api package serve HTTP request for different routes

todb package serve the interaction between server and database

main package to launch server

Front end is still in development

This project is still on the way...

Thank you
