# Init Functions in Go

This section demonstrates the init function in Go - understanding how this special function works for automatic initialization and setup tasks before the main program execution begins.

## Overview

The init function is a special built-in function in Go that is automatically executed before the main function runs. Unlike regular functions, init functions cannot be called explicitly, have no parameters, and return no values. They are primarily used for initialization tasks such as setting up global variables, establishing database connections, loading configuration files, or performing any setup work that needs to happen before the main program logic executes. Understanding init functions is crucial for proper program initialization and ensuring that your application starts in a well-defined state.

## Code Example

```go
package main

import "fmt"

var a int = 0

func main() {
	fmt.Println("main function")
	fmt.Println("a =", a)
}

//! 'init' function -> we cannot call this function, it will be called automatically
//! 'init' function doesn't take any input, doesn't give any output also. It automatically gets executed
func init() {
	a = 10
	fmt.Println("a =", a)
	fmt.Println("init function")
}

//! we all know that, when a program runs, it first runs the 'main' function first, but now we will learn something new. In 'go', a program will first check is there any 'init' function, if yes, it will execute the 'init' function first.
```

## How This Code Works

### 1. Package Declaration and Import

```go
package main

import "fmt"
```

- Declares the package as `main`, making this an executable program
- Imports the `fmt` package for formatted I/O operations
- The import statement itself can trigger init functions in imported packages

### 2. Global Variable Declaration

```go
var a int = 0
```

- **Global Variable**: `a` is declared at the package level with initial value 0
- **Package Scope**: This variable is accessible throughout the entire package
- **Initialization Order**: Global variables are initialized before init functions run
- **Modification Target**: This variable will be modified by the init function

### 3. Main Function Declaration

```go
func main() {
	fmt.Println("main function")
	fmt.Println("a =", a)
}
```

- **Entry Point**: The `main` function serves as the program's entry point
- **Execution Order**: This runs after all init functions have completed
- **Variable Access**: Displays the value of `a` after init function modification
- **Expected Output**: Will show `a = 10` (not the original value 0)

### 4. Init Function Declaration

```go
func init() {
	a = 10
	fmt.Println("a =", a)
	fmt.Println("init function")
}
```

- **Variable Modification**: Changes the global variable `a` from 0 to 10
- **Automatic Execution**: Called automatically before main function execution
- **No Parameters**: Init functions cannot accept any parameters
- **No Return Values**: Init functions cannot return any values
- **Initialization Purpose**: Demonstrates how init functions can modify global state

### 5. Program Execution Flow

1. **Package Loading**: Go runtime loads the main package
2. **Global Variable Initialization**: `a` is set to 0
3. **Init Function Execution**: `a` is modified to 10, and init messages are printed
4. **Main Function Execution**: Main function runs and shows the modified value of `a`
5. **Output Order**: Init output appears before main output

## Init Function Characteristics

### Basic Properties

```go
// Basic init function syntax
func init() {
    // initialization code here
    // no parameters allowed
    // no return values allowed
}
```

### Key Features of Init Functions

1. **Automatic Execution**: Called automatically, cannot be invoked manually
2. **No Signature**: Cannot have parameters or return values
3. **Multiple Allowed**: A single package can have multiple init functions
4. **Execution Order**: Run before main function in dependency order
5. **Package Scope**: Each package can have its own init functions
6. **One-Time Execution**: Each init function runs exactly once

## Init Function Execution Order

### Single Package with Multiple Init Functions

```go
package main

import "fmt"

func init() {
	fmt.Println("First init function")
}

func init() {
	fmt.Println("Second init function")
}

func init() {
	fmt.Println("Third init function")
}

func main() {
	fmt.Println("Main function")
}

// Output:
// First init function
// Second init function
// Third init function
// Main function
```

### Multiple Files in Same Package

```go
// file1.go
package main

import "fmt"

func init() {
	fmt.Println("Init from file1")
}

// file2.go
package main

import "fmt"

func init() {
	fmt.Println("Init from file2")
}

// main.go
package main

import "fmt"

func init() {
	fmt.Println("Init from main")
}

func main() {
	fmt.Println("Main function")
}

// Execution order depends on file compilation order
```

