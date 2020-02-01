package errors

// Message represents a textual error.
//
//  const ErrInterrupted errors.Message = "operation interrupted"
//
//  func Action() error {
//  	...
//  }
//
//  switch err := Action(); err {
//  case ErrInterrupted:
//  	http.Error(..., err.Error(), http.StatusRequestTimeout)
//  case ...:
//  	http.Error(..., http.StatusInternalServerError)
//  }
//
type Message string

// Message returns a string representation of the error.
func (err Message) Error() string {
	return string(err)
}
