# Delivery Order Price Calculator (DOPC)

A backend service that calculates the total price and breakdown of a delivery order, including the delivery fee, distance, and other charges.

---

## **Features**
- Calculate delivery price based on distance and cart value.
- Integrates with a mock Home Assignment API for venue data.
- Implements clean and testable architecture with mocked services for unit testing.

---

## **Requirements**

To run this project, ensure the following are installed:
- [Go](https://golang.org/doc/install) (Version 1.22 or higher recommended)
- A terminal or command-line interface
- `Git` for version control (optional but recommended)

---


## **Running the Service**

1. **Start the server**:
   - Run the following command in the project root directory:
     ```bash
     go run cmd/server/main.go
     ```

2. **Access the API**:
   - The service runs on port `8080` by default. You can access the endpoint at:
     ```
     GET http://localhost:8080/api/v1/delivery-order-price
     ```
   - Example usage:
     ```bash
     curl "http://localhost:8080/api/v1/delivery-order-price?venue_slug=test-venue&cart_value=1500&user_lat=52.5&user_lon=13.4"
     ```

---

## **Testing**

Run the unit tests to verify the functionality:

1. **Execute tests**:
   ```bash
   go test ./... -v
   ```

2. **Check test coverage**:
   ```bash
   go test ./... -cover
   ```

---

## **Project Structure**

```plaintext
.
├── cmd
│   └── server
│       └── main.go          # Entry point for the application
├── internal
│   ├── handlers             # HTTP handlers for API endpoints
│   ├── models               # Data structures used across the service
│   ├── services             # Core business logic and integrations
│   └── utils                # Utility functions like distance calculations
├── tests                    # Unit tests for handlers and services
├── go.mod                   # Go module dependencies
└── README.md                # Project documentation
```

---

## **API Endpoint**

### **GET `/api/v1/delivery-order-price`**

#### **Query Parameters**
| Parameter   | Type    | Description                                   | Required |
|-------------|---------|-----------------------------------------------|----------|
| `venue_slug`| String  | Unique identifier for the venue               | Yes      |
| `cart_value`| Integer | Total value of items in the shopping cart (in cents) | Yes      |
| `user_lat`  | Float   | Latitude of the user's location               | Yes      |
| `user_lon`  | Float   | Longitude of the user's location              | Yes      |

#### **Example Request**
```bash
curl "http://localhost:8080/api/v1/delivery-order-price?venue_slug=test-venue&cart_value=1500&user_lat=52.5&user_lon=13.4"
```

#### **Example Response**
```json
{
  "total_price": 1690,
  "small_order_surcharge": 0,
  "cart_value": 1500,
  "delivery": {
    "fee": 190,
    "distance": 80
  }
}
```

---

## **Notes**
- All monetary values are returned in the smallest unit of the local currency (e.g., cents for EUR).
- Ensure you have network access if the Home Assignment API endpoints are hosted remotely.

---
