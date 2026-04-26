package di

func (c *container) Provide(constructor interface{}) {
	dep := c.registry.validateConstructor(constructor)
	c.registry.provide(dep, constructor)
}
