package main

import (
	"fmt"
	"solid-principles/adapters"
	"solid-principles/domain"
)

func main() {

	repo := adapters.NewProductosEnMemoriaRepository()

	repo.Guardar(&domain.Producto{
		ID:     1,
		Nombre: "Producto 1",
		Precio: 100,
	},)
	repo.Guardar(&domain.Producto{
		ID:     2,
		Nombre: "Producto 2",
		Precio: 200,
	},)

	productos, err := repo.Listar()
	if err != nil {
		panic(err)
	}
	fmt.Println(productos)
	
	producto, err := repo.Buscar(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(producto)
	
	//agregar producto
	repo.Guardar(&domain.Producto{
		ID:     3,
		Nombre: "Producto 3",
		Precio: 300,
	},)
	//listar de nuevo
	productos, err = repo.Listar()
	if err != nil {
		panic(err)
	}
	fmt.Println(productos)
	//eliminar producto
	repo.Eliminar(1)
	//listar de nuevo
	productos, err = repo.Listar()
	if err != nil {
		panic(err)
	}
	fmt.Println(productos)

}
