---
id: Error handling
aliases:
  - Error handling in Go
tags: []
---

# Error handling in Go

In go there's no try catch blocks. You get the error returned along side the "correct" result. So you can handle the error as if it was any other value. Here's an example:

```go
func main () {
  user, err := getUser()
  if err != nil {
    fmt.Println(err)
    return
  }
  // no errors code here
}

func getUser() (user, error) {
  // do something that might return an error
  return user
}
```

Here, `error` is a built-in interface provided by go to express, well, errors, duh!

## The error interface

The error interface is just an interface and errors are just values. The interface has only one method and that's the `Error` method:

```go
type error interface {
  Error() string
}
```

This means that we can create whatever struct we want to represent an error, we only need to implement the `Error` method and that meets the interface criteria! You don't need to extend the error class or anything like that. So to implement an error you can just do something like this:

```go
type userError struct {
  name string
}

func (u userError) Error() {
  return fmt.Sprintf("%v has a problem with his account.", u.name)
}

func sendSms(msg string, userName string) (bool,error) {
  if!canSendSms(userName) {
    return false, userError{name: userName}
  }
  return true, nil
}
```

## The errors package

Even though the error interface is very siple there's quite a lot of boilerplate if you just want an error. Thats where the error package comes in handy. You can use the `errors.New` method to create an error on the fly with just a simple string.

```go
var err error = errors.New("this is an easy error")
```

## Deeper error handling

In Go errors are basically strings. If you want deeper errors that just returning that string, you'll need to do some string parsing. Luckily Go has a great `strings` module on the standard library. Here's an example to ignore a certain error logging on a web scraper:

```go
_, err = db.CreatePost(context.Background(), database.CreatePostParams{
    ID:          uuid.New(),
    CreatedAt:   time.Now().UTC(),
    UpdatedAt:   time.Now().UTC(),
    Title:       item.Title,
    Description: desc,
    Url:         item.Link,
    PublishedAt: pubAt,
    FeedID:      feed.ID,
})
if err != nil {
    // the error we want to ignore is: 'pq: duplicate key value violates unique constraint "posts_url_key"'
    if strings.Contains(err.Error, "duplicate key") {
        continue
    }
    log.Println("unable to save post:", err)
}
```

Since we are inside a loop we can use continue, otherwise we could use return or any other thing to avoid logging that error.

# SQL errors

When using sqlc, the _de facto_ libbrary for sql, we will get different errors related to database operations. To handle them we will want to try and cast the errors returned into SQL errors. To try to make anything into a struct we can just cast it using the following syntax:

```go
<ERROR>.(*<SQL_ERROR_STRUCT>)
// For example:
pqErr, ok := err.(*pq.Error)
```

So if we are trying to process a specific error from the database we might want to do something like:

```go
if pqErr, ok := err.(*pq.Error); ok {
    log.Println(pqErr.Code.Name())
}
```

Here are some examples of codes you might find, they are pretty self explanatory. You can always check the documentation:

- `foreign_key_violation`
- `unique_violation`
