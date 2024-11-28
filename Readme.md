# LeetCode Solutions in Go

This repository contains "boilerplate" code for solving LeetCode problems using Go (Golang). Each problem is organized into its own package, containing a function template and corresponding test cases.

## Directory Structure

Each problem is structured as follows:
```
/id-problem-name
├── problem-name.go        # Function template (boilerplate code)
└── problem-name_test.go   # Test cases for the function
```

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/maxchagin/leetcode.git
    cd leetcode
    ```

2. **Navigate to the desired problem directory:**

    Find the problem that you want to solve under the root directory. Each problem has its own folder, named according to the problem or a descriptive title.

3. **Implement the solution:**

    Edit the `problem-name.go` file and implement your solution in the provided function template.

4. **Test your solution:**

    Use Go's built-in testing tools to run the tests and verify your solution:

    ```bash
    go test ./id-problem-name
    ```

    Replace `id-problem-name` with the directory name of the problem you are working on.

## Example

Here's how a typical problem might be structured:

```
/409-longest-palindrome
├── longest-palindrome.go        # Contains the function template for the problem
└── longest-palindrome_test.go   # Contains test cases for the problem
```

To work on the `longest-palindrome` problem:

- Implement your solution inside `longest-palindrome.go`.
- Run the tests using `go test ./409-longest-palindrome` to check your implementation.


Happy Coding!