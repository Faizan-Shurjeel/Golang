// Mock product data - this would normally come from the server
/* const products = [
    { 
        id: 1, 
        name: "Premium Headphones", 
        description: "High-quality sound experience with noise cancellation. Comfortable ear cups and long battery life make these perfect for long listening sessions.", 
        price: 199.99,
        image: "https://via.placeholder.com/800x600?text=Headphones"
    },
    { 
        id: 2, 
        name: "Smart Watch", 
        description: "Track your fitness and stay connected with this feature-rich smartwatch. Includes heart rate monitoring, GPS, and water resistance.", 
        price: 149.99,
        image: "https://via.placeholder.com/800x600?text=Smart+Watch" 
    },
    { 
        id: 3, 
        name: "Wireless Charger", 
        description: "Fast charging for all your devices. Compatible with most modern smartphones and capable of delivering up to 15W of power.", 
        price: 39.99,
        image: "https://via.placeholder.com/800x600?text=Wireless+Charger" 
    },
    { 
        id: 4, 
        name: "Bluetooth Speaker", 
        description: "Portable speaker with amazing sound quality. Water-resistant design makes it perfect for outdoor activities.", 
        price: 89.99,
        image: "https://via.placeholder.com/800x600?text=Bluetooth+Speaker" 
    },
    { 
        id: 5, 
        name: "Laptop Backpack", 
        description: "Stylish and functional backpack with padded compartment for laptops up to 15 inches. Multiple pockets for organization.", 
        price: 59.99,
        image: "https://via.placeholder.com/800x600?text=Laptop+Backpack" 
    },
    { 
        id: 6, 
        name: "Mechanical Keyboard", 
        description: "Premium typing experience with tactile mechanical switches. RGB backlighting and programmable macros.", 
        price: 129.99,
        image: "https://via.placeholder.com/800x600?text=Mechanical+Keyboard" 
    }
]; */ // Comment out or remove mock product data

// Handle products page
if (document.getElementById('products-container')) {
    const productsContainer = document.getElementById('products-container');
    const searchInput = document.getElementById('search-products');
    let allProducts = []; // To store products fetched from API
    
    // Display all products
    function displayProducts(productsToShow) {
        productsContainer.innerHTML = '';
        
        if (productsToShow.length === 0) {
            productsContainer.innerHTML = '<p>No products found matching your search.</p>';
            return;
        }
        
        productsToShow.forEach(product => {
            const productCard = document.createElement('div');
            productCard.className = 'product-card';
            
            productCard.innerHTML = `
                <img src="${product.imageUrl}" alt="${product.name}" class="product-img">
                <div class="product-info">
                    <h3 class="product-title">${product.name}</h3>
                    <p class="product-desc">${product.description.substring(0, 80)}...</p>
                    <div class="product-price">$${product.price.toFixed(2)}</div>
                </div>
            `;
            
            productCard.addEventListener('click', () => {
                window.location.href = `/products/${product.id}`;
            });
            
            productsContainer.appendChild(productCard);
        });
    }
    
    // Search functionality
    searchInput.addEventListener('input', (e) => {
        const searchTerm = e.target.value.toLowerCase();
        const filteredProducts = allProducts.filter(product => 
            product.name.toLowerCase().includes(searchTerm) || 
            product.description.toLowerCase().includes(searchTerm)
        );
        displayProducts(filteredProducts);
    });
    
    // Fetch products from API and initial display
    async function fetchAndDisplayProducts() {
        try {
            const response = await fetch('/api/products');
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            allProducts = await response.json();
            displayProducts(allProducts);
        } catch (error) {
            console.error("Failed to fetch products:", error);
            productsContainer.innerHTML = '<p>Error loading products. Please try again later.</p>';
        }
    }

    fetchAndDisplayProducts();
}

// Handle product details page
if (document.getElementById('product-details')) {
    const productDetailsContainer = document.getElementById('product-details');
    const urlParts = window.location.pathname.split('/');
    const productId = parseInt(urlParts[urlParts.length - 1]);
    
    async function fetchProductDetails() {
        try {
            const response = await fetch(`/api/products/${productId}`);
            if (!response.ok) {
                let errorMessage = 'Error loading product details.';
                if (response.status === 404) {
                    errorMessage = 'Product not found.';
                }
                productDetailsContainer.innerHTML = `<p>${errorMessage}</p>`;
                return;
            }
            const product = await response.json();
    
            productDetailsContainer.innerHTML = `
                <div>
                    <img src="${product.imageUrl}" alt="${product.name}" class="product-details-img">
                </div>
                <div class="product-details-info">
                    <h1>${product.name}</h1>
                    <div class="product-details-price">$${product.price.toFixed(2)}</div>
                    <p class="product-details-desc">${product.description}</p>
                    
                    <div class="add-to-cart">
                        <input type="number" class="quantity" value="1" min="1">
                        <button class="btn btn-primary">Add to Cart</button>
                    </div>
                    
                    <div style="margin-top: 2rem;">
                        <h3>Product Features</h3>
                        <ul style="margin-left: 1.5rem;">
                            <li>High-quality materials</li>
                            <li>1-year warranty</li>
                            <li>Free shipping on orders over $50</li>
                            <li>30-day money back guarantee</li>
                        </ul>
                    </div>
                </div>
            `;
            
            // Add event listener for Add to Cart button
            const addToCartBtn = productDetailsContainer.querySelector('.btn-primary');
            addToCartBtn.addEventListener('click', () => {
                const quantity = parseInt(productDetailsContainer.querySelector('.quantity').value);
                alert(`Added ${quantity} ${product.name}(s) to cart!`);
            });
        } catch (error) {
            console.error("Failed to fetch product details:", error);
            productDetailsContainer.innerHTML = '<p>Error loading product details. Please try again later.</p>';
        }
    }

    fetchProductDetails();
}
