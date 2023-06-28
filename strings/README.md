In Go, the `strings` package provides several useful functions for manipulating strings. It is part of the standard library, so you don't need to install any additional packages to use it. Here are some commonly used functions from the `strings` package:

1. `strings.Contains(s, substr string) bool`: Checks if a string `s` contains a specific substring `substr` and returns a boolean value indicating the result.

2. `strings.HasPrefix(s, prefix string) bool`: Checks if a string `s` starts with a specific prefix `prefix` and returns a boolean value indicating the result.

3. `strings.HasSuffix(s, suffix string) bool`: Checks if a string `s` ends with a specific suffix `suffix` and returns a boolean value indicating the result.

4. `strings.Index(s, substr string) int`: Returns the index of the first occurrence of a substring `substr` in a string `s`, or -1 if the substring is not found.

5. `strings.LastIndex(s, substr string) int`: Returns the index of the last occurrence of a substring `substr` in a string `s`, or -1 if the substring is not found.

6. `strings.Split(s, sep string) []string`: Splits a string `s` into a slice of strings at each occurrence of a separator `sep`.

7. `strings.Join(elems []string, sep string) string`: Concatenates a slice of strings `elems` into a single string, with each element separated by a specified separator `sep`.

8. `strings.ToLower(s string) string`: Returns a new string with all characters converted to lowercase.

9. `strings.ToUpper(s string) string`: Returns a new string with all characters converted to uppercase.

These are just a few examples of the functions provided by the `strings` package in Go. You can find more functions and details in the official Go documentation: [strings - Go Package Documentation](https://golang.org/pkg/strings/)