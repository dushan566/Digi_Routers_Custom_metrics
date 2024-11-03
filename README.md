# Digi Router SIM Status Monitor

This Go application monitors the status of Digi Router SIM cards using a custom Prometheus metric exporter. It collects data on the primary and secondary SIM status of routers in various farm locations, enabling real-time monitoring of internet connectivity status. The metrics are then visualized on Grafana dashboards.

## Features

- **Custom Metrics**: Uses `prometheus/client_golang` to create and export custom metrics.
- **Prometheus Integration**: Provides metrics in a format consumable by Prometheus.
- **Grafana Visualization**: Metrics can be visualized on Grafana dashboards for easy monitoring.
- **SIM Status Monitoring**: Tracks the active primary and secondary SIM status for each Digi Router, providing insights into the connectivity status of each farm's internet connection.

## Libraries Used

This application uses the following Go libraries:
- [github.com/prometheus/client_golang/prometheus](https://github.com/prometheus/client_golang) - For defining and registering custom metrics.
- [github.com/prometheus/client_golang/prometheus/promhttp](https://github.com/prometheus/client_golang) - For exposing metrics over HTTP.

## How It Works

1. **Fetches Data from Digi Remote Manager**: The application calls the Digi Remote Manager device and stream API endpoints to retrieve the list of connected routers and the status of SIM1 and SIM2 for each router.
2. **Custom Metrics Exporter**: Extracted data is processed and exposed as custom metrics via an HTTP endpoint for Prometheus to scrape.
3. **Grafana Visualization**: Once scraped by Prometheus, the metrics can be visualized on Grafana, enabling real-time insights into the connectivity status of all routers in the monitored farms.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or later)
- [Prometheus](https://prometheus.io/download/)
- [Grafana](https://grafana.com/get)

### Installation (Local)

1. Clone this repository:
   ```
   git clone hhttps://github.com/dushan566/Digi_Routers_Custom_metrics.git
   cd Digi_Routers_Custom_metrics
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```
3. Set up environment variables for authentication and API endpoint in `secrets/credentials.env`:
   - **DRMUSER**: Digi Remote Manager username
   - **DRMPASS**: Digi Remote Manager password
   - **DEVICE_ENDPOINT**: Digi Remote Manager Device API endpoint
   - **STREAM_ENDPOINT**: Digi Remote Manager Stream API endpoint

4. Run the application locally:
   ```
   go run main.go
   ```

### Use in Kubernetes

 1. Build and Push Docker Image
    Build the Docker image for your application.
    Push the image to a container registry (e.g., Docker Hub).

 2. Create ConfigMap and Secret
    Create a ConfigMap for non-sensitive configuration (e.g., API endpoint).
    Create a Secret for sensitive information (e.g., DRMUSER, DRMPASS).

 3. Create a Kubernetes Deployment
    Define a Deployment to run your application with environment variables pointing to the ConfigMap and Secret.

 4. Create a Kubernetes Service
    Define a Service to expose your application to Prometheus within the cluster.

 5. Configure Prometheus Scraping
    Add a scrape configuration in prometheus.yml to target your application's Service.

 6. Deploy and Verify
    Apply all configurations (ConfigMap, Secret, Deployment, Service) to your Kubernetes cluster.
    Check that your application is running and metrics are accessible for Prometheus.

 7. Visualize Metrics in Grafana
    Add Prometheus as a data source in Grafana.
    Create a dashboard to visualize your application’s metrics.


## License
This project is licensed under the MIT License.