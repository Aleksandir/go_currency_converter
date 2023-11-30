# Currency Converter

This is a command line tool for converting between different currencies. It uses an API to fetch the latest conversion rates.

## Installation

To install the tool, you need to have Go installed on your machine. You can download it from the official [Go website](https://golang.org/dl/). Once you have Go installed, you can clone this repository and build the tool using the following commands:

```bash
git clone https://github.com/username/currency-converter.git
cd currency-converter
go build ./cmd/converter
```

This will create an executable file named `converter` in the current directory.

## Usage

To use the tool, you need to pass three arguments: the amount you want to convert, the code of the currency you are converting from, and the code of the currency you are converting to. The currency codes should be in the ISO 4217 format. For example, to convert 100 US dollars to euros, you would use the following command:

```bash
./converter 100 USD EUR
```

The tool will print the converted amount to the standard output.

## Dependencies

This project uses the following dependencies:

- The `net/http` package for making HTTP requests to the API.
- The `encoding/json` package for parsing the API's JSON responses.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License.