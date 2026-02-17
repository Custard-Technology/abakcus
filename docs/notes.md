# Implementation Notes

## 17/02/2026

- **Added `internal/config` package** to centralize configuration handling. It reads `MONGO_URI` and
  `MONGO_DB` from environment variables and validates their presence. This keeps
  configuration logic out of `main.go` and supports future expansion.

- **Created `internal/repository/mongo` package** with `NewClient` constructor. The function
  performs a timed connection and ping to MongoDB, returning an error if the
  handshake fails. Separation respects the "repository handles data storage"
  rule and makes the code easier to test or replace later.

- **Implemented `cmd/api/main.go` entry point** which loads config, establishes
  the MongoDB connection, logs success/failure, and handles graceful shutdown.
  Using `log` from the standard library satisfies the preference for minimal
  dependencies.

- **Added unit tests** for both the config loader and connection error path. Tests
  enforce explicit error handling and improve maintainability by catching
  regressions early.

- **Updated `README.md` and `.env`** with environment variable guidance and
  startup instructions. This ensures other developers or deployment scripts
  understand how to run the service.

These changes align with the project rules (readability, modularity, explicit
error handling, standard library use) and satisfy the acceptance criteria of
ticket BE-001: successful MongoDB connections are logged, and failures are
handled robustly. The structure also makes future enhancements (pooling,
retries, metrics) straightforward.

## Menu Management Implementation (BE-002+)

### Architecture

Implemented a complete menu CRUD API following the three-layer pattern from rules.md:

1. **Domain Layer** (`internal/models/menu.go`)
   - `Menu` struct with BSON tags for MongoDB serialization and JSON tags for HTTP
   - `CreateMenuRequest` and `UpdateMenuRequest` for input validation
   - `MenuItem` struct prepared for future menu items feature

2. **Repository Layer** (`internal/repository/mongo/menu.go`)
   - `MenuRepositoryI` interface defined for testability
   - `MenuRepository` concrete implementation with CRUD operations
   - All database operations include 5-second timeouts to prevent hanging
   - Explicit error handling for duplicate keys, missing documents, and validation failures

3. **Service Layer** (`internal/service/menu.go`)
   - `MenuService` implements business logic and validation
   - UUID generation for menu IDs using `github.com/google/uuid`
   - Timestamp management (created_at, updated_at auto-populated)
   - Separation of concerns: service doesn't know about HTTP

4. **Handler Layer** (`internal/handler/menu.go`)
   - HTTP request/response handling for CRUD endpoints
   - Extracts `X-Business-ID` header as placeholder for session/user ID
   - JSON marshalling/unmarshalling with proper Content-Type headers
   - Consistent error responses with appropriate HTTP status codes (201, 200, 204, 400, 404, 500)

### HTTP Endpoints

- `POST /menus` – Create menu (requires X-Business-ID header, returns 201 Created)
- `GET /menus` – List menus for a business (requires X-Business-ID header, returns 200 OK)
- `GET /menus/:menu_id` – Retrieve single menu (returns 200 OK or 404 Not Found)
- `PUT /menus/:menu_id` – Update menu name/description/active status (returns 200 OK)
- `DELETE /menus/:menu_id` – Soft delete menu (returns 204 No Content)

### Testing Strategy

- **Unit tests for service layer** (`internal/service/menu_test.go`)
  - Mock repository implementation to isolate business logic
  - Tests cover: creation, retrieval, deletion, listing by business
  - Mock verifies interface compliance with compile-time check

- **Unit tests for handler layer** (`internal/handler/menu_test.go`)
  - HTTP request simulation with `httptest`
  - Validates status codes, error messages, and JSON responses
  - No external dependencies or real database calls

### Why These Decisions

- **net/http.ServeMux** chosen over third-party routers to honor "standard library first" rule.
  Go 1.22+ provides sufficient routing via ServeMux; middleware can be layered later if needed.

- **MenuRepositoryI interface** abstraction enables mock testing and future database swaps.
  Testability is a first-class requirement per rules.md.

- **X-Business-ID header** as temporary session ID placeholder allows quick API testing.
  Will be replaced with authenticated user ID on full auth implementation (no schema changes needed).

- **UUID for menu IDs** chosen over sequential IDs for distributed scalability and
  collision-free generation across systems.

- **Explicit error handling throughout** with context timeouts prevents database hanging and
  provides clients with meaningful error messages (e.g., "menu not found", "business_id required").

### Next Steps

- Implement `/menus/:menu_id/items` endpoints for menu items (MenuItem struct ready in models)
- Add authentication middleware to extract real user ID instead of header placeholder
- Consider adding pagination for `GET /menus` when business has many menus
- Add request validation middleware (e.g., max name length, required fields)
- Implement soft delete recovery or permanent delete based on business requirements

### Miscellaneous

- Added simple CORS middleware in `cmd/api/main.go` to allow cross‑origin requests
  from the frontend during development. The handler permits `*` by default; restrict
  to specific origins before deploying to production.


Frontend Developer API Consumption Guide
Overview
The backend exposes a RESTful API for menu management. This guide explains how frontend developers consume these endpoints.

Base URL
Authentication & Headers
Currently, the API uses a temporary session ID approach (no user authentication yet).

Required Header:

Example:

Menu Endpoints
1. Create a Menu
POST /menus

Request:

Headers:

cURL Example:

Response (201 Created):

Error Responses:

400 Bad Request – Missing/invalid fields
500 Internal Server Error – Database failure
2. List All Menus (for a Business)
GET /menus

Headers:

cURL Example:

Response (200 OK):

3. Get a Single Menu
GET /menus/:menu_id

cURL Example:

Response (200 OK):

Error Responses:

404 Not Found – Menu does not exist
4. Update a Menu
PUT /menus/:menu_id

Request:

Headers:

cURL Example:

Response (200 OK):

Error Responses:

400 Bad Request – Invalid input
404 Not Found – Menu does not exist
500 Internal Server Error – Database failure
5. Delete a Menu
DELETE /menus/:menu_id

Headers:

cURL Example:

Response (204 No Content)

Error Responses:

404 Not Found – Menu does not exist
500 Internal Server Error – Database failure
JavaScript/Fetch Examples
Create Menu
List Menus
Get Single Menu
Update Menu
Delete Menu
TypeScript Types
Status Codes Summary
Code	Meaning
200	OK – Request succeeded
201	Created – Resource created successfully
204	No Content – Success with no response body (DELETE)
400	Bad Request – Invalid input/missing fields
404	Not Found – Resource does not exist
500	Internal Server Error – Server-side error
Notes for Frontend Developers
Session ID is temporary – The X-Business-ID header will be replaced with JWT/OAuth authentication when user accounts are implemented.
CORS – If consuming from a different domain, backend CORS headers will need to be configured (currently not set).
Timestamps – All dates are in ISO 8601 format (UTC). Parse with new Date(timestamp) in JavaScript.
Soft Delete – Deleting a menu sets is_active to false rather than permanently removing it. Filtering by is_active is recommended on the frontend.