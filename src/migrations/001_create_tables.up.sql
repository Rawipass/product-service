-- สร้างตาราง users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- สร้างตาราง products
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    gender VARCHAR(50) NOT NULL, 
    style VARCHAR(50) NOT NULL,
    size VARCHAR(10) NOT NULL, 
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- สร้างตาราง orders
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL, 
    address TEXT NOT NULL,
    status VARCHAR(50) NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- สร้างตาราง order_items
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL, 
    product_id INTEGER NOT NULL, 
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
