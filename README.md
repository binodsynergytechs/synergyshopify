# Synergy Shopify API Client (Go)

Synergy Shopify is a Go package that provides a full functional code for managing the Shopify REST API. With this package, you can easily interact with the Shopify API to perform various operations on your Shopify store.

## Installation

To install the Synergy package, you need to have Go installed on your system. Then, you can use the following command to install the package:

```shell
go get -u github.com/binodsynergytechs/synergyshopify
```
## Usage

To use the Synergy package in your Go project, you need to import the `synergyshopify` package. Here's an example of how to create a new API client and fetch the number of products:

```go

import "github.com/binodsynergytechs/synergyshopify"

    // Create a new Shopify app with your API key and password
    app := synergyshopify.App{
        ApiKey:    "apikey",
        Password:  "apipassword",
    }

    // Create a new API client (notice the token parameter is the empty string)
    client := synergyshopify.NewClient(app, "shopname", "")

    // Fetch the number of products.
    numProducts, err := client.Product.Count(nil)
