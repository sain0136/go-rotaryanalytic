### Notes

The `http.HandlerFunc` type is a built-in type in Go's `net/http` package. It's an adapter that allows the use of ordinary functions as HTTP handlers.
If `f` is a function with the appropriate signature (`func(w http.ResponseWriter, r *http.Request)`), `http.HandlerFunc(f)` is a `Handler` that calls `f`.

In your `AuthorizeHandler` function, you're returning an anonymous function with the signature `func(w http.ResponseWriter, r *http.Request)`. When you call `AuthorizeHandler()`, it returns this anonymous function as an `http.HandlerFunc`.

Here's the key part: `http.HandlerFunc` itself satisfies the `http.Handler` interface because it has a `ServeHTTP` method. This method just calls the function `f`.
So when you return an `http.HandlerFunc` from your `AuthorizeHandler` function, you're actually returning something that has a `ServeHTTP` method, even though you
don't see it in your code.

This is why you can use `AuthorizeHandler()` as an argument to `http.Handle`. Even though `AuthorizeHandler()` doesn't explicitly return something with a `ServeHTTP` method, it returns an `http.HandlerFunc`, which does have a `ServeHTTP` method.

This is a common pattern in Go. It allows you to create flexible APIs and use functions as interfaces. It might seem a bit confusing at first, but it's a powerful feature once you get the hang of it.

## Constants

In Go, constants are not considered variables, and they are treated differently by the debugger. Constants are inlined by the compiler, meaning their values are directly substituted into the code where they are used. As a result, they do not appear in the list of local variables during debugging.
Here's a brief explanation:

Why Constants Don't Appear in Debugger
Inlining by Compiler: The Go compiler inlines constants, replacing their usage with their values directly in the code. This optimization means that the constant itself doesn't exist as a separate entity at runtime.
Not Variables: Constants are not variables; they are fixed values that do not change. The debugger typically shows variables that can change state during execution

How to Work Around This
If you need to see the value of a constant during debugging, you can assign it to a variable within your function. For example:

Now, logsPerPageVar will appear in the list of local variables during debugging, and you can inspect its value.

const logsPerPage int = 10

```go
func main() {
    page := 1
    logsPerPageVar := logsPerPage // Assign constant to a variable
    totalLogs := logsPerPageVar * page
    fmt.Println(totalLogs)
}
```

Conclusion
Constants are optimized away by the compiler and do not appear in the debugger's list of local variables. If you need to inspect their values, you can assign them to variables within your functions. This approach allows you to see their values during debugging sessions.
