# Form3 Client API

A client library codebase submission from Dan Abel.

## A few warnings
This client has been written as a demonstration by an engineer new to Go Lang
It would need work to be done to be a valuable and useful client library - these items are detailed below.

For example only a few required parameters of create are supported by the API (Org id and country)

## Usage guide

use `-domain=<domainname>` to set the domain the API is running on - default is 'localhost'

## Test guide

`docker-compose up` gathers the resources and images and runs the tests.

Alternatively run `docker-compose up accountapi` and then run `go test`

## Technical decision log
- used Go Modules - relatively new model, but suitable for this library
- Aiming to use basic Go language constructs and stay away from complexity    
- Early decision to build a version that does operations with minimal 
  set of parameters and config once operating, expose some key params to
  user
- not using test doubles as the docker image provided acts as one (though that 
  might make testing network error handling hard) 
- For demonstration purposes only allow setting of basic parameters (not yet meeting expectations of API docs)   
- Delete account operation differed documentation around error codes - implemented to work with how API work in the mock presented
- Running from docker-compose runs the tests until they pass due to issues in waiting for dependencies to be available. 
  I'd love to have time to look into [wait-for-it](https://github.com/vishnubob/wait-for-it)

### not yet implemented
* Be great to be able to set account name on create (and other 'required fields')  
* Add Documentation and example tests
* status code and api error handling should be consistent and useful to caller

### potential refactors and code improvement 
 * Be consistent in how we make http requests
 * encapsulate http requests for error handling
 * test unit network error handling
 * Support defaults

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