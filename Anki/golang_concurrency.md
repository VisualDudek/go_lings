# Go Concurrency Flashcards

| Question | Answer |
|----------|--------|
| What keyword do you use to spawn a goroutine in Go? | The `go` keyword<br><br>Example: `go say("world")` |
| What happens to goroutines when the main function finishes? | They are terminated immediately, even if not complete.<br><br>The main function does not wait for spawned goroutines. |
| How do you create an unbuffered channel in Go? | Using `make(chan Type)`<br><br>Example: `c := make(chan int)` |
| What is the key behavior of an unbuffered channel? | It blocks until both sender and receiver are ready.<br><br>Send blocks until someone receives; receive blocks until someone sends. |
| What is the difference between buffered and unbuffered channels? | Buffered channels: `make(chan Type, capacity)` – sends block only when full, receives block when empty<br><br>Unbuffered channels: sends and receives always block until the other side is ready |
| How do you send a value to a channel? | Use the send operator `<-`<br><br>Example: `c <- value`<br><br>The channel is on the right side of the arrow. |
| How do you receive a value from a channel? | Use the receive operator `<-`<br><br>Example: `value := <-c`<br><br>The channel is on the right side of the arrow. |
| What does this code do?<br><br>`x, y := <-c, <-c` | It receives two values from channel `c` concurrently, blocking until both are available.<br><br>Both receives happen in parallel before continuing. |
| Can you receive from a buffered channel that is empty? | No, it blocks until a value is sent.<br><br>Even when buffered, receiving from an empty channel blocks the goroutine. |
| What error occurs if you fill a buffered channel without draining it? | Deadlock – if no goroutine receives from the channel, the sender will block forever. |
| How do you close a channel in Go? | Use the `close()` function<br><br>Example: `close(c)`<br><br>Only the sender should close a channel. |
| What happens when you receive from a closed channel? | It returns the zero value immediately without blocking.<br><br>Multiple receives from a closed channel return zero values. |
| How do you iterate over channel values until it closes? | Use `for` with `range`<br><br>Example: `for value := range c { ... }`<br><br>The loop exits automatically when the channel is closed. |
| What is the purpose of the `select` statement in Go? | It waits on multiple channel operations and executes the first one that is ready.<br><br>Useful for coordinating communication between multiple channels. |
| What happens if multiple cases in `select` are ready at the same time? | Go randomly selects one of the ready cases to execute.<br><br>This prevents bias and avoids starvation. |
| What does the `default` case do in a `select` statement? | It executes immediately if no other cases are ready (non-blocking behavior). |
| Is there a deadlock risk with an empty `select` statement? | Yes, `select {}` blocks forever with no cases ready.<br><br>Use carefully or pair with cases to avoid deadlock. |
| What does `time.Tick()` return? | A channel that sends the current time at regular intervals.<br><br>Example: `tick := time.Tick(100 * time.Millisecond)`<br><br>The channel never closes. |
| What does `time.After()` return? | A channel that sends the current time once after a specified duration.<br><br>Example: `boom := time.After(500 * time.Millisecond)`<br><br>Useful for timeouts. |
| For this code, what is printed?<br><br>`ch := make(chan int, 2)`<br>`ch <- 1`<br>`ch <- 2`<br>`fmt.Println(<-ch)`<br>`fmt.Println(<-ch)` | Output:<br>1<br>2<br><br>The buffered channel holds 2 values without blocking. |
| For this code, does it deadlock? Why or why not?<br><br>`ch := make(chan int, 2)`<br>`ch <- 1`<br>`ch <- 2`<br>`ch <- 3` | Yes, it deadlocks. The buffer capacity is 2, but we try to send 3 values without receiving any. |
| What is "channel ownership" in Go concurrency? | The sender owns the channel and is responsible for closing it.<br><br>The receiver should never close a channel because sending on a closed channel causes a panic. |
| What is an idiomatic pattern for closing a channel from a goroutine? | Use a goroutine with `defer close(ch)` to ensure the channel is always closed when the function returns.<br><br>Example:<br>`go func() {`<br>`    defer close(ch)`<br>`    // send values to ch`<br>`}()` |
| What is "in-order tree traversal"? | Visit nodes in ascending order: left subtree, then node, then right subtree.<br><br>For a binary search tree, this produces sorted values. |
| How can you compare two trees concurrently using channels? | Create channels for each tree, spawn goroutines to walk both trees concurrently, and compare received values side-by-side.<br><br>If any value differs, the trees are not the same. |
| What is a data race in Go? | When two or more goroutines access the same variable concurrently and at least one modifies it without synchronization.<br><br>This causes unpredictable behavior and must be prevented. |
| What does `sync.Mutex` do? | It provides mutual exclusion – only one goroutine can hold the lock at a time.<br><br>Prevents concurrent access to shared data. |
| How do you use a mutex to protect a variable? | Lock before accessing, unlock after (use `defer` for safety):<br><br>`mu.Lock()`<br>`defer mu.Unlock()`<br>`// access shared data` |
| In this code, is the map access safe?<br><br>`type SafeCounter struct {`<br>`    mu sync.Mutex`<br>`    v  map[string]int`<br>`}`<br><br>`func (c *SafeCounter) Inc(key string) {`<br>`    c.mu.Lock()`<br>`    c.v[key]++`<br>`    c.mu.Unlock()`<br>`}` | Yes, the mutex protects all access to the map `c.v`. |
| What is a race condition? | A situation where the output depends on the timing of goroutines rather than logic.<br><br>Occurs when multiple goroutines access shared data without synchronization. |
| Can you send on a nil channel? | No, it blocks forever (deadlock).<br><br>Attempting to send on a nil channel causes a panic in a select statement. |
| Can you receive from a nil channel? | No, it blocks forever (deadlock).<br><br>Receiving from a nil channel blocks indefinitely. |
| What is an anonymous goroutine? | A goroutine created from an anonymous function (closure).<br><br>Example: `go func() { ... }()`<br><br>Useful for one-off concurrent tasks. |
| When should you use channels vs mutexes? | Use channels for communicating values between goroutines.<br><br>Use mutexes to protect shared data structures when goroutines don't exchange values. |
| What happens if you try to close a nil channel? | It causes a panic: "close of nil channel". |
| What is the behavior of `cap()` on a channel? | It returns the buffer capacity of the channel.<br><br>For unbuffered channels, it returns 0. |
| What does this select statement do?<br><br>`select {`<br>`case c <- x:`<br>`    x, y = y, x+y`<br>`case <-quit:`<br>`    return`<br>`}` | It tries to send x to channel c; if it blocks, it waits for a receive on quit and returns if received. |
| Why is defer used with mutex unlock? | To ensure the mutex is unlocked even if a panic occurs between Lock and Unlock.<br><br>Example: `c.mu.Lock()` / `defer c.mu.Unlock()` |
| What is the difference between `channel <-` and `<- channel` syntax? | `channel <-` is a send (value goes to channel).<br><br>`<- channel` is a receive (value comes from channel). |
| For this select, what happens?<br><br>`select {`<br>`case <-tick:`<br>`    fmt.Println("tick")`<br>`case <-boom:`<br>`    fmt.Println("BOOM")`<br>`default:`<br>`    fmt.Println("waiting")`<br>`}` | If tick and boom are not ready, the default case executes (non-blocking).<br><br>Otherwise, whichever is ready first executes. |
| What is a goroutine leak? | When goroutines continue running after they are no longer needed, wasting resources.<br><br>Example: spawning goroutines that receive from a channel that is never sent to. |
| How do you prevent goroutine leaks with channels? | Always close the channel on the sender side or provide a way for receivers to exit.<br><br>Use channels in select with quit signals or timeouts. |
| What does `time.Sleep()` do in a select statement? | It cannot be used directly; select does not support sleep cases.<br><br>Use `time.After()` or `time.Tick()` instead for time-based channel operations. |
| What is the output of this code?<br><br>`c := make(chan int)`<br>`go func() {`<br>`    c <- 5`<br>`}()`<br>`x := <-c`<br>`fmt.Println(x)` | Output: 5<br><br>The goroutine sends 5, synchronizing with the main goroutine's receive. |
| What channels does `time.Tick()` vs `time.After()` represent? | `time.Tick()` is a repeating channel (infinite sends at intervals).<br><br>`time.After()` is a one-shot channel (one send after delay). |
| In a concurrent tree walk, why close the channel? | Closing signals the receiver that no more values will be sent, allowing `for range` loops to exit cleanly. |
| What is the difference between buffered and unbuffered in terms of goroutine coupling? | Unbuffered: tight coupling – sender and receiver must rendezvous.<br><br>Buffered: loose coupling – sender can proceed if buffer has space. |
