🚀 Mereb Technologies Challenge: Person Management API
📚 Project Overview
The Person Management API is designed for the Mereb Technologies Challenge. It provides a RESTful interface for managing person records through basic CRUD (Create, Read, Update, Delete) operations, incorporating essential data validation to ensure data integrity.

🛠️ Technologies Used
Go: A statically typed, compiled programming language that is efficient and fast.
Gin: A lightweight web framework for Go, known for its speed and minimal overhead.
In-Memory Database: Utilizes an in-memory database for temporary storage of person records.
✨ Key Features
Create Person: Add new person records to the database.
Get All Persons: Retrieve a list of all person records.
Get Person: Fetch a specific person record by ID.
Update Person: Modify existing person records.
Delete Person: Remove person records from the database.
Data Validation: Validate incoming data to ensure it meets predefined criteria.
📦 Setup Instructions
To set up the project, follow these steps:

Clone or Download the Repository:

bash
Copy code
git clone <repository_url>
cd <repository_directory>
Set Up Environment Variables: Ensure your environment variables are correctly configured.

Install Dependencies: Run the following command to install necessary dependencies:

bash
Copy code
go mod tidy
Start the Application: Launch the application using:

bash
Copy code
go run main.go
🌐 Usage Instructions
You can interact with the API using the following endpoints:

Create Person: POST http://localhost/persons
Get All Persons: GET http://localhost/persons
Get Person: GET http://localhost/persons/:id
Update Person: PUT http://localhost/persons/:id
Delete Person: DELETE http://localhost/persons/:id
Example JSON for Creating a Person:
json
Copy code
{
    "id": "5d045199-4fcc-4713-8505-174305565f6a",
    "name": "Abeneze ADugna 1",
    "age": 81,
    "hobbies": [
        "Watching Movies",
        "Reading Books"
    ]
}
🗂️ Code Structure
The project structure is organized as follows:

controllers/: Handles HTTP requests and responses.
services/: Contains business logic for processing requests.
models/: Defines data structures used in the application (e.g., Person).
constants/: Stores constant values such as messages and responses.
config/: Manages application configuration, including CORS setup.
routes/: Organizes route definitions for the API.
database/: Manages database connections and operations.
helpers/: Contains utility functions, including data validation.
.env: Stores environment variables for configuration.
go.mod & go.sum: Manage module dependencies.
📋 Data Model
The Person model includes the following fields:

PersonID: Unique identifier (auto-incremented).
Name: The person's name (required, min length: 2, max length: 100).
Age: The person's age (required, must be between 0 and 120).
Hobbies: A list of hobbies associated with the person (optional).
✅ Validation Rules
The following validation rules are implemented for the Person model:

Name:
Required
Minimum length: 2 characters
Maximum length: 100 characters
Age:
Required
Must be between 0 and 120 (inclusive)
Hobbies:
Optional list of strings
💡 Conclusion
The Person Management API serves as a robust solution for managing person records using Go and Gin. It emphasizes data validation and structured code organization, making it maintainable and extendable for future enhancements.

