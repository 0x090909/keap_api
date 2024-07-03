# Keap API Go Library

## Overview

Keap bindings api

## Features

- **Lightweight:** Minimalistic and easy to use.
- **Placeholders:** Supports placeholder replacement with values.
- **Extensible:** Easily extendable for more complex use cases.

## Installation

To install Simple Template, use the following `go get` command:

```bash
go get github.com/0x090909/keap_api
```

## Usage

Below is an example of how to use the Simple Template library:

### Basic Usage

1. Import the library:

```go
import "github.com/0x090909/keap_api"
```

2. Create a new template and replace placeholders with actual values:

```go
package main

import (
	"context"
	"fmt"
	"github.com/0x090909/keap_api"
	v2 "github.com/0x090909/keap_api/v2"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"
	log "github.com/sirupsen/logrus"
)

func main() {

	// API requires no authentication, so use the anonymous
	authProvider, _ := auth.NewApiKeyAuthenticationProvider(
		"YOUR_API",
		"X-Keap-API-Key",
		auth.HEADER_KEYLOCATION)
	// Create request adapter using the net/http-based implementation
	adapter, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v\n", err)
	}

	if err != nil {
		fmt.Printf("Error creating request adapter: %v\n", err)
	}

	myKeapClient := keap_api.NewKeapClient(adapter)

	contacts, _ := myKeapClient.V2().Contacts().Get(context.Background(), &v2.ContactsRequestBuilderGetRequestConfiguration{})
	for _, contact := range contacts.GetContacts() {
		log.Info(*contact.GetGivenName())
	}
}

```

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

If you have any questions or suggestions, feel free to open an issue or reach out to us.

Happy coding!