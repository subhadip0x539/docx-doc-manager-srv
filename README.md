# docx-doc-manager-srv

## Description

This service provides the API for uploading and downloading files using MongoDB's GridFS storage system. GridFS is a specification used in MongoDB to store and retrieve large files efficiently.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/docx-doc-manager-srv.git
   ```

2. Install dependencies:

   ```
   cd docx-doc-manager-srv
   go mod tidy
   ```

3. Set up environment variables:

   Create a `.env` file in the root directory and provide the following variables:

   ```bash
   SERVER_NAME=docx-doc-manager-srv
   SERVER_HOST=0.0.0.0
   SERVER_PORT=8080
   DB_URI=$DB_URI # Your mongodb connection url
   DB_DATABASE=filestore
   DB_TIMEOUT=10000
   ```

## Usage

1. Start the server:

   ```bash
   go run ./cmd/app/main.go
   ```

2. Use API endpoints to interact with the service.

## Compiling and Executing the Binary

1. Build the binary:

   ```bash
   make build
   ```

2. Run the executable:

   ```bash
   make run
   ```

## Running with Docker

1. Build the Docker image:

   ```bash
   docker build -t docx-doc-manager-srv .
   ```

2. Run the Docker container:

   ```bash
   sudo docker run -p 8080:8080 -e .env --name docx-doc-manager-srv docx-doc-manager-srv
   ```
