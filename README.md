# Cinema Seating Service

## Features

- Configure cinema layout with rows, columns, and minimum distance rules.
- Query available seats for a group while maintaining social distancing.
- Reserve specific seats by their coordinates.
- Cancel reservations for specific seats.

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
