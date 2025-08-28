package adapters

import (
	"errors"
	"solid-principles/domain"
)

type ProductosEnMemoriaRepository struct {
	productos map[int]*domain.Producto
}

func NewProductosEnMemoriaRepository() *ProductosEnMemoriaRepository {
	return &ProductosEnMemoriaRepository{
		productos: make(map[int]*domain.Producto),
	}
}

func (r *ProductosEnMemoriaRepository) Guardar(p *domain.Producto) error {
	r.productos[p.ID] = p
	return nil
}

func (r *ProductosEnMemoriaRepository) Buscar(id int) (*domain.Producto, error) {
	producto, ok := r.productos[id]
	if !ok {
		return nil, errors.New("producto no encontrado")
	}
	return producto, nil
}

func (r *ProductosEnMemoriaRepository) Eliminar(id int) error {
	if _, ok := r.productos[id]; ok {
		delete(r.productos, id)
		return nil
	}
	return errors.New("producto no encontrado")
}

func (r *ProductosEnMemoriaRepository) Listar() ([]domain.Producto, error) {
	var productos []domain.Producto
	for _, p := range r.productos {
		productos = append(productos, *p)
	}
	return productos, nil
}
