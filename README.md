# go-for-weather

## What the app does

The Go for Weather app is a simple web application that displays the current weather information. It is built using the Gin web framework for the backend and Tailwind CSS for the frontend.

## How to build the app

1. Clone the repository:
   ```sh
   git clone https://github.com/devopsjester/go-for-weather.git
   cd go-for-weather
   ```

2. Build the Docker image:
   ```sh
   docker build -t go-for-weather .
   ```

## How to run the app

1. Run the Docker container:
   ```sh
   docker run -p 8080:8080 go-for-weather
   ```

2. Open your web browser and go to `http://localhost:8080` to see the current weather information.
