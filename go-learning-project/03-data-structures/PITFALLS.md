# PITFALLS

## Go Struct Pitfalls (Bug-Oriented)

### Easy Pitfalls

* **Value receiver used when mutation is expected**
  Method compiles, but changes don’t persist outside the method.

* **Pointer receiver on a nil pointer**
  Calling a method panics when the receiver is dereferenced.

* **Forgetting to initialize embedded or inner struct fields**
  Zero values lead to incorrect behavior or output.

* **Zero-value struct treated as valid data**
  Missing sentinel or validation checks (e.g., empty ID used as real).

* **Field shadowing with local variables**
  Using `:=` instead of assigning to the struct field.

* **Wrong field selected due to similar names**
  Example: `Total` vs `Subtotal`.

* **Comparing structs containing non-comparable fields**
  Compile errors or logic removed to “fix” the build.

* **Incorrect use of `new(T)` vs `&T{}`**
  Required fields remain uninitialized.

* **Unintentional struct copies via assignment**
  State not shared when it was expected to be.

* **Taking the address of a loop variable**
  All pointers end up referencing the same value.

* **Returning a pointer to a loop variable**
  Common in parsers/builders; produces identical results.

* **JSON marshaling ignores unexported fields**
  Fields silently missing from output.

* **Unexported fields cannot be set from other packages**
  Leads to partially initialized structs.

---

### Medium Pitfalls

* **Method set mismatch with interfaces**
  Value vs pointer receivers cause “almost implements interface” bugs.

* **Embedding causes promoted field/method ambiguity**
  The field or method used is not the one intended.

* **Nil pointer to embedded struct**
  Promoted method calls panic.

* **Mutating slice/map fields via copied structs**
  Some changes persist, others don’t—confusing aliasing.

* **Shallow copy of structs with reference fields**
  Internal state is unintentionally shared.

* **`append` on slice field doesn’t persist**
  Slice reallocation + struct copy drops updates.

* **Using `==` for semantic equality**
  Field-by-field equality isn’t business equality.

* **Direct comparison of `time.Time` or floats**
  Causes flaky behavior and incorrect branching.

* **Exported fields bypass invariants**
  Callers mutate state without validation.

* **`omitempty` hides required fields**
  Bugs surface downstream after marshaling.

* **Passing structs by value through channels**
  Receivers see stale copies.

* **Modifying struct elements in `range` loops**
  Changes don’t persist without indexing or pointers.

* **Copying structs containing `sync.Mutex`**
  Locks silently stop protecting shared state.

* **Initialization order issues with embedded structs**
  Later assignments overwrite promoted fields.

---

### Hard Pitfalls

* **Copying structs containing synchronization primitives**
  (`sync.Mutex`, `sync.RWMutex`, `sync.Once`)
  Leads to races, deadlocks, or undefined behavior.

* **Returning pointers to slice-backed data**
  Later reallocation invalidates assumptions.

* **Broken invariants under concurrency**
  Multiple fields updated without atomicity.

* **False sharing due to struct layout**
  Severe performance degradation under load.

* **`unsafe.Pointer` with struct layout assumptions**
  Breaks across architectures or Go versions.

* **Reflection cannot set unaddressable fields**
  `CanSet()` is false; mutations silently fail.

* **JSON + embedded structs + name collisions**
  Fields override each other unexpectedly.

* **Interface holding a typed nil pointer**
  Interface is non-nil, but method calls panic.

* **Structs with pointer fields used as map keys**
  Key identity changes; lookups fail.

* **Data races from copying structs with shared pointers**
  Multiple owners mutate the same memory.

* **Capturing iteration variables in goroutines**
  All goroutines observe the same struct.

* **Custom `UnmarshalJSON` with value receiver**
  Decoding doesn’t modify the target.

* **Using `sync.Pool` with structs holding references**
  Memory retention, stale state, subtle leaks.

