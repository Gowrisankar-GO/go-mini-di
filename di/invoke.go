package di

import "reflect"

func (c *container) Invoke(fn interface{}) {
	fnType := reflect.TypeOf(fn)

	if fnType.Kind() != reflect.Func {
		panic("Invoke expects a function")
	}

	var args []reflect.Value
	for param := range fnType.Ins() {
		args = append(args, c.registry.resolve(param)...)
	}

	reflect.ValueOf(fn).Call(args)
}
