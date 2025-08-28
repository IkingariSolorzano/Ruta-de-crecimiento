package domain

import "testing"

func TestProducto(t *testing.T) {
	producto := Producto{
		ID:     1,
		Nombre: "Producto 1",
		Precio: 100,
	}
	if producto.ID != 1 {
		t.Errorf("ID = %d, expected 1", producto.ID)
	}
	if producto.Nombre != "Producto 1" {
		t.Errorf("Nombre = %s, expected Producto 1", producto.Nombre)
	}
	if producto.Precio != 100 {
		t.Errorf("Precio = %f, expected 100", producto.Precio)
	}
}
