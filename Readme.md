# Form3 Client API

A client library codebase submission from Dan Abel.

## A few warnings
This client has been written as a demonstration by an engineer new to Go Lang
It would need work to be done to be a valuable and useful client library.

For example only a few required parameters of create are supported by the API (Org id and country)

## Usage guide


## Technical decision log
- used Go Modules - relatively new model, but suitable for this library
- Aiming to use basic Go language constructs and stay away from complexity    
- Early decision to build a version that does operations with minimal 
  set of parameters and config once operating, expose some key params to
  user
- not using test doubles as the docker image provided acts as one (though that 
  might make testing network error handling hard) 
- For demonstration purposes only allow setting of basic parameters   
- Delete account operation differed documentation around error codes - implemented to work with how API work in the mock presented 

### not yet implemented
* test and handle error cases that are not yet managed - what errors can API throw
* Allow location of service to be set (param, init or env var?)
* Be great to be able to set account name on create (and other 'required fields')  
* Add Documentation and example tests
* Handle (and test?) network errors 
* Handle (and test?) input errors??
* Support defaults
* status code and api error handling should be consistent and useful to caller

### potential refactors and code improvement 
 * Be consistent in how we make http requests
 * encapsulate http requests for error handling

### Production hardening support
* use built in [wrapped errors](https://blog.golang.org/go1.13-errors) to Ensure error management tell users' what the issues are - so they can be fixed
  - both messages from the API service
  - and faults in the client
* Consider trust model - what data can be trusted from whom?
* Strong TSL support, potentially with 2 way cert checking
* retries, throttling and backoffs
* seek good practise guide to get consistancy and usefulness in Error returns 

### Go lang points of confusion
* Capitalisation of struts and elements
* balance of small methods vs cost of repeated error handling 