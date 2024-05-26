# DoS Attack Simulator

## Overview

DoS Attack Simulator is a simple tool designed to simulate Denial of Service (DoS) attacks by sending a large number of HTTP requests to a target URL. This project is built using the Go programming language and leverages Go's concurrency model to handle multiple requests efficiently. The primary goal of this tool is to measure the performance and resilience of a target system under load.

## Features

- Configurable number of requests and concurrency levels
- Measures and reports average response time
- Utilizes Go's concurrency model for efficient request handling


## Requirements

- Go programming language (version 1.16 or later)
- Internet connection

## Project Structure

  dos-simulator/
- main.go: The main file that runs the DoS attack simulation.
- go.mod: Go module file that defines project dependencies.
- go.sum: File containing checksums of dependencies.

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/TolgaTD/dos-simulator.git
    cd dos-simulator
    ```

2. **Initialize Go modules**:
    ```sh
    go mod tidy
    ```

## Usage

To run the DoS Attack Simulator, use the following command:
```sh
go run main.go <target_url> <num_requests> <concurrency>
```

## Example
```sh
go run main.go http://example.com 1000 50
```

## Contributing
Contributions are welcome!

## License
This project is licensed under the MIT License. See the LICENSE file for details.