### Package Dependency Chain

```go
// Package A
package a

import "fmt"

func init() {
	fmt.Println("Package A init")
}

// Package B (imports A)
package b

import (
	"fmt"
	_ "path/to/a"  // Blank import to trigger init
)

func init() {
	fmt.Println("Package B init")
}

// Main package (imports B)
package main

import (
	"fmt"
	_ "path/to/b"
)

func init() {
	fmt.Println("Main package init")
}

func main() {
	fmt.Println("Main function")
}

// Output:
// Package A init
// Package B init
// Main package init
// Main function
```

## Common Use Cases for Init Functions

### Use Case 1: Global Variable Initialization

```go
package main

import (
	"fmt"
	"time"
)

var (
	startTime time.Time
	config    map[string]string
	isReady   bool
)

func init() {
	fmt.Println("Initializing global variables...")

	// Record application start time
	startTime = time.Now()

	// Initialize configuration map
	config = make(map[string]string)
	config["version"] = "1.0.0"
	config["environment"] = "development"

	// Set ready flag
	isReady = true

	fmt.Println("Initialization complete!")
}

func main() {
	fmt.Printf("Application started at: %v\n", startTime)
	fmt.Printf("Version: %s\n", config["version"])
	fmt.Printf("Ready: %t\n", isReady)
}
```

### Use Case 2: Resource Setup and Configuration

```go
package main

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// Set up custom logger
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	logger = log.New(logFile, "APP: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Logger initialized")

	fmt.Println("Logger setup complete")
}

func main() {
	logger.Println("Application started")
	fmt.Println("Main function executing")
}
```

### Use Case 3: Validation and Precondition Checks

```go
package main

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	fmt.Println("Performing system checks...")

	// Check Go version
	if runtime.Version() < "go1.18" {
		fmt.Fprintf(os.Stderr, "Warning: Go version %s may not be supported\n", runtime.Version())
	}

	// Check required environment variables
	requiredEnvVars := []string{"HOME", "PATH"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			fmt.Fprintf(os.Stderr, "Warning: Required environment variable %s not set\n", envVar)
		}
	}

	fmt.Println("System checks complete")
}

func main() {
	fmt.Println("Application running with validated environment")
}
```

### Use Case 4: Registration and Plugin Loading

```go
package main

import "fmt"

type Plugin interface {
	Name() string
	Execute() string
}

type TextPlugin struct{}

func (p TextPlugin) Name() string { return "text-processor" }
func (p TextPlugin) Execute() string { return "Processing text..." }

var pluginRegistry = make(map[string]Plugin)

func init() {
	fmt.Println("Loading plugins...")

	// Register built-in plugins
	registerPlugin(TextPlugin{})

	fmt.Printf("Loaded %d plugins\n", len(pluginRegistry))
}

func registerPlugin(p Plugin) {
	pluginRegistry[p.Name()] = p
	fmt.Printf("Registered plugin: %s\n", p.Name())
}

func main() {
	fmt.Println("Available plugins:")
	for name, plugin := range pluginRegistry {
		fmt.Printf("- %s: %s\n", name, plugin.Execute())
	}
}
```

## Output

The current program produces the following output:

```
a = 10
init function
main function
a = 10
```

**Explanation:**

1. The init function executes first, modifying `a` from 0 to 10
2. Init function prints "a = 10" and "init function"
3. Then the main function executes as the program entry point
4. Main function prints "main function" and "a = 10" (showing the modified value)
5. This demonstrates both execution order and how init functions can modify global variables

## Running the Code

To run this example:

```bash
go run main.go
```

## Try It Yourself

Experiment with different init function scenarios:

### Experiment 1: Multiple Init Functions with Variable Modification

