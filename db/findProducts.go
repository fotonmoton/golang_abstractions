package db

type Connection struct {
	Products []string
}

func (c *Connection) Find(searchTerm string) []string {
	return c.Products
}
