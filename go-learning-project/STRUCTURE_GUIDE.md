# Detailed Directory Structure Guide

## 01-fundamentals/
**Goal**: Establish Go syntax and mental model

```
01-fundamentals/
├── README.md                 # Go syntax overview
├── variables_constants.go    # var, const, type declarations
├── data_types.go            # int, float, string, bool
├── operators.go             # arithmetic, logical, comparison
├── control_flow.go          # if/else, switch, for loops
├── exercises/
│   ├── ex01_hello_world.go
│   ├── ex02_temperature_converter.go
│   ├── ex03_fibonacci.go
│   └── solutions/
└── _test.go                 # Basic unit test examples
```

**Key Concepts**: Variables, constants, primitive types, control structures
**Exercises**: Simple calculators, temperature converters, sequence generators

---

## 02-functions/
**Goal**: Master function design and Go's functional capabilities

```
02-functions/
├── README.md                # Function patterns in Go
├── basic_functions.go       # Parameters, return values, multiple returns
├── variadic_functions.go    # ...args pattern
├── named_returns.go         # Named return values (and when to avoid)
├── closures.go              # Closures and lexical scoping
├── defer_panic_recover.go   # Resource management patterns
├── higher_order.go          # Functions as values, callbacks
├── exercises/
│   ├── ex01_function_signatures.go
│   ├── ex02_calculator_refactor.go
│   └── ex03_pipeline.go
└── _test.go
```

**Key Concepts**: Parameter passing, multiple returns, defer, panic/recover, first-class functions
**Exercises**: Custom middleware, functional pipelines, resource cleanup patterns

---

## 03-data-structures/
**Goal**: Work with Go's fundamental collection types

```
03-data-structures/
├── README.md                # Slices vs arrays, maps overview
├── arrays.go                # Fixed-size arrays, bounds
├── slices.go                # Dynamic arrays, append, reslicing
├── slice_internals.go       # Slice headers, capacity, backing arrays
├── maps.go                  # Key-value storage, iteration, deletion
├── structs.go               # Struct definition, embedding, methods
├── struct_tags.go           # Field tags, reflection patterns
├── exercises/
│   ├── ex01_slice_operations.go
│   ├── ex02_word_frequency.go
│   ├── ex03_student_records.go
│   └── ex04_json_serialization.go
└── _test.go
```

**Key Concepts**: Arrays, slices, maps, structs, struct composition, tags
**Exercises**: Data manipulation, CSV processing, struct validation

---

## 04-interfaces/
**Goal**: Understand Go's compositional approach to abstraction

```
04-interfaces/
├── README.md                # Interface design philosophy
├── basic_interfaces.go      # Interface definition and implementation
├── interface_satisfaction.go # Implicit satisfaction, type assertions
├── empty_interface.go       # interface{}, type switches
├── reader_writer.go         # io.Reader, io.Writer patterns
├── composition.go           # Embedding interfaces
├── exercises/
│   ├── ex01_shape_interface.go
│   ├── ex02_reader_writer.go
│   ├── ex03_plugin_system.go
│   └── ex04_strategy_pattern.go
└── _test.go
```

**Key Concepts**: Interface design, implicit satisfaction, type assertions, standard library patterns
**Exercises**: Building extensible systems, implementing standard interfaces

---

## 05-concurrency/
**Goal**: Master goroutines and channels

```
05-concurrency/
├── README.md                # Concurrency model overview
├── goroutines_basics.go     # Creating goroutines, go keyword
├── channels_basics.go       # Unbuffered channels, send/receive
├── buffered_channels.go     # Channel buffering, capacity
├── select_statement.go      # Multiplexing, timeouts
├── channel_patterns.go      # Fan-in, fan-out, pipelines
├── exercises/
│   ├── ex01_goroutine_spawning.go
│   ├── ex02_worker_pool.go
│   ├── ex03_data_pipeline.go
│   └── ex04_rate_limiter.go
├── benchmarks/
│   └── concurrent_vs_sequential.go
└── _test.go
```

**Key Concepts**: Goroutines, channels, select, deadlock prevention, channel patterns
**Exercises**: Concurrent workers, producer-consumer, pipeline patterns

