# Data_Storage

## Technologies

- Go (Golang)
- Go Standard Library


## Architecture

This project follows a simple service-based architecture.

Application
     │
     ▼
Manager (Business Logic)
     │
     ▼
Users Slice (In-Memory Storage)

### Components

User Struct
Defines the user model with fields like FirstName, LastName, and Email.

Manager
Handles user operations such as:
- Adding users
- Searching users
- Validating input

Users Slice
Acts as an in-memory storage for user data.
 
