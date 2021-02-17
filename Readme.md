# Form3 Client API

A client library codebase submission from Dan Abel.

## A few warnings
This client has been written as a code sample by an engineer new to Go Lang.

Current operations are a minimal implementation as a demonstration.

There are four levels of further work that would need to be done to make this a useful API

1. Complete current interface so that they provide a useful interface. 
   Only a few required parameters of create are supported by the API (Org id and country)
2. Add additional methods to provide a useful service
3. Be reviewed and refreshed by engineers experienced in Go Lang
4. Harden for production usage. 

There are further details in the *Usage Guide* and the *Development details* sections below  



## Usage guide

The Api presents three operations and allows the setting of the API domain.

Use `setDomain(apiDomain, apiPort)` to set the location of the API to run against.

The defaults are `localhost:8080`

### Error responses

`operation failed` is a standard error returned when a request 
cannot be sent, or a response cannot be received or parsed. Please
check the logs for further details of the issue.

`API reported fail` is a standard error when the api reports an error - further detail
is supplied in the error. 

### `createAccount`

Creates an account with a generated account number using OrgId and county code provided. 

`createAccount(orgId uuid.UUID, country string)`

Where `country` should be the two-letter country code specified in https://api-docs.form3.tech/api.html

On failure will return an error

 - `validation failure` - API validation failures are reported as an error with further details
 - `API reported fail` - Other API failures
 - `operation failed` - for other failures

On success will return the account details with these key fields 

```
{
Id uuid.UUID  // account id
OrganisationId uuid.UUID // org ID
Attributes.country // Country code 
Version int32 // current version of account
}
```

### `fetchAccount`

Gets the account details for the given account number

`fetchAccount(id uuid.UUID)` 

On failure will return an error beginning
 - `account does not exist` - if account requested does not exist. 
 - `API reported fail ` - on API errors - further detail is contained within the error string
 - `operation failed` for other failures

On success will return the account details with the key fields as `CreateAccount` above.

### `deleteAccount`

Requests the deletion of the given account, using the known current version to avoid removal of 
account that have been changed since being retrieved.

`deleteAccount(id uuid.UUID, version int32)`

On failure will return an error
- `delete failed` - if versions do not match or other API operation error
- `operation failed` for other failures

## Test guide

`docker-compose up` gathers the resources and images, and runs the tests.

Alternatively, run `docker-compose up accountapi` and then run `go test`

## Development details
### Technical decision log
- used Go Modules - relatively new model, but suitable for this library
- Aiming to use basic Go language constructs and stay away from complexity    
- Early decision to build a version that does operations with minimal 
  set of parameters and config once operating, expose some key params to
  user
- Not using test doubles as the docker image provided acts as one (though that 
  might make testing network error handling hard) 
- For demonstration purposes only allow setting of basic parameters (not yet meeting expectations of API docs)   
- Delete account operation differed documentation around error codes - implemented to work with how API work in the mock presented
- Running from docker-compose runs the tests until they pass due to issues in waiting for dependencies to be available.

### What I want to do, but did not have time to complete 
  I'd love to have time to look into [wait-for-it](https://github.com/vishnubob/wait-for-it)
- I'm unhappy with the errors and would like to both support nested errors and also have structure to contain details from the API
- I would like to use HTTP stubbing to allow me to test both network errors and http 500s - I ran out of time to find out how to do this well
- offer a richer API for create that is more like an usable API
  
### potential refactors and code improvement
* Add [testable examples](https://blog.golang.org/examples)  
* Be consistent in how we make http requests
 * encapsulate http requests for error handling in one place
 * test unit network error handling
 * Support defaults in the API
 * Improve error returning in the API. [This guide](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) has good points. 

### Production hardening support
* use built in [wrapped errors](https://blog.golang.org/go1.13-errors) to Ensure error management tell users' what the issues are - so they can be fixed
  - messages from the API service
  - faults in the client
* Consider trust model - what data can be trusted from whom?
* Strong TSL support, potentially with 2 way cert checking
* retries, throttling and backoffs
* seek good practise guide to get consistency and usefulness in Error returns
* ensure logging does not reveal secrets or personal information

### Go lang points of confusion
* Capitalisation of struts and elements
* balance of small methods vs cost of repeated error handling