# Cinema Seating Service

This service provides a RESTful API for managing cinema seating arrangements. It allows configuring the cinema layout, querying available seats, reserving seats, and canceling reservations while enforcing social distancing rules.

## Features

- Configure cinema layout with rows, columns, and minimum distance rules.
- Query available seats for a group while maintaining social distancing.
- Reserve specific seats by their coordinates.
- Cancel reservations for specific seats.

## Prerequisites

- Go 1.18 or later
- Postman (optional, for testing the API)

## Run the Service

To start the service, run the following command:

```bash
go run main.go
```

The service will start on `http://localhost:8080`.

## API Endpoints

### 1. Configure Cinema

**Endpoint**: `POST /configure`  
**Request Body**:

```json
{
  "rows": 10,
  "cols": 15,
  "min_distance": 6
}
```

**Response**:

- `200 OK`: `Cinema configured successfully`
- `400 Bad Request`: Validation errors (e.g., missing or invalid fields)

---

### 2. Query Available Seats

**Endpoint**: `POST /query`  
**Request Body**:

```json
{
  "group_size": 3
}
```

**Response**:

- `200 OK`: List of available seats (e.g., `[{"row": 0, "col": 0}, {"row": 0, "col": 1}, {"row": 0, "col": 2}]`)
- `400 Bad Request`: Validation errors

---

### 3. Reserve Seats

**Endpoint**: `POST /reserve`  
**Request Body**:

```json
{
  "seats": [
    { "row": 0, "col": 0 },
    { "row": 0, "col": 1 }
  ]
}
```

**Response**:

- `200 OK`: `Seats reserved successfully`
- `400 Bad Request`: Validation errors (e.g., invalid seat coordinates)
- `409 Conflict`: Seat already reserved

---

### 4. Cancel Seats

**Endpoint**: `POST /cancel`  
**Request Body**:

```json
{
  "seats": [{ "row": 0, "col": 0 }]
}
```

**Response**:

- `200 OK`: `Seats canceled successfully`
- `400 Bad Request`: Validation errors (e.g., invalid seat coordinates)
- `409 Conflict`: Seat not reserved

---

## Testing the API

1. **Using Postman**:

   - Import the Postman collection from the `/docs` folder.
   - Use the pre-configured requests to test the API.

2. **Using `curl`**:

   - Example: Configure cinema

     ```bash
     curl -X POST http://localhost:8080/configure \
     -H "Content-Type: application/json" \
     -d '{"rows": 10, "cols": 15, "min_distance": 6}'
     ```

   - Example: Reserve seats

     ```bash
     curl -X POST http://localhost:8080/reserve \
     -H "Content-Type: application/json" \
     -d '{"seats": [{"row": 0, "col": 0}, {"row": 0, "col": 1}]}'
     ```

## Validation Rules

- **Configure Cinema**:
  - `rows`, `cols`, and `min_distance` must be integers greater than or equal to 1.
- **Query Available Seats**:
  - `group_size` must be an integer greater than or equal to 1.
- **Reserve/Cancel Seats**:
  - Each seat must have `row` and `col` as integers greater than or equal to 0.

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux): For routing.
- [go-playground/validator](https://github.com/go-playground/validator): For input validation.
