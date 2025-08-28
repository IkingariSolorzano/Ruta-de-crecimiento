package adapters

import (
	"os"
	"solid-principles/domain"
	"testing"
)

func TestRepositorioEnArchivo(t *testing.T) {
	// Archivo temporal para pruebas
	testFile := "test_productos.json"
	defer func() {
		err := os.Remove(testFile)
		if err != nil {
			t.Errorf("Error al eliminar archivo: %v", err)
		}
	}()

	// 1️⃣ Crear el repositorio de archivo
	repo := NewProductosEnArchivoRepository(testFile)

	// 2️⃣ Crear productos de prueba
	p1 := &domain.Producto{ID: 1, Nombre: "Laptop", Precio: 1500.50}
	p2 := &domain.Producto{ID: 2, Nombre: "Mouse", Precio: 25.99}

	// 3️⃣ Guardar productos
	if err := repo.Guardar(p1); err != nil {
		t.Errorf("Error al guardar p1: %v", err)
	}
	if err := repo.Guardar(p2); err != nil {
		t.Errorf("Error al guardar p2: %v", err)
	}

	// 4️⃣ Buscar un producto
	encontrado, err := repo.Buscar(1)
	if err != nil {
		t.Errorf("Error al buscar p1: %v", err)
	}
	if encontrado.Nombre != "Laptop" {
		t.Errorf("Se esperaba 'Laptop', se obtuvo '%s'", encontrado.Nombre)
	}

	// 5️⃣ Listar productos
	lista, err := repo.Listar()
	if err != nil {
		t.Errorf("Error al listar productos: %v", err)
	}
	if len(lista) != 2 {
		t.Errorf("Se esperaban 2 productos, se obtuvieron %d", len(lista))
	}

	// 6️⃣ Eliminar un producto
	if err := repo.Eliminar(1); err != nil {
		t.Errorf("Error al eliminar p1: %v", err)
	}

	// 7️⃣ Verificar que fue eliminado
	_, err = repo.Buscar(1)
	if err == nil {
		t.Errorf("Se esperaba error al buscar p1 eliminado")
	}

	// 8️⃣ Listar después de eliminar
	lista, err = repo.Listar()
	if err != nil {
		t.Errorf("Error al listar productos: %v", err)
	}
	if len(lista) != 1 {
		t.Errorf("Se esperaba 1 producto, se obtuvieron %d", len(lista))
	}
}
