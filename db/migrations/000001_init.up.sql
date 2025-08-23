-- Tables
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE warehouse_locations (
    id SERIAL PRIMARY KEY,
    address_line1 VARCHAR(255) NOT NULL,
    address_line2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    country VARCHAR(2) NOT NULL,
    phone VARCHAR(50),
    email VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE warehouses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL UNIQUE,
    location_id INTEGER NOT NULL REFERENCES warehouse_locations(id) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_by_user_id INTEGER REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    sku VARCHAR(100) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    unit VARCHAR(50) NOT NULL,
    net_price NUMERIC(15, 2) NOT NULL,
    vat_rate NUMERIC(5, 3) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE stock_balances (
    id SERIAL PRIMARY KEY,
    warehouse_id INTEGER NOT NULL REFERENCES warehouses(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity NUMERIC(15, 2) NOT NULL DEFAULT 0,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (warehouse_id, product_id)
);

CREATE TABLE partners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    tax_number VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    partner_id INTEGER NOT NULL REFERENCES partners(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    address_line1 VARCHAR(255) NOT NULL,
    address_line2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    country VARCHAR(2) NOT NULL,
    phone VARCHAR(50),
    email VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    number VARCHAR(100) NOT NULL UNIQUE,
    date DATE NOT NULL,
    partner_id INTEGER NOT NULL REFERENCES partners(id) ON DELETE RESTRICT,
    created_by_user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    warehouse_id INTEGER REFERENCES warehouses(id) ON DELETE SET NULL,
    total_amount NUMERIC(15, 2) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    completed_at TIMESTAMP,
    related_document_id INTEGER REFERENCES documents(id) ON DELETE SET NULL
);

CREATE TABLE document_lines (
    id SERIAL PRIMARY KEY,
    document_id INTEGER NOT NULL REFERENCES documents(id) ON DELETE CASCADE,
    product_id INTEGER  NOT NULL REFERENCES products(id) ON DELETE RESTRICT,
    quantity NUMERIC(15, 2) NOT NULL,
    unit_price NUMERIC(15, 2) NOT NULL,
    total_price NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Default admin
INSERT INTO users (username, email, password, role) -- password is: admin.ADMIN.1234!
SELECT 'admin', 'admin@example.com', '$2a$12$Xn0TF.Pc6cT3WvagkECNFeVpwGCZKYDlhTo6VkuujyDKU/9Ury8Ia', 'ADMIN'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin');

-- Indexes
CREATE INDEX idx_warehouses_location_id ON warehouses(location_id);
CREATE INDEX idx_warehouses_created_by_user_id ON warehouses(created_by_user_id);

CREATE INDEX idx_stock_balances_warehouse_id ON stock_balances(warehouse_id);
CREATE INDEX idx_stock_balances_product_id ON stock_balances(product_id);

CREATE INDEX idx_contacts_partner_id ON contacts(partner_id);

CREATE INDEX idx_documents_partner_id ON documents(partner_id);
CREATE INDEX idx_documents_created_by_user_id ON documents(created_by_user_id);
CREATE INDEX idx_documents_warehouse_id ON documents(warehouse_id);
CREATE INDEX idx_documents_related_document_id ON documents(related_document_id);

CREATE INDEX idx_document_lines_document_id ON document_lines(document_id);
CREATE INDEX idx_document_lines_product_id ON document_lines(product_id);
