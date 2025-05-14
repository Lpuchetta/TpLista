package diccionario_test

import (
	"fmt"
	"strings"
	TDADiccionarioOrdenado "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararEnteros(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func compararStrings(a, b string) int {
	return strings.Compare(a, b)
}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	dicInt := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)

	require.EqualValues(t, 0, dicInt.Cantidad())
	require.False(t, dicInt.Pertenece(21))

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Obtener(5)
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Borrar(4)
	})

	dicInt.Guardar(15, 20)
	require.EqualValues(t, 1, dicInt.Cantidad())
	require.EqualValues(t, 20, dicInt.Obtener(15))

	require.True(t, dicInt.Pertenece(15))
	require.False(t, dicInt.Pertenece(20))

	require.EqualValues(t, 20, dicInt.Borrar(15))
	require.EqualValues(t, 0, dicInt.Cantidad())

	require.False(t, dicInt.Pertenece(15))

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Obtener(15)
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Borrar(121)
	})
}

func TestGuardarYBorrar(t *testing.T) {
	dicStr := TDADiccionarioOrdenado.CrearABB[string, string](compararStrings)

	require.EqualValues(t, 0, dicStr.Cantidad())
	require.False(t, dicStr.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicStr.Borrar("")
	})

	dicStr.Guardar("A", "B")
	dicStr.Guardar("L", "w")
	dicStr.Guardar("s", "l")

	require.EqualValues(t, 3, dicStr.Cantidad())
	require.True(t, dicStr.Pertenece("A"))
	require.True(t, dicStr.Pertenece("L"))
	require.True(t, dicStr.Pertenece("s"))

	require.EqualValues(t, "w", dicStr.Obtener("L"))
	require.EqualValues(t, "w", dicStr.Borrar("L"))
	require.False(t, dicStr.Pertenece("L"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicStr.Obtener("L")
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicStr.Borrar("L")
	})
	require.EqualValues(t, 2, dicStr.Cantidad())

	require.EqualValues(t, "B", dicStr.Obtener("A"))
	dicStr.Guardar("A", "H")

	require.True(t, dicStr.Pertenece("A"))
	require.EqualValues(t, "H", dicStr.Obtener("A"))
	require.EqualValues(t, 2, dicStr.Cantidad())
	require.EqualValues(t, "H", dicStr.Borrar("A"))

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicStr.Obtener("A")
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicStr.Borrar("A")
	})
	require.EqualValues(t, 1, dicStr.Cantidad())

	require.EqualValues(t, "l", dicStr.Obtener("s"))
	dicStr.Guardar("s", "m")
	dicStr.Guardar("s", "s")
	dicStr.Guardar("s", "x")
	dicStr.Guardar("s", "o")
	dicStr.Guardar("s", "w")
	dicStr.Guardar("s", "q")

	require.EqualValues(t, 1, dicStr.Cantidad())
	require.EqualValues(t, "q", dicStr.Obtener("s"))
	require.EqualValues(t, "q", dicStr.Borrar("s"))

	require.EqualValues(t, 0, dicStr.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicStr.Obtener("y")
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicStr.Borrar("s")
	})

}

func TestVolumenDiccionarioOrdenado(t *testing.T) {
	const N = 1000
	dicInt := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)

	for i := 0; i < N; i++ {
		dicInt.Guardar(i, i*2)
		require.True(t, dicInt.Pertenece(i))
		require.EqualValues(t, i*2, dicInt.Obtener(i))
		require.EqualValues(t, i+1, dicInt.Cantidad())
	}

	eliminados := N
	for i := 0; i < N; i++ {
		require.EqualValues(t, i*2, dicInt.Borrar(i))
		require.False(t, dicInt.Pertenece(i))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
			dicInt.Obtener(i)
		})
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
			dicInt.Borrar(i)
		})
		eliminados--
		require.EqualValues(t, eliminados, dicInt.Cantidad())
	}

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Obtener(2)
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Borrar(12)
	})
	require.EqualValues(t, 0, dicInt.Cantidad())
	require.False(t, dicInt.Pertenece(2))
}

