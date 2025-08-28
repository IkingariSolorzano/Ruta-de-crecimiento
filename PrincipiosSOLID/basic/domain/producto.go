package domain

type Producto struct {
	ID     int
	Nombre string
	Precio float64
}

type ProductoRepository interface {
	Guardar(p *Producto) error
	Buscar(id int) (*Producto, error)
	Eliminar(id int) error
	Listar() ([]Producto, error)
}
