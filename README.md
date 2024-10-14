
# 🚀 **Mereb Technologies Challenge: Person Management API**

## 📚 **Project Overview**
The **Person Management API** is designed for the **Mereb Technologies Challenge**. It provides a RESTful interface for managing person records through basic CRUD (Create, Read, Update, Delete) operations, incorporating essential data validation to ensure data integrity.

---

## 🛠️ **Technologies Used**
- **Go**: A statically typed, compiled programming language that is efficient and fast.
- **Gin**: A lightweight web framework for Go, known for its speed and minimal overhead.
- **In-Memory Database**: Utilizes an in-memory database for temporary storage of person records.

---

## ✨ **Key Features**
- **Create Person**: Add new person records to the database.
- **Get All Persons**: Retrieve a list of all person records.
- **Get Person**: Fetch a specific person record by ID.
- **Update Person**: Modify existing person records.
- **Delete Person**: Remove person records from the database.
- **Data Validation**: Validate incoming data to ensure it meets predefined criteria.

---

## 📦 **Setup Instructions**
To set up the project, follow these steps:

1. **Clone or Download the Repository**:
   ```bash
   git clone https://github.com/nuredinbedruyimer/Mereb_task-V1
2. **Install Dependencies: Run the following command to install necessary dependencies**:
   ```bash
      go mod tidy
3. **Start the Application: Launch the application using**:
   ```
      go run server.go
**Example JSON for Creating a Person**:

    {
    "id": "5d045199-4fcc-4713-8505-174305565f6a",
    "name": "Abeneze ADugna 1",
    "age": 81,
    "hobbies": [
        "Watching Movies",
        "Reading Books"
    ]
}
## 🌐 **Usage Instructions**
You can interact with the API using the following endpoints:

| HTTP Method | Endpoint                       | Description                          |
|-------------|--------------------------------|--------------------------------------|
| **POST**    | `http://localhost/persons`     | Create a new person record          |
| **GET**     | `http://localhost/persons`     | Retrieve a list of all persons      |
| **GET**     | `http://localhost/persons/:id` | Fetch a specific person by ID       |
| **PUT**     | `http://localhost/persons/:id` | Update an existing person record     |
| **DELETE**  | `http://localhost/persons/:id` | Delete a person record              |

## 🗂️ **Code Structure**:
 The project structure is organized as follows:
 



person-management-api/ ├── controllers/ # Handles HTTP requests and responses ├── services/ # Contains business logic for processing requests ├── models/ # Defines data structures used in the application (e.g., Person) ├── constants/ # Stores constant values such as messages and responses ├── config/ # Manages application configuration, including CORS setup ├── routes/ # Organizes route definitions for the API ├── database/ # Manages database connections and operations ├── helpers/ # Contains utility functions, including data validation ├── .env # Stores environment variables for configuration ├── go.mod # Manages module dependencies └── go.sum # Manages module dependencies




## 💡 **Conclusion**
The Person Management API serves as a robust solution for managing person records using Go and Gin. It emphasizes data validation and structured code organization, making it maintainable and extendable for future enhancements.


   
