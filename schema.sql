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
