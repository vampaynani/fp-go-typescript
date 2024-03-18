# Test 1 - Backend

## Description

The purpose of this code is to create a Go Wrapper for [JokeAPI](https://sv443.net/jokeapi/v2/).

In order to test it you can clone the repo and run it using:

```bash
go run . '{ "categories": ["programming"], "language": "en" }'
```

Also you can compile the code and use it as follows:

```bash
jsontoapi '{ "categories": ["programming"], "language": "en" }'
```

## Architectural Considerations and Requirements

1.  The input must be in JSON format. On this we should consider structures and types complexity -nested objects, arrays, date or formatted types-.

2.  Implementing the correct validation rules will help us maintain data integrity and ensure the security of our system, which are the required fields? Are there any optional fields? Which are the validation boundaries for each field?.

3.  Logging the input to STDOUT in Go is quite straightforward and well-documented. It would be worth considering whether there is any information that could be deemed sensitive and should not be displayed.

4.  Making an HTTP call to an external API is simple thanks to the libraries included in Go. It is essential to ensure that the data sent has been previously validated. Questions that came to me mind, do we want to do something with the response body? Do we want to track the requests? Is there a a retry policy?

5.  It is important to consider the handling of errors that may occur at points involving interactions that are not under our code control, such as when parsing the JSON, validating it, and making the HTTP request.

## Use cases

### Happy paths

With category and language declared

```json
{ "categories": ["programming"], "language": "en" }
```

With more than one category and language declared

```json
{ "categories": ["programming", "dark"], "language": "en" }
```

#### Flags and amount should be optional

With more than one category, language and flags declared

```json
{
  "categories": ["programming", "dark"],
  "language": "en",
  "flags": ["nsfw", "political"]
}
```

With more than one category, language and amount declared

```json
{ "categories": ["programming", "dark"], "language": "en", "amount": 5 }
```

### Error triggering

#### Parsing error

```json
{"categories":"programming,"language":"es"}
```

#### Required error

- The JSON string should contain categories and language properties.

```json
{ "amount": 0, "flags": ["nsfw", "political"] }
```

#### Categories error

- Categories is not an slice
  ```json
  { "categories": "programming", "language": "es" }
  ```
- Categories is not a list which includes any of `programming`,`misc`,`dark`,`pun`,`spooky` or `christmas` values.

  ```json
  { "categories": ["something", "else"], "language": "es" }
  ```

#### Language error

- Language is not a string with any of the values `cs`, `de`, `en`, `es`, `fr` or `pt` values.
  ```json
  { "categories": ["programming"], "language": "cd" }
  ```

#### Flags error

- Flags is not an slice
  ```json
  {
    "categories": ["programming", "dark"],
    "language": "en",
    "flags": "nsfw"
  }
  ```
- Flags is not a list which includes any of `nsfw`, `religious`, `political`, `racist`, `sexist`or `explicit` values.

  ```json
  {
    "categories": ["programming"],
    "language": "en",
    "flags": ["dark"]
  }
  ```

  #### Amount error

- Amount is not int
  ```json
  {
    "categories": ["programming", "dark"],
    "language": "en",
    "amount": "a"
  }
  ```
- Amount is less than 1

  ```json
  {
    "categories": ["programming"],
    "language": "en",
    "amount": 0
  }
  ```

- Amount is more than 10

  ```json
  {
    "categories": ["programming"],
    "language": "en",
    "amount": 11
  }
  ```
