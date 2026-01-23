| Question | Answer |
|----------|--------|
| What is a Go slice and what three components does it contain? | A slice is a dynamic array backed by a fixed-size array. It contains:<br>1. Pointer to the first element<br>2. Length (number of elements)<br>3. Capacity (total space available) |
| How do you create a slice in Go using literal syntax? | Using square brackets with initial values: {{c1::`s := []int{1, 2, 3}`}} |
| What is the difference between length and capacity of a slice? | Length is the number of elements currently in the slice (accessed with `len()`). Capacity is the total space allocated for the underlying array (accessed with `cap()`). |
| What does `make([]int, 5)` create? | A slice of integers with length 5, capacity 5, with all elements initialized to zero value (0) |
| What does `make([]int, 5, 10)` create? | A slice of integers with length 5 and capacity 10. You can append up to 5 more elements before reallocation. |
| When you pass a slice to a function, what gets passed by value? | The slice header (pointer, length, capacity) is copied. The underlying array is {{c1::NOT}} copied - both slices point to the same array. |
| What happens when you call `append()` on a slice that is at capacity? | Go allocates a new, larger underlying array, copies elements to it, and returns a new slice header pointing to the new array. |
| Why doesn't modifying a slice inside a function affect the caller's slice when using `append()`? | Because `append()` may allocate a new array, returning a new slice header. The function receives a copy of the original header, so the assignment `s = append(s, ...)` doesn't affect the caller's variable. |
| How do you modify a slice passed to a function and ensure changes are visible to the caller? | Return the modified slice: {{c1::`s = properModify(s)`}} instead of modifying it in-place in a function |
| What is the syntax for appending one slice to another? | Use the spread operator `...`: {{c1::`result := append(s1, s2...)`}} |
| What is reslicing in Go? | Creating a new slice view of an existing slice using the syntax {{c1::`s[start:end]`}}. The new slice shares the same underlying array. |
| What does `s[2:4]` create? | A new slice containing elements at indices 2 and 3 (start inclusive, end exclusive). Length is 2, but capacity depends on remaining space in the original array. |
| What is the danger of reslicing when combined with `append()`? | If the resliced slice has available capacity in the shared array, `append()` may overwrite elements in the original slice that you didn't intend to modify. |
| How do you safely reslice to prevent `append()` from affecting the original slice? | Use a full slice expression: {{c1::`y := s[low:high:high]`}}. This sets the capacity to `high - low`, preventing append from accessing the original array. |
| If you do `s2 := append(s1, 4)` and both s1 and s2 have capacity for the element, what happens to s1? | If s1 has capacity, s2 points to the same underlying array. Modifying s2 will modify the shared array, affecting s1's elements even though s1's length doesn't change. |
| What does `cap()` return for an empty slice declared as `var s []int`? | 0 - an empty slice has zero capacity |
| What does the `clear()` function do to a slice? | Sets all elements in the slice to their zero values without changing the length: {{c1::`clear(s)` makes all elements 0 (for int) or empty string (for string), etc.}} |
| When iterating with `for i, v := range slice`, what is important to remember about `v`? | `v` is a copy of each element, not a reference. Modifying `v` does not modify the original slice elements. |
| What is the correct way to iterate and modify slice elements in place? | Use the index: {{c1::`for i := range slice { slice[i] = newValue }`}} |
| What is the two-value form of `range` over a slice? | {{c1::`for index, value := range slice {}`}} - index starts at 0, value is a copy of the element |
| What is the one-value form of `range` over a slice? | {{c1::`for i := range slice {}`}} - only iterates over indices. In Go 1.22+, this iterates over indices not values. |
| What happens when you do `append()` repeatedly in a loop without pre-allocating? | The slice may reallocate multiple times, creating new underlying arrays and copying data repeatedly. This is inefficient. |
| How do you avoid repeated reallocations when you know the final slice size? | Use `make()` to pre-allocate with the expected capacity: {{c1::`s := make([]int, 0, 100)`}} |
| What is unsafe.SliceData() used for? | It returns a pointer to the first element of the slice's underlying array: {{c1::`ptr := unsafe.SliceData(s)`}} |
| What happens to the underlying array address when `append()` causes a reallocation? | The underlying array address changes because Go allocates a new, larger array and the slice now points to it. |
| How do you create a slice of structs with literal syntax? | Use struct literals within slice brackets: {{c1::`s := []Item{{ID: 1, Ready: true}, {ID: 2, Ready: false}}`}} |
| Can you iterate over a slice using only the values without the index? | Yes: {{c1::`for v := range slice {}`}} - only iterates over values (the variable name represents the value, not index) |
| What does `append(s, 4)` do if you don't assign the result? | In Go, `append()` returns a new slice. If you don't assign it (like `_ = append(s, 4)`), the append is discarded and the original slice `s` is unchanged. |
| Why should you assign the result of `append()` back to the variable? | Because append returns a potentially new slice header. To capture the changes: {{c1::`s = append(s, element)`}} |
| What is the difference between `[]int` and `[5]int`? | {{c1::`[]int`}} is a slice (dynamic), `[5]int` is an array (fixed size of 5). Arrays cannot grow; slices can. |
| When you pass an array to a function, what gets passed? | The entire array is copied by value (all elements). This is inefficient for large arrays. |
| When you pass a slice to a function, what gets passed? | Only the slice header (pointer, length, capacity) is copied. The underlying array is shared. |
| What is the zero value of a slice? | `nil` - an uninitialized slice with length 0 and capacity 0 |
| Is a nil slice equal to an empty slice? | No: `nil` is an uninitialized slice, while an empty slice `[]int{}` or `make([]int, 0)` has been explicitly initialized but has no elements. |
| What's the result of iterating over a nil slice? | The loop runs zero times - no iteration occurs |
| What does `s[i:j]` evaluate to if i equals j? | An empty slice with length 0 but same capacity as if it pointed to the shared array at position j |
| What are the constraints for a slice expression `s[low:high]`? | {{c1::0 ≤ low ≤ high ≤ len(s)}} - low must be ≤ high, and both must be within the slice's length |
| What does `s[:]` do? | Creates a new slice header pointing to the entire slice (equivalent to `s[0:len(s)]`) |
| What does `s[2:]` do? | Creates a slice from index 2 to the end |
| What does `s[:3]` do? | Creates a slice from the beginning to index 3 (exclusive), so elements at indices 0, 1, 2 |
| When you use the full slice expression `s[low:high:max]`, what must be true? | {{c1::low ≤ high ≤ max ≤ len(s)}} - the new capacity will be `max - low` |
| What is the purpose of using a full slice expression like `s[1:3:3]`? | To limit the capacity to prevent future appends from accessing beyond the intended slice range |
| What happens if `append()` needs more capacity than the current allocation? | Go typically doubles the capacity (or uses a growth factor) to amortize reallocations |
| What is the growth strategy for slice capacity in Go? | Go uses a threshold: for small slices (< 256 elements), capacity is roughly doubled; for larger slices, it grows by ~25% |
| If you need a slice of slices (2D), how do you create it? | {{c1::`s := [][]int{{1, 2}, {3, 4}}`}} or allocate and initialize each row separately |
| How do you print a slice of strings nicely? | Use `fmt.Println()` or `strings.Join()`: {{c1::`fmt.Println(strings.Join(slice, " "))`}} |
| What is shallow copy in the context of slices? | When you assign a slice to another: `b := a` - both point to the same underlying array. It's a shallow copy of the header. |
| How do you make a deep copy of a slice? | Use `copy()` or create a new slice and populate it: {{c1::`dst := make([]int, len(src)); copy(dst, src)`}} |
| What does `copy(dst, src)` do? | Copies elements from src to dst. Returns the number of elements copied (minimum of `len(dst)` and `len(src)`). |
| What is the risk of using the same slice as both source and destination in `copy()`? | If ranges overlap and copy happens element-by-element, you may get unexpected results. Generally safe if dst comes before src. |
| When should you use `make()` vs. slice literals? | Use `make()` for dynamic allocation with size/capacity; use literals like `[]int{1,2,3}` for known, fixed initial values. |
| What is a common mistake when using `append()` in nested loops? | Not realizing that append may reallocate, causing unexpected behavior if you thought you were modifying a shared array |
| How do you check if a slice is empty? | Use `len(s) == 0` or simply `if len(s) == 0` |
| What does appending to a `nil` slice do? | It works correctly - Go allocates a new underlying array and returns a new slice with the element |
| Can you compare two slices for equality using `==`? | No - slices cannot be compared with `==`. Use `reflect.DeepEqual()` or iterate and compare elements manually. |
| What happens when you reslice a slice to have length equal to its capacity? | The slice now contains all elements up to its capacity position. Further appends will require reallocation. |
| What is the memory layout of a slice header? | A pointer to the underlying array, a length field, and a capacity field (24 bytes on 64-bit systems) |
| When are pointers to slice elements invalidated? | When the underlying array is reallocated during `append()` - old pointers become invalid |
| How do you safely get a pointer to a slice element? | Use `unsafe.Pointer(&s[i])` or `unsafe.SliceData(s)`, but be aware it becomes invalid after reallocation |
| What is the relationship between array and slice in Go? | A slice is a view into an array. Arrays are fixed-size; slices are dynamic views backed by arrays. |
| Can you convert an array to a slice? | Yes: `a := [5]int{1,2,3,4,5}; s := a[:]` - creates a slice pointing to the entire array |
| Can you convert a slice back to an array? | Partial: `a := [5]int(slice)` only works if the slice has exactly 5 elements. Usually use copy instead. |
