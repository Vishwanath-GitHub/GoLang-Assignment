package main
import (
  "fmt"
  "errors"
  "io/ioutil"
)
//The errors.Wrap function returns a new error that adds context to the original error.
func main() {
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.Wrap(err, "read failed")
  }
switch err := errors.Cause(err).(type) {
case *MyError:
        // handle specifically
default:
        // unknown error
}
  cause := errors.New("whoops")
	err := errors.Unwrap(cause) //Unwrap returns the next error in err's chain. If there is no next error, Unwrap returns nil.
	fmt.Println(err)
}
//errors.Cause will recursively retrieve the topmost error which does not implement causer, which is assumed to be the original cause.
type causer interface {
        Cause() error
}