---

## 06-error-handling/
**Goal**: Go's error handling approach and best practices

```
06-error-handling/
├── README.md                # Error philosophy in Go
├── error_interface.go       # error type, implementing errors
├── multiple_returns.go      # Error as second return value
├── custom_errors.go         # Custom error types, assertions
├── error_wrapping.go        # errors.Is, errors.As, fmt.Errorf %w
├── context_errors.go        # Context cancellation, error propagation
├── exercises/
│   ├── ex01_custom_error_types.go
│   ├── ex02_error_wrapping.go
│   ├── ex03_robust_parsing.go
│   └── ex04_context_timeouts.go
└── _test.go
```

**Key Concepts**: Error interface, custom errors, error wrapping, sentinel errors
**Exercises**: Robust argument validation, graceful degradation

---

## 07-packages-modules/
**Goal**: Organize code professionally with packages and modules

```
07-packages-modules/
├── README.md                # Package design, module management
├── package_basics/
│   ├── main.go
│   └── helpers.go
├── internal_packages/
│   ├── cmd/app/main.go
│   └── internal/config/config.go
├── module_example/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── vendor_example/
│   ├── go.mod
│   └── vendor/
├── exercises/
│   ├── ex01_multi_package_app.go
│   └── ex02_module_publishing.go
└── README.md
```

**Key Concepts**: Package organization, module management, internal packages, version control
**Exercises**: Refactoring monolith into packages, creating shareable modules

---

## 08-io-files/
**Goal**: Handle I/O operations efficiently

```
08-io-files/
├── README.md                # Reader/Writer interfaces, buffering
├── file_operations.go       # Open, read, write, close
├── buffered_io.go           # bufio package, buffering strategies
├── readers_writers.go       # io.Reader, io.Writer abstractions
├── encoders_decoders.go     # JSON, CSV, binary encoding
├── streaming.go             # Processing large files without buffering
├── exercises/
│   ├── ex01_file_copy.go
│   ├── ex02_csv_processor.go
│   ├── ex03_log_parser.go
│   └── ex04_streaming_json.go
├── data/
│   ├── sample.csv
│   └── sample.json
└── _test.go
```

**Key Concepts**: File I/O, buffering, encoding/decoding, streaming patterns
**Exercises**: Log processing, data transformation, format conversion

---

## 09-testing/
**Goal**: Write comprehensive tests

```
09-testing/
├── README.md                # Testing philosophy and tools
├── unit_tests.go            # Basic *testing.T patterns
├── unit_tests_test.go
├── table_driven_tests.go    # Data-driven testing
├── table_driven_tests_test.go
├── benchmarks_test.go       # Performance testing, comparing implementations
├── subtests.go              # t.Run, nested test organization
├── subtests_test.go
├── fixtures.go              # Test data, setup/teardown
├── mocking.go               # Interface-based mocking
├── mocking_test.go
├── exercises/
│   ├── ex01_write_unit_tests.go
│   ├── ex02_improve_coverage.go
│   └── ex03_benchmark_comparison.go
└── coverage/
    └── coverage_report.html
```

**Key Concepts**: Unit testing, table-driven tests, benchmarks, mocking, coverage
**Exercises**: Achieving high coverage, testing concurrent code, performance analysis

---

## 10-web-basics/
**Goal**: Build HTTP applications

```
10-web-basics/
├── README.md                # net/http fundamentals
├── http_server.go           # Basic server, handlers
├── routing.go               # Route patterns, multiplexing
├── middleware.go            # Middleware chains, logging
├── request_response.go      # Reading requests, writing responses
├── json_api.go              # JSON encoding/decoding
├── exercises/
│   ├── ex01_simple_server.go
│   ├── ex02_rest_api.go
│   ├── ex03_middleware_chain.go
│   └── ex04_api_versioning.go
├── tests/
│   └── integration_test.go
└── main.go
```

**Key Concepts**: HTTP handlers, routing, middleware, JSON APIs, testing HTTP servers
**Exercises**: Building REST APIs, authentication middleware, CORS handling

---

## 11-databases/
**Goal**: Work with databases

