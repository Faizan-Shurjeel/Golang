# Gommerce - Modern E-Commerce Platform

A modern e-commerce web application built with Go using the Gin framework. Gommerce features a clean, responsive design and provides a seamless shopping experience with product browsing, search functionality, and detailed product views.

## Features

- 🏠 **Homepage** with featured products showcase
- 📱 **Responsive Design** optimized for desktop and mobile
- 🔍 **Product Search** with real-time filtering
- 📦 **Product Catalog** with detailed product pages
- 🛒 **Product Details** with add-to-cart functionality
- 🎨 **Modern UI** with CSS Grid and Flexbox
- ⚡ **Fast Performance** with Go backend
- 🔗 **RESTful API** endpoints for product data

## Tech Stack

- **Backend**: Go 1.24.2 with Gin framework
- **Frontend**: HTML5, CSS3, Vanilla JavaScript
- **Styling**: Custom CSS with CSS Grid and Flexbox
- **HTTP Router**: Gin-Gonic
- **Template Engine**: Go HTML templates

## Project Structure

```
e:\Golang\
├── main.go              # Main application entry point
├── go.mod               # Go module dependencies
├── views/               # HTML templates
│   ├── home.html        # Homepage template
│   ├── products.html    # Products listing page
│   ├── product_details.html # Individual product page
│   └── 404.html         # Error page
├── public/              # Static assets
│   ├── style.css        # Main stylesheet
│   └── app.js           # Frontend JavaScript
└── assets/              # Product images and media
```

## Prerequisites

- Go 1.19 or higher
- Git

## Installation & Setup

1. **Clone the repository:**

```bash
git clone <repository-url>
cd Golang
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Run the application:**

```bash
go run main.go
```

4. **Access the application:**
   Open your browser and navigate to `http://localhost:3000`

## Available Routes

### Web Routes

- `GET /` - Homepage with featured products
- `GET /products` - Products listing page
- `GET /products/:id` - Individual product details page

### API Routes

- `GET /api/products` - JSON API for all products
- `GET /api/products/:id` - JSON API for specific product

### Static Routes

- `/public/*` - CSS and JavaScript files
- `/assets/*` - Product images and media

## Usage

### Browsing Products

1. Visit the homepage to see featured products
2. Click "Browse Products" or navigate to `/products`
3. Use the search bar to filter products by name or description
4. Click on any product card to view detailed information

### Product Details

- View high-resolution product images
- Read detailed product descriptions
- See pricing information
- Adjust quantity and add items to cart (UI only)

## Development

### Running in Development Mode

```bash
# Run with auto-reload (requires air)
air

# Or run directly
go run main.go
```

### Building for Production

```bash
# Build binary
go build -o gommerce main.go

# Run binary
./gommerce
```

### Code Structure

#### Backend (main.go)

- Product struct definition
- Mock product database
- Gin router setup
- HTML template rendering
- API endpoints for AJAX requests

#### Frontend (public/app.js)

- Product fetching from API
- Search functionality
- Dynamic content rendering
- Product detail page handling

#### Styling (public/style.css)

- CSS custom properties for theming
- Responsive grid layouts
- Modern card-based design
- Mobile-first approach

## API Documentation

### Get All Products

```http
GET /api/products
```

**Response:**

```json
[
  {
    "id": 1,
    "name": "Premium Headphones",
    "description": "High-quality sound experience...",
    "price": 199.99,
    "imageUrl": "assets/headphone.avif"
  }
]
```

### Get Product by ID

```http
GET /api/products/:id
```

**Response:**

```json
{
  "id": 1,
  "name": "Premium Headphones",
  "description": "High-quality sound experience...",
  "price": 199.99,
  "imageUrl": "assets/headphone.avif"
}
```

# eCommerce Backend API

A comprehensive eCommerce backend API built with Go, Gin framework, and MongoDB. This project implements user authentication, product management, and shopping cart functionality.

## Features

- **Product Management**: Full CRUD operations for products
- **User Authentication**: JWT-based authentication with bcrypt password hashing
- **Shopping Cart**: User-specific cart management
- **Search & Filtering**: Search products by name and filter by category
- **Pagination**: Efficient pagination for product listings
- **Error Handling**: Comprehensive error handling and validation

## Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin (HTTP web framework)
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **CORS**: gin-contrib/cors

## Project Structure

```
ecommerce-backend/
├── main.go                 # Application entry point
├── config/
│   └── database.go         # Database configuration
├── models/
│   ├── product.go          # Product model
│   ├── user.go            # User model
│   └── cart.go            # Cart model
├── controllers/
│   ├── product_controller.go   # Product CRUD operations
│   ├── auth_controller.go      # Authentication logic
│   └── cart_controller.go      # Cart operations
├── middleware/
│   └── auth_middleware.go      # JWT authentication middleware
├── routes/
│   └── routes.go              # Route definitions
├── scripts/
│   └── seed.go               # Database seeding script
├── .env                      # Environment variables
├── go.mod                    # Go module dependencies
└── README.md                 # Project documentation
```

## Installation & Setup

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd ecommerce-backend
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   Create a `.env` file with the following variables:

   ```
   MONGODB_URI=mongodb://localhost:27017
   DATABASE_NAME=ecommerce_db
   JWT_SECRET=your-super-secret-jwt-key-here
   PORT=8080
   ```

4. **Start MongoDB**
   Make sure MongoDB is running on your system or use MongoDB Atlas.

5. **Seed the database (optional)**

   ```bash
   go run scripts/seed.go
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## API Endpoints

### Authentication

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login user

### Products

- `GET /api/products` - Get all products (supports search, filtering, pagination)
- `GET /api/products/:id` - Get product by ID
- `POST /api/products` - Create new product (protected)
- `PUT /api/products/:id` - Update product (protected)
- `DELETE /api/products/:id` - Delete product (protected)

### Cart

- `POST /api/cart` - Add item to cart (protected)
- `GET /api/cart` - Get user's cart (protected)
- `DELETE /api/cart/:id` - Remove item from cart (protected)

### Query Parameters for Products

- `search` - Search products by name (case-insensitive)
- `category` - Filter products by category
- `page` - Page number for pagination (default: 1)
- `limit` - Number of items per page (default: 10)

Example: `GET /api/products?search=phone&category=Electronics&page=1&limit=5`

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

## Sample Requests

### Register User

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Get Products

```bash
curl -X GET "http://localhost:8080/api/products?page=1&limit=5"
```

### Create Product (Protected)

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "name": "New Product",
    "price": 99.99,
    "description": "Product description",
    "category": "Electronics",
    "stock": 10,
    "image": "https://example.com/image.jpg"
  }'
```

## Deployment

### Environment Variables for Production

```
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/
DATABASE_NAME=ecommerce_production
JWT_SECRET=your-production-secret-key
PORT=8080
```

### Build for Production

```bash
go build -o main .
./main
```

## Error Handling

The API includes comprehensive error handling:

- 400: Bad Request (validation errors)
- 401: Unauthorized (authentication required)
- 404: Not Found (resource not found)
- 409: Conflict (duplicate resources)
- 500: Internal Server Error

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

- **Project**: Gommerce E-Commerce Platform
- **Email**: support@gommerce.com
- **Demo**: http://localhost:3000

---

Built with ❤️ using Go and modern web technologies.