```go
package main

import "fmt"

var counter int = 0
var message string = "uninitialized"

func init() {
	counter++
	fmt.Printf("First init: counter = %d\n", counter)
}

func init() {
	counter++
	message = "partially initialized"
	fmt.Printf("Second init: counter = %d, message = %s\n", counter, message)
}

func init() {
	counter++
	message = "fully initialized"
	fmt.Printf("Third init: counter = %d, message = %s\n", counter, message)
}

func main() {
	fmt.Printf("Main: counter = %d, message = %s\n", counter, message)
}
```

### Experiment 2: Init with Global Variable Setup

```go
package main

import (
	"fmt"
	"os"
)

var appConfig map[string]string
var isConfigured bool = false

func init() {
	fmt.Println("Initializing application configuration...")

	// Initialize configuration map
	appConfig = make(map[string]string)
	appConfig["version"] = "1.0.0"
	appConfig["environment"] = os.Getenv("APP_ENV")

	if appConfig["environment"] == "" {
		appConfig["environment"] = "development"
	}

	isConfigured = true
	fmt.Printf("Configuration loaded: %+v\n", appConfig)
}

func main() {
	if isConfigured {
		fmt.Printf("App version: %s\n", appConfig["version"])
		fmt.Printf("Environment: %s\n", appConfig["environment"])
	} else {
		fmt.Println("Configuration not loaded!")
	}
}
```

### Experiment 3: Init with Complex Initialization

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	randomSeed  int64
	gameConfig  map[string]int
	playerStats []string
)

func init() {
	fmt.Println("Game initialization starting...")

	// Initialize random seed
	randomSeed = time.Now().UnixNano()
	rand.Seed(randomSeed)

	// Setup game configuration
	gameConfig = map[string]int{
		"maxPlayers":  4,
		"timeLimit":   300, // 5 minutes
		"scoreTarget": 100,
	}

	// Initialize player stats
	playerStats = make([]string, 0)

	fmt.Printf("Game initialized with seed: %d\n", randomSeed)
}

func main() {
	fmt.Printf("Game config: %+v\n", gameConfig)
	fmt.Println("Game ready to start!")
}
```

## Best Practices for Init Functions

### 1. Keep Init Functions Simple and Fast

```go
// Good: Simple, fast initialization
func init() {
	config = loadDefaultConfig()
	logger = setupLogger()
}

// Avoid: Complex, time-consuming operations
func init() {
	// Don't do expensive operations like:
	// - Network calls
	// - Heavy file I/O
	// - Complex computations
}
```

### 2. Handle Errors Appropriately

```go
// Good: Proper error handling
func init() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	globalConfig = config
}

// Consider: Graceful degradation
func init() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Printf("Warning: Using default config due to error: %v", err)
		config = getDefaultConfig()
	}

	globalConfig = config
}
```

### 3. Use Init for Setup Only

```go
// Good: Use init for initialization
func init() {
	database.Connect(connectionString)
	cache.Initialize(cacheSize)
	metrics.Setup()
}

// Avoid: Business logic in init
func init() {
	// Don't put main application logic here
	// processUserData()
	// generateReports()
}
```

### 4. Document Init Function Purpose

```go
// init initializes the HTTP client with custom timeout settings
// and sets up the default request headers for API communication.
func init() {
	httpClient = &http.Client{
		Timeout: 30 * time.Second,
	}

	defaultHeaders = map[string]string{
		"User-Agent":   "MyApp/1.0",
		"Content-Type": "application/json",
	}
}
```

## Advanced Init Function Patterns

### Pattern 1: Conditional Initialization

```go
package main

import (
	"fmt"
	"os"
)

func init() {
	environment := os.Getenv("APP_ENV")

	switch environment {
	case "production":
		setupProductionEnvironment()
	case "staging":
		setupStagingEnvironment()
	default:
		setupDevelopmentEnvironment()
	}
}

func setupProductionEnvironment() {
	fmt.Println("Production environment initialized")
	// Production-specific setup
}

func setupStagingEnvironment() {
	fmt.Println("Staging environment initialized")
	// Staging-specific setup
}

func setupDevelopmentEnvironment() {
	fmt.Println("Development environment initialized")
	// Development-specific setup
}

func main() {
	fmt.Println("Application running")
}
```

### Pattern 2: Init with Dependency Management

```go
package main