```
11-databases/
├── README.md                # Database patterns and practices
├── sql_basics.go            # database/sql package, queries
├── transactions.go          # ACID properties, rollback patterns
├── prepared_statements.go   # Parameter binding, SQL injection prevention
├── orm_patterns.go          # ORM concepts without heavy library
├── migration_patterns.go    # Schema versioning
├── exercises/
│   ├── ex01_simple_crud.go
│   ├── ex02_transaction_handling.go
│   ├── ex03_migration_system.go
│   └── ex04_query_optimization.go
├── migrations/
│   ├── 001_create_tables.sql
│   └── 002_add_indexes.sql
├── schema.sql
└── tests/
    └── db_test.go
```

**Key Concepts**: SQL queries, transactions, connection pooling, prepared statements
**Exercises**: Building data access layer, implementing migrations, query optimization

---

## 12-reflection/
**Goal**: Inspect and manipulate types dynamically

```
12-reflection/
├── README.md                # reflect package, introspection
├── reflect_basics.go        # Type, Value, Kind
├── struct_introspection.go  # Inspecting struct fields, tags
├── dynamic_calls.go         # Calling functions dynamically
├── custom_unmarshaling.go   # Implementing custom serialization
├── exercises/
│   ├── ex01_type_inspection.go
│   ├── ex02_struct_copier.go
│   ├── ex03_json_encoder.go
│   └── ex04_dependency_injection.go
└── _test.go
```

**Key Concepts**: Type inspection, value reflection, dynamic invocation
**Exercises**: Building struct mappers, custom marshaling, simple DI container

---

## 13-advanced-concurrency/
**Goal**: Advanced concurrent patterns and synchronization

```
13-advanced-concurrency/
├── README.md                # Advanced patterns and primitives
├── sync_primitives.go       # Mutex, RWMutex, WaitGroup, Once
├── sync_primitives_test.go
├── context_package.go       # Context hierarchy, cancellation, timeouts
├── context_test.go
├── semaphore_pattern.go     # Resource limiting patterns
├── worker_pool.go           # Thread pool implementations
├── worker_pool_test.go
├── bounded_queue.go         # Backpressure and flow control
├── lock_free.go             # sync/atomic patterns
├── exercises/
│   ├── ex01_concurrent_map.go
│   ├── ex02_rate_limiter.go
│   ├── ex03_circuit_breaker.go
│   ├── ex04_fan_in_fan_out.go
│   └── ex05_graceful_shutdown.go
└── benchmarks/
    └── sync_strategies_bench_test.go
```

**Key Concepts**: Mutex, channels vs mutexes, context propagation, atomic operations
**Exercises**: Building thread-safe data structures, implementing worker pools, graceful shutdown

---

## 14-systems-programming/
**Goal**: Unix/Linux system programming concepts

```
14-systems-programming/
├── README.md                # OS concepts, syscalls
├── signal_handling.go       # os.Signal, signal.Notify
├── process_management.go    # Spawning processes, os/exec
├── unix_sockets.go          # Unix domain sockets, IPC
├── file_permissions.go      # Chmod, file modes, umask
├── environment_variables.go # os.Environ, configuration
├── exercises/
│   ├── ex01_signal_handler.go
│   ├── ex02_process_launcher.go
│   ├── ex03_unix_socket_server.go
│   ├── ex04_daemon_process.go
│   └── ex05_system_monitor.go
└── _test.go
```

**Key Concepts**: Signals, process spawning, IPC mechanisms, file permissions
**Exercises**: Building servers with graceful shutdown, process spawning, system utilities

---

## 15-performance/
**Goal**: Optimize and profile Go applications

```
15-performance/
├── README.md                # Profiling tools and optimization
├── profiling_basics.go      # pprof, CPU profiling
├── memory_profiling.go      # Heap profiling, allocations
├── benchmarking.go          # Writing meaningful benchmarks
├── benchmarks_test.go
├── optimization_techniques.go # String building, slice preallocation
├── gc_tuning.go             # GOGC, GC behavior
├── exercises/
│   ├── ex01_cpu_profile.go
│   ├── ex02_memory_leaks.go
│   ├── ex03_allocation_analysis.go
│   └── ex04_optimize_hot_path.go
├── profiles/                # Generated profile outputs
│   └── cpu.prof
└── scripts/
    └── profile_analysis.sh
```

