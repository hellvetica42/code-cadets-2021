// Defining the name of this package. This is the name that will be used when this package is
// imported by other packages. Most often the same name as the folder it is contained in.
// `main` is a special package with the `main` function which is the entry point of the program.
package main

// Import statements in order to import functionality from other packages.
import (
	// Importing of Go standard library packages.
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	// Importing of 3rd party packages. Note that the URL of the package is actually the import.
	// This is because Go uses vendoring for its package management - once a 3rd party package
	// is downloaded it's added to your GOPATH with other packages in the `src` folder.
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

// Definition of a struct/class. Names (methods, functions, attributes, variables...)
// which are capitalised are exported which means they can be used by other packages. There
// are no public/private keywords in Go.
type yesNoResponse struct {
	Answer string
	Forced bool
	Image  string
}

// Constant definition. Notice that the `string` type is inferred.
const yesNoURL = "http://yesno.wtf/api"

// Function definition. Notice that the type of the argument comes after the name and that
// the type of the returned value comes at the end of function prototype.
func linearBackoff(retry int) time.Duration {
	// `time.Duration` is a type so `time.Duration(retry)` is casting and `retry` to that type.
	// `time.Second` is a constant.
	return time.Duration(retry) * time.Second
}

// `main` function, entry-point of a Go program.
func main() {
	// Calling `New()` from `pester` package to create an http client. Notice that by using the
	// `:=` notation the type of the variable will be inferred.
	httpClient := pester.New()

	// Setting the struct's attribute.
	httpClient.Backoff = linearBackoff

	// Functions in Go can return multiple values. The last returned value is usually an error
	// value. Go has an in-built `error` type for errors. Go handles errors in code via values
	// which can be easily passed around and modified while they can't be accidentally ignored.
	// There is a `panic` in-built function which has a behaviour similar to exceptions in other
	// languages but should be rarely used.
	httpResponse, err := httpClient.Get(yesNoURL)
	// `if err != nil` is the usual Go idiom to check whether the previous
	// function call was successful.
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards yesno API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of yesno API response"),
		)
	}

	// If not explicitly initialised values in Go are zero-initialised by default.
	// e.g. ints will be 0, bools will be false, strings will be empty...
	var decodedContent yesNoResponse

	// Go has pointers in order to give explicit control to the programmer regarding
	// pass-by-value and pass-by-reference. There is no pointer arithmetic in Go.
	// Well, there is pointer arithmetic via the `unsafe` package but you probably never need it.
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	// Go support some interesting verbs in `printf` (and similar functions) like `%v` here.
	// `%v` prints the whole contents of the struct.
	log.Printf("Response from yesno: %v", decodedContent)
}
