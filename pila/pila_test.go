package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVaciaDeEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pila.EstaVacia(), "La pila debe estar vacía al inicio.")

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "VerTope() debe lanzar un panic cuando la pila está vacía")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Desapilar() debe lanzar un panic cuando la pila está vacía")
}

func TestPilaUnEntero(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(2)
	require.False(t, pila.EstaVacia(), "La pila no debe estar vacía tras apilar")
	require.Equal(t, 2, pila.VerTope(), "El tope debe ser 2")

	require.Equal(t, 2, pila.Desapilar(), "Desapilar() debe devolver 2")

	require.True(t, pila.EstaVacia(), "La pila debe estar vacía luego de desapilar")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "VerTope() debe lanzar un panic tras vaciar la pila")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Desapilar() debe lanzar un panic tras vaciar la pila")
}

func TestPilaApilarYDesapilarVariosEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(3)
	pila.Apilar(2)
	pila.Apilar(1)

	require.False(t, pila.EstaVacia(), "La pila no debe estar vacía tras apilar")

	require.Equal(t, 1, pila.VerTope(), "El tope debe ser 1")
	require.Equal(t, 1, pila.Desapilar(), "Desapilar() debe devolver 1")

	pila.Apilar(4)

	require.Equal(t, 4, pila.VerTope(), "El tope debe ser 4")
	require.Equal(t, 4, pila.Desapilar(), "Desapilar() debe devolver 4")

	require.Equal(t, 2, pila.VerTope(), "El tope debe ser 2")
	require.Equal(t, 2, pila.Desapilar(), "Desapilar() debe devolver 2")

	require.Equal(t, 3, pila.VerTope(), "El tope debe ser 3")
	require.Equal(t, 3, pila.Desapilar(), "Desapilar() debe devolver 3")

	require.True(t, pila.EstaVacia(), "La pila debe estar vacía tras desapilar")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "VerTope() debe lanzar un panic tras vaciar la pila")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Desapilar() debe lanzar un panic tras vaciar la pila")
}

func TestApilarYDesapilarMilesDeEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	cantidad := 10000

	require.True(t, pila.EstaVacia(), "La pila debe estar vacía al inicio.")

	for i := 0; i < cantidad; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
		require.False(t, pila.EstaVacia(), "La pila no debe estar vacía tras apilar")
	}
	require.Equal(t, cantidad-1, pila.VerTope())

	for i := cantidad - 1; i >= 0; i-- {
		require.False(t, pila.EstaVacia(), "La pila no está vacía antes de desapilar")
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
	}

	require.True(t, pila.EstaVacia(), "La pila debe estar vacía tras desapilar todos los elementos")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "VerTope() debe lanzar un panic tras vaciar la pila")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Desapilar() debe lanzar un panic tras vaciar la pila")
}

func TestPilaVaciaDeFloats(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaUnFloat(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	pila.Apilar(2.5)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 2.5, pila.VerTope())
	require.Equal(t, 2.5, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaVaciaDeStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaUnString(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("hola")
	require.Equal(t, "hola", pila.VerTope())
	require.Equal(t, "hola", pila.Desapilar())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaApilarYDesapilarVariosStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("mundo")
	pila.Apilar("")
	pila.Apilar("hola")
	require.Equal(t, "hola", pila.Desapilar())
	require.Equal(t, "", pila.Desapilar())
	require.Equal(t, "mundo", pila.Desapilar())
}

func TestPilaVaciaDeBooleans(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestPilaUnBoolean(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()
	pila.Apilar(true)
	require.Equal(t, true, pila.VerTope())
	pila.Apilar(false)
	require.Equal(t, false, pila.Desapilar())
}

func TestPilaApilarYDesapilarVariosBooleans(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()
	pila.Apilar(true)
	pila.Apilar(false)
	pila.Apilar(true)
	require.Equal(t, true, pila.Desapilar())
	require.Equal(t, false, pila.Desapilar())
	require.Equal(t, true, pila.Desapilar())
}
