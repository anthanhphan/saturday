package routine

import (
	"fmt"
	"reflect"

	"go.uber.org/zap"
)

// Run starts a new goroutine and invokes the provided function with the given arguments.
// Any panic that occurs within the goroutine is recovered and logged.
//
// Parameters:
// - fn: The function to be invoked in the new goroutine. It must be of type `any` to allow flexibility.
// - args: A variadic list of arguments to pass to the invoked function.
//
// Notes:
// - The function is called asynchronously, and any panics inside the goroutine will not crash the program.
// - Panics are handled by the recoverPanic function, which logs the error.
//
// Usage example:
//
//	Run(func(msg string) {
//	    log.Infof("Message: %s", msg)
//	}, "Hello, world!")
func Run(fn any, args ...any) {
	go func() {
		defer recoverPanic() // Recover from any panics within the goroutine and log the error.
		invoke(fn, args)     // Dynamically call the provided function with arguments.
	}()
}

// invoke calls the provided function with the specified arguments.
//
// Parameters:
// - fn: The function to be invoked. It must be a valid function, or an error will be logged.
// - args: A list of arguments to pass to the function.
//
// Notes:
// - The function is dynamically invoked using reflection.
// - If fn is not of type `func`, an error is logged, and the function is not called.
func invoke(fn any, args []any) {
	log := zap.L().With(zap.String("prefix", "invoke")).Sugar()
	funcValue := reflect.ValueOf(fn)

	if funcValue.Kind() != reflect.Func {
		log.Error("invoke: provided value is not a function") // Log an error if the provided value is not a function.
		return
	}

	// Convert the arguments to reflect.Value to match the signature of reflect.Call.
	funcArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		funcArgs[i] = reflect.ValueOf(arg)
	}

	funcValue.Call(funcArgs) // Call the function with the converted arguments.
}

// recoverPanic handles any panic that occurs within a goroutine.
// The panic is caught and logged as an error, preventing the goroutine from crashing.
//
// Notes:
// - If a panic is recovered, it checks whether the recovered value is an error.
// - If it's not an error, it converts it to a string and logs it.
func recoverPanic() {
	log := zap.L().With(zap.String("prefix", "recoverPanic")).Sugar()
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r) // Convert non-error panic values into an error type.
		}
		log.Errorf("recoverPanic: %v", err.Error()) // Log the recovered panic as an error.
	}
}
