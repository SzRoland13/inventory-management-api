# Inventory Management Backend API

A Go-based backend API for managing inventory, warehouses, products, partners, and documents.  
This project provides a robust foundation for building an inventory management application.

---

## Table of Contents

- [Project Overview](#project-overview)  
- [Database Structure](#database-structure)  
- [Setup and Installation](#setup-and-installation)  
- [Usage](#usage)  
- [Project Structure](#project-structure)  
- [Contributing](#contributing)  
- [License](#license)

---

## Project Overview

This project implements a backend API for inventory management using Go (Golang).  
It handles core functionalities such as managing users, warehouses, products, stock balances, partners, contacts, documents, and document lines.  

The API is designed to be consumed by frontend applications or other services to track inventory and related operations efficiently.

---

## Database Structure

The backend uses a PostgreSQL database with the following core tables:

- **users** — system users with roles and authentication data
- **warehouse_locations** — physical address data for warehouse locations
- **warehouses** — warehouses linked to locations and users
- **products** — product catalog with SKU, pricing, and VAT rates
- **stock_balances** — current stock quantities per warehouse and product
- **partners** — business partners with types and tax information
- **contacts** — partner contact details
- **documents** — inventory documents like orders, invoices, etc.
- **document_lines** — individual lines/items within a document

## Setup and Installation

### Prerequisites

- Docker

### Setup

1. Clone this repository:

   git clone https://github.com/yourusername/inventory-management.git
   cd inventory-management

2. Start docker with:

    docker compose up -d

## Usage

The backend exposes RESTful endpoints for managing users, warehouses, products, stock, partners, and documents.
API documentation and example requests will be provided soon.

## Project Structure

    /cmd           # Main application entrypoints
    /internal      # Core business logic and services
    /scripts       # Setup and utility scripts (e.g., run-migrations.sh)
    /db            # Database initialization and migrations

## Contributing

Contributions are welcome!
Please open issues or pull requests for bug fixes, new features, or improvements.