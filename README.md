# Unexpected Go Memory Address Behavior

This repository demonstrates a common misunderstanding regarding memory addresses in Go.  The example shows that while you can get the memory address of a variable using the `&` operator, these addresses are not guaranteed to be stable or directly comparable across different variables.  They can change between different runs and even during a single program's execution depending on memory allocation and garbage collection.

## Key Points

* The memory address printed is specific to the current execution environment. 
* Comparing memory addresses directly is usually not meaningful or reliable.
* Focus on the values of variables, not their memory locations for program logic.