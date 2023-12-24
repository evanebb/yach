# Yet Another Configuration Handler

This small package aims to be a simple and extensible way to retrieve key-value bindings from different configuration
sources, primarily used to configure your application on startup.

## Introduction

This package contains the following:

- A `ConfigRetriever` interface - the interface that should be implemented by all configuration sources.
- A `ConfigHandler` struct - you add all the desired configuration sources to this object, and you can then retrieve
  values from it. If two sources have values for the same key, the value from the source that was added last will take
  precedence.
- A few different included configuration sources:
    - `ManualConfigSource` - can be used to manually set values, useful for setting default values and testing.
    - `EnvironmentConfigSource` - retrieves values from environment variables, with the ability to 'bind' abstract names
      to underlying environment variables and optional auto-binding.

Adding new/custom sources is simple; all you have to do is ensure that your type implements the `ConfigRetriever`
interface.

## Usage

```go
package main

import (
	"fmt"
	"github.com/evanebb/yach"
	"log"
)

func main() {
	// Create an instance of the ConfigHandler.
	ch := yach.NewConfigHandler()

	// Create and configure your configuration sources.
	manual := yach.NewManualConfigSource()
	manual.Set("database-host", "localhost")

	environment := yach.NewEnvironmentConfigSource()
	// We bind the 'database-host' key to the 'DB_HOST' environment variable.
	// This means that if the value for 'database-host' is requested, it will look up the value of the 'DB_HOST' environment variable.
	environment.Bind("database-host", "DB_HOST")

	// Push your configuration sources onto the ConfigHandler instance.
	// If multiple sources specify values for the same key, the order in which the sources were added determines which value takes precedence.
	// In this example, values from the EnvironmentConfigSource take precedence over those from the ManualConfigSource, since it was added onto the stack later.
	ch.Add(manual, environment)

	// Retrieve the configuration value for 'database-host' from the ConfigHandler instance.
	dbHost, err := ch.Get("database-host")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dbHost)
}

```
