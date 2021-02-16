# Form3 Client API

A client library codebase submission from Dan Abel.

## A warning
This client has been written as a demonstration by an engineer new to Go Lang
It would need work to be done to be a valuable and useful client library.

## Usage guide


## Technical decision log
- used Go Modules - relatively new model, but suitable for this library
- Aiming to use basic Go language constructs and stay away from complexity    
- Early decision to build a version that does operations with minimal 
  set of parameters and config once operating, expose some key params to
  user
- not using test doubles as the docker image provided acts as one (though that 
  might make testing network error handling hard) 
  

### not yet implemented
* `delete` needs to accept a version string, rather than being hardcoded
* `create` needs to return account details
* test error cases that exist
* tet and handle error cases that are not yet managed
* Allow input of params to users to operate the api (remove hardcoding)
* Allow location of service to be set (param, init or env var?)
* error handling is inconsistent
* logging needs removing or improving.
* Add Documentation and example tests
* Handle (and test?) network errors 
* Handle (and test?) input errors
* Support defaults

### Issues
* tests using countAccountsBelow100 only work if there are less than 100 accouns

### potential refactors and code improvement 
 * Be consistent in how we make http requests
 * encapsulate http requests for error handling
 * 

### Production hardening support
* much better error management
* Consider trust model - what data can be trusted from whom?
* Strong TSL support, potentially with 2 way cert checking

### Go lang points of confusion
* Capitalisation of struts and elements
* balance of small methods vs cost of repeated error handling 