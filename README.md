Gotry [![Build Status](https://travis-ci.org/selimekizoglu/gotry.svg?branch=master)](https://travis-ci.org/selimekizoglu/gotry)
===============

A tiny library for retrying failing operations


Installation
------------
```shell
$ go get github.com/selimekizoglu/gotry
```

### Example

```go
operation := func() error {
    return nil // or return errors.New("some error")
}

// retry an operation up to 5 times
retry := &Retry{Max: 5, Timeout: 2 * time.Second}
err := Try(operation, retry)
if err != nil {
    // operation failed after 5 retries
    return err
}

// Success
return nil
```
