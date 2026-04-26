package di

import (
	"fmt"
	"reflect"
)

type registry struct {
	providers map[reflect.Type]interface{}
}

func newRegistry() *registry {
	return &registry{
		providers: make(map[reflect.Type]interface{}),
	}
}

func (r *registry) provide(outType reflect.Type, constructor interface{}) {
	r.providers[outType] = constructor
}

func (r *registry) validateConstructor(constructor interface{}) reflect.Type {
	ctorType := reflect.TypeOf(constructor)

	fmt.Printf("ctorType <%v>\n", ctorType)

	if ctorType.Kind() != reflect.Func {
		panic("constructor must be a function type")
	}

	errType := reflect.TypeOf((*error)(nil)).Elem()

	var targetOut reflect.Type

	for fnOut := range ctorType.Outs() {
		if fnOut == errType {
			continue
		}

		if targetOut != nil {
			panic("constructor can return only (T) or (T, error)")
		}

		targetOut = fnOut
	}

	if targetOut == nil {
		panic("constructor must return at least one value")
	}

	return targetOut
}
