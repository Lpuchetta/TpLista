package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVaciaDeEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia(), "La cola debe estar vacía al inicio.")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "VerPrimero() debe lanzar un panic cuando la cola está vacía")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Desencolar() debe lanzar un panic cuando la cola está vacía")
}

func TestEncolarYDesencolarUnEntero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())
	require.False(t, cola.EstaVacia(), "EstaVacia() debe ser false luego de apilar un número entero")

	require.Equal(t, 1, cola.Desencolar())

	require.True(t, cola.EstaVacia(), "La cola debe estar vacía luego de desencolar el único elemento encolado")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "VerPrimero() debe lanzar un panic cuando la cola está vacía")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Desencolar() debe lanzar un panic cuando la cola está vacía")

}

func TestEncolarYDesencolarVariosEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola debe estar vacía al inicio")

	cola.Encolar(1)

	require.False(t, cola.EstaVacia(), "La cola no debe estar vacía luego de encolar")
	require.Equal(t, 1, cola.VerPrimero(), "El primer elemento debe ser el 1")

	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero(), "El primer elemento debe ser el 1")

	require.Equal(t, 1, cola.Desencolar())
	require.False(t, cola.EstaVacia(), "La cola tiene aún elementos")
	require.Equal(t, 2, cola.VerPrimero(), "El primero debe ser el 2")

	cola.Encolar(3)
	require.Equal(t, 2, cola.VerPrimero(), "El primer elemento debe ser el 2")

	require.Equal(t, 2, cola.Desencolar())
	require.False(t, cola.EstaVacia(), "La cola aún tiene elementos")

	require.Equal(t, 3, cola.VerPrimero())
	require.Equal(t, 3, cola.Desencolar())

	require.True(t, cola.EstaVacia(), "La cola debe estar vacía luego de desencolar todos los elementos encolados")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "VerPrimero() debe lanzar un panic cuando la cola está vacía")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Desencolar() debe lanzar un panic cuando la cola está vacía")
}

func TestEncolarDespuesDeDesencolarEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	cola.Encolar(2)
	require.Equal(t, 1, cola.Desencolar(), "Desencolar() debe devolver el primer elemento encolado (1)")

	cola.Encolar(3)
	require.Equal(t, 2, cola.VerPrimero(), "VerPrimero() debe mostrar el nuevo primero (2) después de desencolar")

	require.Equal(t, 2, cola.Desencolar(), "Desencolar debe devolver 2")
	require.Equal(t, 3, cola.Desencolar(), "Desencolar debe devolver 3")

	require.True(t, cola.EstaVacia(), "La cola debe estar vacía después de desencolar todos los elementos")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Desencolar() debe lanzar un panic cuando la cola está vacía")
}

func TestEncolarYDesencolarMilesDeEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	inicio := 0
	cantidad := 100000

	for i := inicio; i < cantidad; i++ {
		cola.Encolar(i)
		require.False(t, cola.EstaVacia(), "La cola no debe estar vacía luego de encolar un elemento")
		require.Equal(t, inicio, cola.VerPrimero())
	}

	require.Equal(t, inicio, cola.VerPrimero())
	require.False(t, cola.EstaVacia())

	for i := inicio; i < cantidad; i++ {
		require.False(t, cola.EstaVacia(), "La cola no debe estar vacía antes de desencolar")
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia(), "La cola debe estar vacía luego de desencolar todos los números enteros")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "VerPrimero() debe lanzar un panic cuando la cola está vacía")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Desencolar() debe lanzar un panic cuando la cola está vacía")
}

func TestEncolarYDesencolarUnFloat(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	cola.Encolar(3.14)
	require.Equal(t, 3.14, cola.VerPrimero(), "VerPrimero() debe devolver 3.14")
	require.Equal(t, 3.14, cola.Desencolar(), "Desencolar() debe devolver 3.14")

	require.True(t, cola.EstaVacia(), "La cola debe quedar vacía después de desencolar el único float")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Desencolar() debe lanzar un panic cuando la cola está vacía")
}

func TestEncolarYDesencolarStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	require.True(t, cola.EstaVacia(), "La cola debe estar vacía al inicio")

	cola.Encolar("a")
	require.False(t, cola.EstaVacia())
	require.Equal(t, "a", cola.VerPrimero())
	require.Equal(t, "a", cola.Desencolar())
	require.True(t, cola.EstaVacia(), "La cola debe estar vacía luego de desencolar 'a'")

	cola.Encolar("b")
	require.False(t, cola.EstaVacia())
	require.Equal(t, "b", cola.VerPrimero())
	require.Equal(t, "b", cola.Desencolar())
	require.True(t, cola.EstaVacia(), "La cola debe estar vacía luego de desencolar 'b'")

	cola.Encolar("c")
	cola.Encolar("d")
	require.Equal(t, "c", cola.VerPrimero())
	require.Equal(t, "c", cola.Desencolar())
	require.Equal(t, "d", cola.VerPrimero())
	require.Equal(t, "d", cola.Desencolar())
	require.True(t, cola.EstaVacia(), "La cola debe estar vacía luego de desencolar todos los elementos")
}
