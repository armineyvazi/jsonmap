# JSONMap API

JSONMap API is a Go-based web service built using the Fiber framework. It leverages a GPT service to process and return laptop details based on provided input data. The application is equipped with Swagger for API documentation and is containerized using Docker for easy deployment.

## Getting Started

### Prerequisites

- **Go**: Install [Go](https://golang.org/doc/install).

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/armineyvazi/jsonmap.git
   cd jsonmap
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

3. Generate Swagger documentation:

   ```bash
   swag init -g cmd/api.go --output ./docs
   ```

### Running the Application

1. Run the application:

   ```bash
    go run cmd/api.go
   ```

2. The API will be available at `http://localhost:3005`.

3. Access the Swagger UI at:

   ```
   http://localhost:3005/swagger/index.html
   ```

## API Documentation

The API documentation is automatically generated using Swagger and can be accessed via the Swagger UI.

Visit `http://localhost:3000/swagger/index.html` to view the documentation, test endpoints, and see request/response schemas.

## Example Request

### POST /api/v1/gpt

Processes a request to generate laptop details based on the provided data.

**Request:**

```bash
curl --location 'http://127.0.0.1:3005/api/v1/gpt' \
--form 'data="\"Laptop: Dell Inspiron; Processor i7-10510U ; RAM 16GB; 512GB SSD Missing battery\"
\"MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed\"
\"ThinkPad, i5 CPU, 8GB memory, storage: 1TB HDD\"
\"Asus ROG, Processor: AMD Ryzen 7; RAM 16 GB; 1TB SSD; Damaged battrey\"
\"Dell Inspiron; Processor: i5-1135G7; RAM 8GB; Storage: 256.123548 SSD; Missing charger
"'
```

**Response:**

```json
[
  {
    "brand": "Dell",
    "model": "Inspiron",
    "processor": "Intel Core i7-10510U",
    "ramCapacity": "16GB",
    "ramType": "DDR4",
    "storageCapacity": "512GB",
    "batteryStatus": "No"
  }
]
```

