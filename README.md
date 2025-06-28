# Inventory

A fullstack inventory management application with a Go backend API and a Next.js (Material UI) frontend.  
Easily add, view, update, and delete products, with in-memory storage and a modern UI.

## Project Structure

```
Inventory/
├── Inventoryinventorydashboardfrontend/   # Next.js frontend (Material UI, Ant Design icons)
│   ├── src/
│   │   └── app/
│   │       ├── layout.js
│   │       └── page.js
│   ├── components/
│   │   ├── AddProduct.js
│   │   └── ProductList.js
│   ├── lib/
│   │   └── api.js
│   ├── package.json
│   ├── Dockerfile
│   └── ... (other frontend files)
├── simpleInvetoryApiBackend/              # Go backend API
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
├── docker-compose.yml
└── README.md
```

## Features

- **Backend (Go):**

  - RESTful API: `GET /products`, `POST /product`, `PUT /product/:id`, `DELETE /product/:id`
  - In-memory product storage (no database)
  - Filtering by name and sorting by quantity or name (via query params)
  - Simple, readable code

- **Frontend (Next.js + Material UI):**

  - Add, view, edit, and delete products
  - Filter and sort products
  - Form validation (no empty name, positive quantity)
  - Responsive, modern design with Material UI and Ant Design icons
  - No custom CSS required

- **DevOps:**
  - Dockerfiles for both backend and frontend
  - `docker-compose.yml` for easy fullstack orchestration

## Getting Started

### 1. Prerequisites

- [Docker](https://www.docker.com/) and Docker Compose installed
- (For manual run) Node.js 20+ and Go 1.24.4 installed

### 2. Running with Docker Compose (Recommended)

From the root `Inventory/` directory:

```bash
docker-compose up --build
```

- **Frontend:** [http://localhost:2222](http://localhost:2222)
- **Backend API:** [http://localhost:1111/products](http://localhost:1111/products)

To stop:

```bash
docker-compose down
```

### 3. Running Locally (Without Docker)

**Backend:**

```bash
cd simpleInvetoryApiBackend
go run main.go
```

**Frontend:**

```bash
cd Inventoryinventorydashboardfrontend
npm install
npm run dev
```

- Open [http://localhost:2222](http://localhost:2222)

## API Endpoints

| Method | Endpoint       | Description                   |
| ------ | -------------- | ----------------------------- |
| GET    | `/products`    | List (optionally filter/sort) |
| POST   | `/product`     | Add a new product             |
| PUT    | `/product/:id` | Update a product              |
| DELETE | `/product/:id` | Delete a product              |

**Filtering/Sorting:**

- `/products?name=apple&sort=quantity`
- `/products?sort=name`
- `/products?sort=quantity`

## Frontend Usage

- **Add Product:** Fill the form and click "Add"
- **Edit Product:** Click the edit icon, modify, then save
- **Delete Product:** Click the delete icon
- **Filter/Sort:** Use the filter input and sort dropdown above the product list

## Design & Decisions

- **Clarity:** Each component/file has a clear, single responsibility
- **Simplicity:** In-memory storage and minimal dependencies for demo purposes
- **Correctness:** All CRUD, filter, and sort operations are covered and tested
- **Styling:** 100% Material UI (no custom CSS), Ant Design icons for actions
- **Communication:** Code and structure are easy to follow and explain

## Customization

- To persist data, replace the in-memory Go slice with a database.
- To deploy, use the provided Docker setup on any cloud or server.

## License

MIT License

Copyright 2025 Yashrajsinhdev

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
