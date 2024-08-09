# GameBuy

GameBuy is a backend application for managing game purchases and transactions.

## Purpose

The purpose of this project is to provide a robust backend system for a game store, allowing users to browse and purchase games

## Usage

To run the GameBuy application:

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Set up your PostgreSQL database and update the configuration in `configs/config.json`
4. Run the application: `go run main.go`
5. Access all the API with Role Admin ( Username : admin & Password : admin )
6. If Sign up A new user, just have role buyer

The server will start on `localhost:8080` by default.

## Available Endpoints

### Users
- `POST /api/login` - Login user
- `POST /api/signup` - Create a new user just buyer

### Games
- `GET /api/games` - Get a list of all games
- `POST /api/games` - Add a new game
- `GET /api/games/{id}` - Get game details
- `PUT /api/games/{id}` - Update game information
- `DELETE /api/games/{id}` - Remove a game

### transaksi
- `GET /api/transaksi` - Get a list of all transactions
- `POST /api/transaksi` - Create a new transaction
- `GET /api/transaksi/{id}` - Get transaction details
- `DELETE /api/transaksi/{id}` - Remove a transaction

### Categories
- `GET /api/categories` - Get a list of all categories
- `POST /api/categories` - Create a new category
- `GET /api/categories/{id}` - Get category details
- `PUT /api/categories/{id}` - Update category information
- `DELETE /api/categories/{id}` - Delete a category

### Platforms
- `GET /api/platforms` - Get a list of all platforms
- `POST /api/platforms` - Add a new platform
- `GET /api/platforms/{id}` - Get platform details
- `PUT /api/platforms/{id}` - Update platform information
- `DELETE /api/platforms/{id}` - Remove a platform

## Configuration

The application configuration is stored in `configs/config.json`. Make sure to update the database connection details before running the application.
