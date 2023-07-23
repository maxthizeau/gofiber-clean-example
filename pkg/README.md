# ./pkg

Mainly used to use third party libraries.
Functions in this folder should have no information about the incoming/outgoing requests.
These packages are used in the internal but should not use any of the internal package. They should be as much isolated as possible.

Example :

- auth (handle new JWT token, parsing JWT, refreshing token )
- email (handle send mail)
- cache (manages cache)
- hash (handle hash password)
- otp (create a otp)
- payment (handle requests to third party payment processor)
- logger
- ...
