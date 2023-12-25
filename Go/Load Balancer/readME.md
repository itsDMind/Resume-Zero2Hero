# Load Balancer with Reverse Proxy in Go

## Introduction

This document provides an overview and explanation of the implemented code in Go for a simple load balancer with a reverse proxy. The code is designed to distribute incoming HTTP requests among multiple backend servers using a round-robin approach.
Code Structure
LoadBalancer Struct

The LoadBalancer struct manages the distribution of requests among backend servers. It contains the following fields:

    servers: A slice of strings representing the backend server addresses.
    index: An integer representing the index of the next server to which a request will be forwarded.
    mutex: A mutex for ensuring thread safety in a concurrent environment.

The NewLoadBalancer function initializes a new LoadBalancer instance with the provided list of backend servers.

The ServeHTTP method handles incoming HTTP requests. It acquires a lock to access and update the index, selects the next backend server in a round-robin fashion, and forwards the request to the selected server using a reverse proxy.
Proxy Struct

The Proxy struct represents a simple reverse proxy. It contains the following field:

    target: A string representing the target backend server address.

The NewProxy function creates a new Proxy instance with the specified target server.

The ServeHTTP method of the Proxy struct forwards the incoming HTTP request to the target server, copies the response headers, and writes the response back to the original response writer.
Main Function

The main function demonstrates the usage of the load balancer. It defines a list of backend servers (servers) and creates a new LoadBalancer instance. The load balancer is then started on port 8080 using http.ListenAndServe.
Usage

To run the code:

  - Ensure you have Go installed on your system.

  - Save the code in a file, e.g., main.go.

  - Open a terminal and navigate to the directory containing the file.

    Run the following command:

    ```bash

    go run main.go
    ```
    The load balancer will start, and you can access it by sending HTTP requests to http://localhost:8080.

Conclusion

This Go code provides a basic implementation of a load balancer with a reverse proxy. It can be extended and customized for more advanced use cases, such as load balancing algorithms, health checks, and logging.