**Key Concepts**: CPU profiling, memory profiling, benchmarking, GC tuning
**Exercises**: Identifying bottlenecks, reducing allocations, memory leak detection

---

## 16-distributed-systems/
**Goal**: Build distributed applications

```
16-distributed-systems/
├── README.md                # Distributed patterns
├── rpc_basics.go            # net/rpc package
├── grpc_basics.go           # Protocol Buffers, gRPC basics
├── grpc_advanced.go         # Streaming, interceptors
├── load_balancing.go        # Client-side load balancing
├── service_discovery.go     # Service registration patterns
├── circuit_breaker.go       # Resilience patterns
├── exercises/
│   ├── ex01_basic_rpc.go
│   ├── ex02_grpc_service.go
│   ├── ex03_streaming_rpc.go
│   ├── ex04_load_balancer.go
│   └── ex05_resilience.go
├── proto/
│   └── service.proto
├── go.mod
└── tests/
    └── distributed_test.go
```

**Key Concepts**: RPC protocols, gRPC, service discovery, resilience patterns
**Exercises**: Building microservices, implementing load balancing, circuit breaker pattern

---

## 17-real-world-projects/
**Goal**: Integrate all concepts into complete applications

```
17-real-world-projects/
├── project-1-cli-tool/
│   ├── main.go
│   ├── cmd/
│   │   ├── list.go
│   │   ├── add.go
│   │   └── delete.go
│   ├── internal/
│   │   ├── storage/
│   │   └── cli/
│   ├── go.mod
│   └── README.md
│
├── project-2-web-api/
│   ├── main.go
│   ├── internal/
│   │   ├── api/handlers.go
│   │   ├── db/queries.go
│   │   └── middleware/auth.go
│   ├── migrations/
│   ├── tests/
│   ├── go.mod
│   └── README.md
│
├── project-3-system-monitor/
│   ├── main.go
│   ├── internal/
│   │   ├── metrics/collector.go
│   │   └── output/formatter.go
│   ├── Dockerfile
│   ├── go.mod
│   └── README.md
│
├── project-4-distributed-cache/
│   ├── main.go
│   ├── cmd/
│   │   ├── server.go
│   │   └── client.go
│   ├── proto/cache.proto
│   ├── internal/
│   │   ├── cache/lru.go
│   │   ├── replication/
│   │   └── rpc/
│   ├── go.mod
│   └── README.md
│
└── project-5-task-scheduler/
    ├── main.go
    ├── internal/
    │   ├── scheduler/executor.go
    │   ├── storage/persistence.go
    │   └── api/server.go
    ├── migrations/
    ├── docker-compose.yml
    ├── go.mod
    └── README.md
```

**Goal**: Each project reinforces multiple concepts:

- **Project 1 (CLI Tool)**: Package organization, flag parsing, file I/O, error handling
- **Project 2 (Web API)**: HTTP handling, databases, middleware, testing, error handling
- **Project 3 (System Monitor)**: Signal handling, concurrency, metrics collection, output formatting
- **Project 4 (Distributed Cache)**: gRPC, replication, concurrency, performance optimization
- **Project 5 (Task Scheduler)**: Concurrency, persistence, HTTP APIs, graceful shutdown

---

## Progression Summary

| Phase | Directories | Focus |
|-------|-----------|-------|
| **Fundamentals** | 01-03 | Language basics |
| **Core Features** | 04-07 | Go-specific capabilities |
| **Practical Dev** | 08-12 | Real application patterns |
| **Advanced** | 13-17 | Performance, distribution, systems programming |

## Time Estimates (per directory)

- **Fundamentals (01-03)**: 2-3 weeks
- **Core Features (04-07)**: 3-4 weeks
- **Practical Dev (08-12)**: 4-5 weeks
- **Advanced (13-17)**: 6-8 weeks
- **Projects**: 4-6 weeks

**Total**: ~4-5 months of consistent learning

