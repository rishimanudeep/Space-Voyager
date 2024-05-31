# Space-Voyager

This is a Go application for managing exoplanet data. It provides RESTful API endpoints for performing CRUD operations on exoplanets and estimating fuel requirements for specific exoplanets.


## Getting Started with Space-Voyager

### Requirements

- A working Go environment - [https://go.dev/dl/](https://go.dev/dl/)
- Check the go version with command: go version.
- One should also be familiar with the Golang syntax. [Golang Tour](https://tour.golang.org/) has an excellent guided tour and highly recommended.

### Installation

## GOFR as dependency

- To get the GOFR as a dependency, use the command:
  `go get gofr.dev`

- Then use the command `go mod tidy`, to download the necessary packages.


### To Run Server

Run `go run main.go` command in CLI.

## Usage

The application provides the following RESTful endpoints:

- `GET /exoplanet`: Retrieve all exoplanets.
- `POST /exoplanet`: Create a new exoplanet.
- `GET /exoplanet/{id}`: Retrieve an exoplanet by ID.
- `PUT /exoplanet/{id}`: Update an exoplanet by ID.
- `DELETE /exoplanet/{id}`: Delete an exoplanet by ID.
- `GET /exoplanet/{id}/fuel_estimation`: Get fuel estimation for an exoplanet by ID.

## Dependencies

The application uses the following dependencies:

- `gofr.dev/pkg/gofr`: A Go web framework used for handling HTTP requests.
- `Exo-planet/handlers`: Handlers package for handling HTTP requests related to exoplanets.
- `Exo-planet/services`: Services package for business logic related to exoplanets.


For any information please reach out to me via rishimanudeepg@gmail.com