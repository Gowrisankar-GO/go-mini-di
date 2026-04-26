package di

import (
	"fmt"
	"reflect"
)

func (r *registry) resolve(t reflect.Type) []reflect.Value {
	ctor, ok := r.providers[t]
	if !ok {
		panic(fmt.Sprintf("unknown provider: cannot able to find the constructor %v dependency", t.String()))
	}

	ctorType := reflect.TypeOf(ctor)

	numIn := ctorType.NumIn()

	var args []reflect.Value

	for i := 0; i < numIn; i++ {

		ctorIn := ctorType.In(i)

		v := r.resolve(ctorIn)

		args = append(args, v...)
	}

	return reflect.ValueOf(ctor).Call(args)
}
