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

- **User** — system users with roles and authentication data
- **Warehouse_location** — physical address data for warehouse locations
- **Warehouse** — warehouses linked to locations and users
- **Product** — product catalog with SKU, pricing, and VAT rates
- **Stock_balance** — current stock quantities per warehouse and product
- **Partner** — business partners with types and tax information
- **Contact_info** — partner contact details
- **Document** — inventory documents like orders, invoices, etc.
- **Documentline** — individual lines/items within a document

## Setup and Installation

### Prerequisites

- Go 1.20+ installed  
- PostgreSQL database  
- Linux or Windows Subsystem for Linux (WSL) recommended for running the setup script  

### Setup

1. Clone this repository:

   git clone https://github.com/yourusername/inventory-management.git
   cd inventory-management

2. Run the setup script to initialize the database and environment:

    ./scripts/setup.sh

Note: The setup script is designed to run in a Linux or WSL environment.

3. Configure your environment variables (e.g., database credentials) as required.

4. Run or build and run the Go API:

    go build -o inventory-api ./cmd/inventory-api
    ./inventory-api

## Usage

The backend exposes RESTful endpoints for managing users, warehouses, products, stock, partners, and documents.
API documentation and example requests will be provided soon.

## Project Structure

    /cmd           # Main application entrypoints
    /internal      # Core business logic and services
    /scripts       # Setup and utility scripts (e.g., setup.sh)
    /db            # Database initialization and migrations