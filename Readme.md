# 🛠️ Principios SOLID con Go -- Ejemplo Básico

[![Go](https://img.shields.io/badge/Go-1.22-blue?logo=go)](https://golang.org/)
[![Arquitectura](https://img.shields.io/badge/Arquitectura-Hexagonal-orange)]()
[![SOLID](https://img.shields.io/badge/Principios-SOLID-green)]()

Este proyecto forma parte de mi **Ruta de Crecimiento Profesional como
Ingeniero de Software**.\
Aquí practico los principios **SOLID** y una aproximación a la
**Arquitectura Hexagonal** en Go, con un caso sencillo de gestión de
productos.

------------------------------------------------------------------------

## 🚀 Objetivos

-   Aplicar los **principios SOLID** en un proyecto pequeño.
-   Separar **dominio (lógica de negocio)** de la **infraestructura
    (implementaciones)**.
-   Preparar el código para ser fácilmente extensible (ejemplo: cambiar
    de almacenamiento en memoria a archivo o base de datos sin modificar
    el dominio).

------------------------------------------------------------------------

## 📂 Estructura del Proyecto

    proyectos-david/Practicas/Ruta de crecimiento/PrincipiosSOLID/basic/
    │── domain/              # Núcleo del negocio (modelo + interfaces)
    │   └── producto.go
    │── infra/               # Adaptadores concretos (ej. almacenamiento en memoria)
    │   └── repo_memoria.go
    │── main.go              # Punto de entrada
    │── go.mod               # Definición del módulo Go
    │── repo_memoria_test.go # Tests unitarios

------------------------------------------------------------------------

## ⚙️ Ejecución

### 1. Clonar el repositorio

``` bash
git clone https://github.com/IkingariSolorzano/Ruta-de-crecimiento.git
cd "Ruta-de-crecimiento/PrincipiosSOLID/basic"
```

### 2. Inicializar el módulo (si no existe `go.mod`)

``` bash
go mod init basic
```

### 3. Ejecutar el programa

``` bash
go run .
```

### 📌 Ejemplo de salida

    Productos: [{1 Laptop 1200.5} {2 Mouse 25.99}]
    Encontrado: &{1 Laptop 1200.5}
    Después de eliminar: [{1 Laptop 1200.5}]

------------------------------------------------------------------------

## 🧪 Test Unitario (ejemplo)

Archivo: `infra/repo_memoria_test.go`

``` go
package infra

import (
    "basic/domain"
    "testing"
)

func TestGuardarYBuscarProducto(t *testing.T) {
    repo := NewInMemoryProductoRepository()
    p := &domain.Producto{ID: 1, Nombre: "Teclado", Precio: 45.90}

    err := repo.Guardar(p)
    if err != nil {
        t.Errorf("Error al guardar producto: %v", err)
    }

    encontrado, err := repo.Buscar(1)
    if err != nil {
        t.Errorf("Error al buscar producto: %v", err)
    }

    if encontrado.Nombre != "Teclado" {
        t.Errorf("Se esperaba 'Teclado', se obtuvo '%s'", encontrado.Nombre)
    }
}
```

Ejecutar los tests:

``` bash
go test ./...
```

------------------------------------------------------------------------

## 📖 Principios SOLID aplicados

-   **S (Single Responsibility)** → El `domain` define las reglas de
    negocio, `infra` solo almacena datos.\
-   **O (Open/Closed)** → Puedo añadir nuevos repositorios (Postgres,
    archivo, memoria) sin modificar el dominio.\
-   **L (Liskov Substitution)** → Cualquier repositorio que implemente
    `ProductoRepository` puede usarse de forma intercambiable.\
-   **I (Interface Segregation)** → El dominio define una interfaz
    clara, sin métodos innecesarios.\
-   **D (Dependency Inversion)** → El dominio depende de la
    **abstracción** `ProductoRepository`, no de una implementación
    concreta.

------------------------------------------------------------------------

## 🏗️ Próximos pasos

-   [ ] Crear un **segundo adaptador** que guarde productos en archivo
    JSON.\
-   [ ] Añadir más **tests unitarios** para validar las operaciones
    `Eliminar` y `Listar`.\
-   [ ] Desplegar la aplicación con Docker.

------------------------------------------------------------------------

✍️ **Autor:** [Ikíngari
Solórzano](https://github.com/IkingariSolorzano)\
📅 Proyecto en curso como parte de mi **camino de crecimiento
profesional**.
