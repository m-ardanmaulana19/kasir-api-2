-- Create categories table first (parent)
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- Create products table with category_id (child)
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
);

-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    total_amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create transaction_details table
CREATE TABLE IF NOT EXISTS transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    subtotal INT NOT NULL
);

-- Insert sample categories
INSERT INTO categories (name, description) VALUES
('Minuman', 'Berbagai macam minuman'),
('Makanan', 'Berbagai macam makanan'),
('Snack', 'Makanan ringan');

-- Insert sample products with category_id
INSERT INTO products (name, price, stock, category_id) VALUES
('Teh Pucuk', 3500, 10, 1),
('Le Minerale 600ml', 3000, 24, 1),
('Milku', 5000, 15, 1),
('Indomie Goreng', 3500, 50, 2),
('Chitato', 8000, 30, 3);