func TestGuardarMismaClave(t *testing.T) {
	dicInt := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)
	const veces = 100
	clave := 21
	for i := 0; i < veces; i++ {
		dicInt.Guardar(clave, i)
		require.EqualValues(t, 1, dicInt.Cantidad())
		require.EqualValues(t, i, dicInt.Obtener(clave))
		require.True(t, dicInt.Pertenece(clave))
	}
	require.EqualValues(t, veces-1, dicInt.Obtener(clave))
	require.EqualValues(t, 1, dicInt.Cantidad())
	require.EqualValues(t, veces-1, dicInt.Borrar(clave))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Obtener(clave)
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
		dicInt.Borrar(clave)
	})
	require.EqualValues(t, 0, dicInt.Cantidad())
}

func TestBorrarNodoConDosHijos(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)

	dic.Guardar(15, 150)
	dic.Guardar(10, 100)
	dic.Guardar(20, 200)
	dic.Guardar(5, 50)
	dic.Guardar(12, 120)

	valor := dic.Borrar(10)
	require.EqualValues(t, 100, valor)

	require.True(t, dic.Pertenece(12))
	require.False(t, dic.Pertenece(10))
	require.EqualValues(t, 120, dic.Obtener(12))

	require.EqualValues(t, 4, dic.Cantidad())
}

func TestBorrarNodoUnHijoDerecho(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, string](compararEnteros)
	dic.Guardar(10, "a")
	dic.Guardar(20, "b")

	require.EqualValues(t, "a", dic.Borrar(10))
	require.True(t, dic.Pertenece(20))
	require.False(t, dic.Pertenece(10))
}

func TestBorrarNodoUnHijoIzquierdo(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, string](compararEnteros)
	dic.Guardar(20, "a")
	dic.Guardar(10, "b")

	require.EqualValues(t, "a", dic.Borrar(20))
	require.True(t, dic.Pertenece(10))
	require.False(t, dic.Pertenece(20))
}

func TestStressDiccionarioOrdenado(t *testing.T) {
	const N = 10000
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)

	for i := N - 1; i >= 0; i-- {
		dic.Guardar(i, i*10)
	}

	require.Equal(t, N, dic.Cantidad())

	for i := 0; i < N; i++ {
		require.True(t, dic.Pertenece(i))
		require.Equal(t, i*10, dic.Obtener(i))
	}

	for i := 0; i < N; i++ {
		require.Equal(t, i*10, dic.Borrar(i))
		require.False(t, dic.Pertenece(i))
	}

	require.Equal(t, 0, dic.Cantidad(), "El diccionario debe quedar vacío después de borrar todo")
}

func TestClavesCadenasValoresLongitudes(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[string, int](compararStrings)

	claves := []string{"a", "aa", "aaa", "aaaa", "aaaaa"}
	for _, clave := range claves {
		dic.Guardar(clave, len(clave))
	}

	var resultado []int
	dic.Iterar(func(clave string, valor int) bool {
		resultado = append(resultado, valor)
		return true
	})
	require.Equal(t, []int{1, 2, 3, 4, 5}, resultado)
}

func TestIteradorInternoABBClaves(t *testing.T) {
	clave1 := "Manzana"
	clave2 := "Banana"
	clave3 := "Cereza"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionarioOrdenado.CrearABB[string, *int](compararStrings)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[2], nil)

	cs := make([]string, 3)
	cantidad := 0
	dic.Iterar(func(clave string, _ *int) bool {
		cs[cantidad] = clave
		cantidad++
		return true
	})

	require.EqualValues(t, []string{"Banana", "Cereza", "Manzana"}, cs)
}

func TestIteradorInternoABBValores(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)
	dic.Guardar(5, 10)
	dic.Guardar(3, 20)
	dic.Guardar(7, 30)

	suma := 0
	dic.Iterar(func(_ int, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 60, suma, "La suma de los valores no coincide")
}

