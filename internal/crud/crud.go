package crud

type crud struct {
	id       uint32
	name     string
	required bool
}

type Crud interface {
	create()
	read()
	update()
	delete()
}

func NewCrud() Crud {
	return &crud{}
}

func (c *crud) create() {
	//TODO
}

func (c *crud) read() {
	//TODO
}

func (c *crud) update() {
	//TODO
}

func (c *crud) delete() {
	//TODO
}
