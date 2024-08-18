# Forecast

Welcome to Forecast!

This is my first GoLang CLI program.

## Overview

Forecast is a simple command-line interface (CLI) application written in Go that fetches and displays weather information for a specified city. It uses the OpenWeatherMap API to retrieve current weather data.

## Features

- Fetches current weather information for a specified city.
- Displays temperature, humidity, and weather description.
- Handles city names with multiple words.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/biowulf21/forecast
    cd forecast
    ```

2. Build the application:

    ```sh
    go build -o forecast
    ```

## Usage

To use the Forecast CLI, run the following command:

```sh
./forecast [city name]
```

- If no city name is provided, it defaults to "Cagayan de Oro City".
- Example:

    ```sh
    ./forecast "New York"
    ```

## Example Output

```sh
Weather for New York is New York (clear sky), which has a humidity of 50%, and a temperature of 25.00°C
```

## Configuration

You need to set your OpenWeatherMap api key via env variables

`export FORECAST_API=`

## Dependencies

- Go 1.16 or later

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [OpenWeatherMap](https://openweathermap.org/) for providing the weather API.

## Contributing

Feel free to submit issues or pull requests if you have any improvements or suggestions.

## Contact

For any questions or inquiries, please contact jamesjilhaney21@gmail.com
