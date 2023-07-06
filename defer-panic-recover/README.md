#### Summary
- Blog: https://go.dev/blog/defer-panic-and-recover

This blog post discusses three less common control flow mechanisms in Go: defer, panic, and recover. It explains how defer statements are used to simplify clean-up actions, how panic stops the ordinary flow of control and begins panicking, and how recover regains control of a panicking goroutine. Examples and rules for using these mechanisms are provided.

#### Facts
- A defer statement pushes a function call onto a list, and the list of saved calls is executed after the surrounding function returns.
- defer statements are commonly used to ensure proper clean-up actions, such as closing files.
- Deferred function calls are executed in Last In First Out (LIFO) order after the surrounding function returns.
- Deferred functions can read and assign to the named return values of the returning function.
- panic is a built-in function in Go that stops the ordinary flow of control and begins panicking. It can be caused by runtime errors or manually invoked.
- recover is a built-in function that regains control of a panicking goroutine. It can only be used inside deferred functions.
- The example program provided demonstrates the mechanics of panic and defer, showing how panics propagate up the call stack and how recovery can be achieved.