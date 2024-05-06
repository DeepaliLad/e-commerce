the Ecommerce REST API:

markdown

# Ecommerce REST API

This RESTful HTTP API built with Python Flask enables users to efficiently manage their ecommerce platform. Users can perform CRUD operations on products, categories, and subcategories. Categories can have multiple subcategories, and a subcategory can belong to multiple categories. Products can belong to multiple categories and subcategories.

Fetching product details also retrieves the categories and subcategories it belongs to. Additionally, the API provides search functionality for products by name, category, and subcategories. Pagination is implemented when fetching products by categories or subcategories, ensuring efficient data retrieval.

## Requirements

```bash
pip install -r requirements.txt
Usage
Start the server:

bash

python app.py
This command starts the server on 127.0.0.1:5000.

The project includes a preloaded dummy SQLite database located in the instance directory. To start with an empty database, delete the instance directory and restart the server.

To test the API using Postman, install the Postman agent on your OS and make requests to the API.

Endpoints
Fetch Products
GET /product/<name: string> - Get product with the specified name.
GET /subcategory/<subcategory_id: int>/products - Get products within a subcategory. Returns the first page of paginated results.
GET /subcategory/<subcategory_id: int>/products?page=<page_no> - Get products within a subcategory. Returns the specified page of paginated results.
GET /category/<category_id: int>/products - Get products within a category. Returns the first page of paginated results.
GET /category/<category_id: int>/products?page=<page_no> - Get products within a category. Returns the specified page of paginated results.
Category
GET /categories - Get all categories.
GET /category/(int: category_id) - Get category by ID.
DELETE /category/(int: category_id) - Delete category by ID.
POST /category/create - Create a new category.
json

{
  "name": "category_name",
  "subcategories": [subcat_id1, subcat_id2]  // optional
}
PUT /category/(int: category_id)/update - Update category by ID.
json

{
  "name": "category_name",
  "subcategories": [subcat_id1, subcat_id2]  // optional
}
Subcategory
GET /subcategories - Get all subcategories.
GET /subcategory/(int: subcategory_id) - Get subcategory by ID.
DELETE /subcategory/(int: subcategory_id) - Delete subcategory by ID.
POST /subcategory/create - Create a new subcategory.
json

{
  "name": "subcategory_name",
  "categories": [cat_id1, cat_id2],  // optional
  "products": [product_id1, product_id2]  // optional
}
PUT /subcategory/(int: subcategory_id)/update - Update subcategory by ID.
json

{
  "name": "subcategory_name",
  "categories": [cat_id1, cat_id2],  // optional
  "products": [product_id1, product_id2]  // optional
}
Product
GET /products - Get all products.
GET /product/(int: product_id) - Get product by ID.
DELETE /product/(int: product_id) - Delete product by ID.
POST /product/create - Create a new product.
json

{
  "name": "product_name",
  "description": "product_description",
  "subcategories": [subcat_id1, subcat_id2]  // optional
}
PUT /product/(int: product_id)/update - Update product by ID.
json

{
  "name": "product_name",
  "description": "product_description",
  "subcategories": [subcat_id1, subcat_id2]  // optional
}