import "fmt"

type Service struct {
	name string
}

func (s *Service) Start() error {
	fmt.Printf("Starting service: %s\n", s.name)
	return nil
}

var services []*Service

func init() {
	// Initialize services in dependency order
	services = []*Service{
		{name: "Database"},
		{name: "Cache"},
		{name: "Web Server"},
	}

	// Start all services
	for _, service := range services {
		if err := service.Start(); err != nil {
			panic(fmt.Sprintf("Failed to start service %s: %v", service.name, err))
		}
	}

	fmt.Println("All services initialized")
}

func main() {
	fmt.Printf("Application ready with %d services\n", len(services))
}
```

### Pattern 3: Init with Resource Pool

```go
package main

import (
	"fmt"
	"sync"
)

type ConnectionPool struct {
	connections chan *Connection
	maxSize     int
}

type Connection struct {
	id string
}

var dbPool *ConnectionPool

func init() {
	fmt.Println("Initializing database connection pool...")

	dbPool = &ConnectionPool{
		connections: make(chan *Connection, 10),
		maxSize:     10,
	}

	// Pre-populate the pool
	for i := 0; i < 5; i++ {
		conn := &Connection{id: fmt.Sprintf("conn-%d", i)}
		dbPool.connections <- conn
	}

	fmt.Printf("Connection pool initialized with %d connections\n", len(dbPool.connections))
}

func main() {
	fmt.Printf("Database pool ready with capacity: %d\n", dbPool.maxSize)
}
```

## Init Function Limitations and Considerations

### 1. Cannot Be Called Directly

```go
// ERROR: Cannot call init function
func main() {
	init() // Compilation error
}
```

### 2. No Parameters or Return Values

```go
// ERROR: Init cannot have parameters
func init(config string) { // Compilation error
}

// ERROR: Init cannot return values
func init() string { // Compilation error
	return "initialized"
}
```

### 3. Execution Order Dependencies

```go
// Be careful with global variable dependencies
var a int
var b int = a + 1 // b depends on a being initialized

func init() {
	a = 10
	// b is already calculated as 0 + 1 = 1, not 10 + 1 = 11
}
```

### 4. Testing Considerations

```go
// Init functions run during testing too
func init() {
	// This will run for every test
	setupTestEnvironment()
}

// Consider using build tags for test-specific init
//go:build !test
func init() {
	// This won't run during testing
	setupProductionEnvironment()
}
```

## Memory and Performance Considerations

### Startup Time Impact

```go
// Keep init functions lightweight
func init() {
	// Good: Fast operations
	config = getDefaultConfig()

	// Avoid: Slow operations that delay startup
	// downloadLargeFile()
	// connectToSlowService()
}
```

### Memory Usage

```go
// Be mindful of memory allocation in init
func init() {
	// Consider lazy initialization for large objects
	largeCache = make(map[string][]byte) // OK if needed immediately

	// Or use sync.Once for lazy initialization
	var once sync.Once
	getLargeCache := func() map[string][]byte {
		once.Do(func() {
			largeCache = initializeLargeCache()
		})
		return largeCache
	}
}
```

## Key Takeaways

1. **Automatic Execution**: Init functions run automatically before main, cannot be called manually
2. **No Signature**: Cannot have parameters or return values
3. **Multiple Allowed**: Packages can have multiple init functions
4. **Execution Order**: Run in dependency order across packages, source order within packages
5. **One-Time Execution**: Each init function runs exactly once per program execution
6. **Initialization Purpose**: Best used for setup tasks, not business logic
7. **Error Handling**: Use log.Fatal or panic for critical initialization failures
8. **Performance Impact**: Keep init functions fast to avoid delaying program startup

## Next Steps

- Learn about [anonymous functions](../c.%20anonymous%20function/) to understand functions without names
- Study [higher-order functions](../d.%20higher%20order%20function/) for advanced function composition
- Explore [method functions](../e.%20method%20function/) for associating functions with types
- Investigate [recursive functions](../f.%20recursive%20function/) for self-calling function patterns