func TestIteradorInternoABBValoresConBorrados(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[string, int](compararStrings)
	claves := []string{"A", "B", "C", "D", "E"}
	for i, c := range claves {
		dic.Guardar(c, i+1)
	}
	dic.Borrar("B")
	dic.Borrar("D")

	suma := 0
	dic.Iterar(func(_ string, dato int) bool {
		suma += dato
		return true
	})

	require.EqualValues(t, 9, suma)
}
func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[string, int](compararStrings)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoInterno(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, string](compararEnteros)
	dic.Guardar(5, "A")
	dic.Guardar(3, "B")
	dic.Guardar(7, "C")
	dic.Guardar(2, "D")
	dic.Guardar(4, "E")

	desde := 3
	hasta := 5
	resultado := []string{}
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		resultado = append(resultado, valor)
		return true
	})

	require.EqualValues(t, []string{"B", "E", "A"}, resultado)
}

func TestIteradorCorteInterno(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)
	for i := 0; i < 10; i++ {
		dic.Guardar(i, i)
	}

	contador := 0
	dic.Iterar(func(clave int, _ int) bool {
		if clave == 5 {
			return false
		}
		contador++
		return true
	})

	require.EqualValues(t, 5, contador)
}

func TestIteradorUnicoElemento(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[string, int](compararStrings)
	dic.Guardar("Pucherito de gallina...", 100)

	contador := 0
	dic.Iterar(func(_ string, _ int) bool {
		contador++
		return true
	})

	require.EqualValues(t, 1, contador)
}

func TestIteradorRangoConDesdeYHastaIguales(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, string](compararEnteros)
	dic.Guardar(10, "A")
	dic.Guardar(20, "B")
	dic.Guardar(30, "C")

	desde := 20
	hasta := 20
	var resultado string
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		resultado = valor
		return true
	})

	require.EqualValues(t, "B", resultado)
}

func TestIteradorRangoVacio(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)
	for i := 0; i < 5; i++ {
		dic.Guardar(i, i)
	}

	desde := 10
	hasta := 20
	contador := 0
	dic.IterarRango(&desde, &hasta, func(_ int, _ int) bool {
		contador++
		return true
	})

	require.Equal(t, 0, contador)
}

func TestIteradorExternoOrdenCorrecto(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, string](compararEnteros)
	claves := []int{5, 3, 7, 1, 4}
	for _, c := range claves {
		dic.Guardar(c, fmt.Sprintf("%d", c))
	}

	esperado := []int{1, 3, 4, 5, 7}
	var resultado []int
	iter := dic.Iterador()
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
		iter.Siguiente()
	}

	require.EqualValues(t, esperado, resultado)
}

func TestIteradorExternoConRango(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)
	for i := 0; i < 10; i++ {
		dic.Guardar(i, i*2)
	}

	desde := 3
	hasta := 7
	iter := dic.IteradorRango(&desde, &hasta)
	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	require.EqualValues(t, []int{3, 4, 5, 6, 7}, claves)
}

func TestIteradorExternoBorradoPrevio(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[string, int](compararStrings)
	dic.Guardar("A", 1)
	dic.Guardar("B", 2)
	dic.Guardar("C", 3)
	dic.Borrar("B")

	iter := dic.Iterador()
	var claves []string
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	require.EqualValues(t, []string{"A", "C"}, claves)
}

func TestIteradorExternoModificarValores(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, *int](compararEnteros)
	valores := []int{1, 2, 3}
	for i := range valores {
		dic.Guardar(i, &valores[i])
	}

	iter := dic.Iterador()
	for iter.HaySiguiente() {
		_, valor := iter.VerActual()
		*valor *= 10
		iter.Siguiente()
	}

	require.EqualValues(t, 10, valores[0])
	require.EqualValues(t, 20, valores[1])
	require.EqualValues(t, 30, valores[2])
}

func TestIteradorVolumenABB(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)
	const N = 10000

	for i := 0; i < N; i++ {
		dic.Guardar(N-i, i)
	}

	iter := dic.Iterador()
	contador := 0
	for iter.HaySiguiente() {
		iter.Siguiente()
		contador++
	}

	require.Equal(t, N, contador)
}
