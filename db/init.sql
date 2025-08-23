CREATE TABLE "User" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE Warehouse_location (
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

CREATE TABLE Warehouse (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL UNIQUE,
    location_id INTEGER NOT NULL REFERENCES Warehouse_location(id) ON DELETE RESTRICT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    created_by_user_id INTEGER REFERENCES "User"(id) ON DELETE SET NULL
);

CREATE TABLE Product (
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

CREATE TABLE Stock_balance (
    id SERIAL PRIMARY KEY,
    warehouse_id INTEGER NOT NULL REFERENCES Warehouse(id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES Product(id) ON DELETE CASCADE,
    quantity NUMERIC(15, 2) NOT NULL DEFAULT 0,
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (warehouse_id, product_id)
);

CREATE TABLE Partner (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    tax_number VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE Contact_info (
    id SERIAL PRIMARY KEY,
    partner_id INTEGER NOT NULL REFERENCES Partner(id) ON DELETE CASCADE,
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

CREATE TABLE Document (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    number VARCHAR(100) NOT NULL UNIQUE,
    date DATE NOT NULL,
    partner_id INTEGER NOT NULL REFERENCES Partner(id) ON DELETE RESTRICT,
    created_by_user_id INTEGER REFERENCES "User"(id) ON DELETE SET NULL,
    warehouse_id INTEGER REFERENCES Warehouse(id) ON DELETE SET NULL,
    total_amount NUMERIC(15, 2) NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    completed_at TIMESTAMP,
    related_document_id INTEGER REFERENCES Document(id) ON DELETE SET NULL
);

CREATE TABLE Documentline (
    id SERIAL PRIMARY KEY,
    document_id INTEGER NOT NULL REFERENCES Document(id) ON DELETE CASCADE,
    product_id INTEGER  NOT NULL REFERENCES Product(id) ON DELETE RESTRICT,
    quantity NUMERIC(15, 2) NOT NULL,
    unit_price NUMERIC(15, 2) NOT NULL,
    total_price NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

-- Create default admin user (only if not exists) password is: admin.ADMIN.1234!
INSERT INTO "User" (username, email, password, role)
SELECT 'admin', 'admin@example.com', '$2a$12$Xn0TF.Pc6cT3WvagkECNFeVpwGCZKYDlhTo6VkuujyDKU/9Ury8Ia', 'ADMIN'
WHERE NOT EXISTS (SELECT 1 FROM "User" WHERE username = 'admin');