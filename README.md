# Asset Locator Application

Welcome to the Asset Locator Application! This guide will help you run the application on your local system.

## Prerequisites

Before getting started, ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (if not already installed)
- [PostgreSQL](https://www.postgresql.org/download/)

## Instructions

1.  **Clone the Repository:**

    ```bash
    git clone https://github.com/vikash-parashar/asset-locator.git
    cd asset-locator
    ```

2.  **Set Up Environment Variables:**

        Create a `.env` file in the root of the project and set the following variables:

        ```env
        DB_HOST=your_database_host
        DB_PORT=your_database_port
        DB_USER=your_database_user
        DB_PASSWORD=your_database_password
        DB_NAME=your_database_name
        PORT=8080  # Or any desired port
        JWT_SECRET=your_custom_jwt_secret
        EMAIL_PASSWORD=your_email_password
        EMAIL_USERNAME=your_email
        S_SERVER=your_external_server_host
        S_PORT=your_external_server_port
        S_USER=your_external_server_username
        S_PASS=your_external_server_password
        APP_ENV=development

````

3. **Install Dependencies:**

   ```bash
   go get -d -v ./...
````

4. **Build the Executable:**

   For Windows:

   ```powershell
   $env:GOOS="windows"; $env:GOARCH="amd64"; go build -o asset_locator.exe
   ```

   For Linux:

   ```bash
   go build -o asset_locator
   ```

5. **Run the Application:**

   For Windows:

   ```powershell
   .\asset_locator.exe
   ```

   For Linux:

   ```bash
   ./asset_locator
   ```

6. **Access the Application:**

   Open your web browser and navigate to [http://localhost:8080](http://localhost:8080) (replace `8080` with the port you specified in the `.env` file).

7. **Explore the Application:**

   You can now use the Asset Locator application to manage and track your assets.

## Additional Notes

- If you encounter any issues during the setup, refer to the error messages and ensure that the prerequisites are correctly installed and configured.

- For production use, make sure to secure your PostgreSQL database and update the `.env` file accordingly.

- Feel free to reach out for assistance if you have any questions or encounter difficulties.

Happy exploring!
