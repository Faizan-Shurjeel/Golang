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

## Dependencies

- **gin-gonic/gin** - HTTP web framework
- **supabase-community/supabase-go** - Database integration (future use)
- Various supporting libraries for JSON, validation, etc.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Future Enhancements

- [ ] User authentication and registration
- [ ] Shopping cart persistence
- [ ] Payment integration
- [ ] Admin panel for product management
- [ ] Order tracking and history
- [ ] Product reviews and ratings
- [ ] Database integration (PostgreSQL/MySQL)
- [ ] Search filters (price, category, brand)

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

- **Project**: Gommerce E-Commerce Platform
- **Email**: support@gommerce.com
- **Demo**: http://localhost:3000

---

Built with ❤️ using Go and modern web technologies.
