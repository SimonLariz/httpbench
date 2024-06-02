# httpbench

`httpbench` is a command-line tool for benchmarking HTTP endpoints. It allows you to measure the performance of your web applications and APIs by sending concurrent requests and analyzing the response times and throughput.

## Features

- Send concurrent HTTP requests to a specified URL
- Customize the number of requests, concurrency level, and timeout duration
- Measure response times and calculate average, minimum, and maximum values
- Calculate throughput (requests per second)
- Display progress bar to track benchmarking progress
- Generate colorful and formatted output for easy readability

## Installation

To install `httpbench`, make sure you have Go installed on your system. Then, run the following command:

```bash
go get github.com/simonlariz/httpbench
```

## Usage 

To start benchmarking an HTTP endpoint, run the following command:

```bash
httpbench -url "http://example.com/api" -n 1000 -c 10 -timeout 5s
```

This command will send 1000 requests to `http://example.com/api` with a concurrency level of 10 and a timeout of 5 seconds.

- `-url`: The URL of the HTTP endpoint to benchmark
- `-n`: The total number of requests to send
- `-c`: The concurrency level (number of requests to send concurrently)
- `-timeout`: The timeout duration for each request

## Example Output

''ADD EXAMPLE OUTPUT HERE''

## Customization
`httpbench` provides various customization options for the progress bar and output formatting. You can modify the following properties:

- Progress bar width
- Progress bar template string
- Progress bar fill character
- Progress bar empty character
- Current and total count format
- Progress bar color
- Progress bar refresh rate

Refer to the source code and the `github.com/cheggaaa/pb` package documentation for more details on customization options.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).