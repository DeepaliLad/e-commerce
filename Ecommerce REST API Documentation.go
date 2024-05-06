Ecommerce REST API Documentation
Authentication
This API employs token-based authentication using JSON Web Tokens (JWT). To access protected routes, clients must include a valid JWT in the Authorization header of their requests.

Token Generation
Authentication endpoints for token generation are not provided in this API. Users are expected to obtain tokens from an authentication service using their credentials.

Protected Routes
GET /protected_route
Requires a valid JWT in the request headers.
Returns a message indicating successful access to the protected route.
Example Request:
http
GET /protected_route
Authorization: Bearer <JWT>
Example Response:
json

{
    "message": "This route is protected by token-based authorization."
}
Public Routes
GET /
Returns a simple message indicating the existence of the Ecommerce REST API.
Example Request:
http

GET /
Example Response:
json

{
    "message": "Ecommerce REST API"
}
Categories
POST /category/create
Creates a new category.
Requires a valid JWT in the request headers.
Request Body:
json

{
    "name": "Category Name",
    "subcategories": [1, 2, 3]  // Optional: List of subcategory IDs
}
Response:
json

{
    "id": 1,
    "name": "Category Name",
    "subcategories": [1, 2, 3]
}
GET /categories
Retrieves all categories.
Requires a valid JWT in the request headers.
Response:
json

{
    "categories": [
        {
            "id": 1,
            "name": "Category 1",
            "subcategories": [1, 2, 3]
        },
        {
            "id": 2,
            "name": "Category 2",
            "subcategories": [4, 5]
        }
    ]
}
GET /category/<int:c_id>
Retrieves a specific category by ID.
Requires a valid JWT in the request headers.
Response:
json

{
    "id": 1,
    "name": "Category Name",
    "subcategories": [1, 2, 3]
}
PUT /category/<int:c_id>/update
Updates a category by ID.
Requires a valid JWT in the request headers.
Request Body:
json

{
    "name": "Updated Category Name",
    "subcategories": [4, 5, 6]  // Optional: List of subcategory IDs
}
Response:
json

{
    "id": 1,
    "name": "Updated Category Name",
    "subcategories": [4, 5, 6]
}
DELETE /category/<int:c_id>
Deletes a category by ID.
Requires a valid JWT in the request headers.
Response:
json

{
    "result": true
}
Subcategories
Document Subcategories endpoints similarly to Categories endpoints.

Products
Document Products endpoints similarly to Categories endpoints.