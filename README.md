# Go STAR Bus 83 Live

A Go application that consumes STAR Rennes Open Data APIs to display real-time Bus 83 arrivals, stops, destinations, and live transit information.

## Description

This project connects to the STAR Rennes Open Data API and retrieves real-time information about Bus 83 operations across the Rennes metropolitan area.

The application fetches live transit data in JSON format, processes the response using Go structures, and displays upcoming Bus 83 arrivals, stop names, destinations, and real-time schedule information directly in the terminal.

## Features

- Retrieve live data from the STAR Rennes Open Data API
- Display Bus 83 real-time arrivals
- Show stop names
- Show bus destinations
- Display expected arrival times
- Filter data to keep only Bus 83 records
- Count the number of Bus 83 passages found
- Parse and process JSON responses
- Handle HTTP and JSON errors
- Use a custom HTTPS client for environments requiring TLS certificate bypass during testing
