package adapters

import (
	"encoding/json"
	"errors"
	"os"
	"solid-principles/domain"
)

type ProductosEnArchivoRepository struct {
	filePath string
}

func NewProductosEnArchivoRepository(filePath string) *ProductosEnArchivoRepository {
	//Crear el archivo si no existe
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	file.Close()
	return &ProductosEnArchivoRepository{
		filePath: filePath,
	}
}

func (r *ProductosEnArchivoRepository) readAll() ([]domain.Producto, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}
	var productos []domain.Producto
	if len(data) == 0 {
		return productos, nil
	}
	err = json.Unmarshal(data, &productos)
	if err != nil {
		return nil, err
	}
	return productos, nil
}

// Guardar todos los productos en el archivo
func (r *ProductosEnArchivoRepository) writeAll(productos []domain.Producto) error {
	data, err := json.MarshalIndent(productos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

func (r *ProductosEnArchivoRepository) Guardar(p *domain.Producto) error {
	if p.ID == 0 {
		return errors.New("id no puede ser 0")
	}
	productos, err := r.readAll()
	if err != nil {
		return err
	}
	for i := range productos {
		if productos[i].ID == p.ID {
			return errors.New("El producto o su Id ya existen")
		}
	}
	productos = append(productos, *p)
	return r.writeAll(productos)
}

func (r *ProductosEnArchivoRepository) Listar() ([]domain.Producto, error) {
	return r.readAll()
}

func (r *ProductosEnArchivoRepository) Buscar(id int) (*domain.Producto, error) {
	productos, err := r.readAll()
	if err != nil {
		return nil, err
	}
	for i := range productos {
		if productos[i].ID == id {
			return &productos[i], nil
		}
	}
	var producto domain.Producto
	return &producto, errors.New("producto no encontrado")
}

func (r *ProductosEnArchivoRepository) Eliminar(id int) error {
	productos, err := r.readAll()
	if err != nil {
		return err
	}

	nuevaLista := []domain.Producto{}
	found := false

	for i := range productos {
		if productos[i].ID == id {
			found = true
			continue // saltar este producto
		}
		nuevaLista = append(nuevaLista, productos[i])
	}

	if !found {
		return errors.New("producto no encontrado")
	}

	return r.writeAll(nuevaLista)
}
