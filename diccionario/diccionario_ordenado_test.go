package diccionario_test

import(
	"testing"
	"github.com/stretchr/testify/require"
	TDADiccionarioOrdenado "tdas/diccionario"
	"strings"
	"math/rand"
	"time"
)

func compararEnteros(a,b int) int{
	if a < b{
		return -1
	}else if a > b{
		return 1
	}
	return 0
}

func compararStrings(a, b string) int{
	return strings.Compare(a, b)
}

func TestDiccionarioOrdenadoVacio(t *testing.T){
	dicInt := TDADiccionarioOrdenado.CrearABB[int,int](compararEnteros)

	require.EqualValues(t, 0, dicInt.Cantidad())
	require.False(t, dicInt.Pertenece(21))

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicInt.Obtener(5)
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicInt.Borrar(4)
	})

	dicInt.Guardar(15,20)
	require.EqualValues(t, 1, dicInt.Cantidad())
	require.EqualValues(t, 20, dicInt.Obtener(15))

	require.True(t, dicInt.Pertenece(15))
	require.False(t, dicInt.Pertenece(20))

	require.EqualValues(t, 20, dicInt.Borrar(15))
	require.EqualValues(t, 0, dicInt.Cantidad())

	require.False(t, dicInt.Pertenece(15))

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicInt.Obtener(15)
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicInt.Borrar(121)
	})
}

func TestGuardarYBorrar(t *testing.T){
	dicStr := TDADiccionarioOrdenado.CrearABB[string,string](compararStrings)

	require.EqualValues(t, 0, dicStr.Cantidad())
	require.False(t, dicStr.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicStr.Borrar("")
	})

	dicStr.Guardar("A","B")
	dicStr.Guardar("L","w")
	dicStr.Guardar("s","l")

	require.EqualValues(t, 3, dicStr.Cantidad())
	require.True(t, dicStr.Pertenece("A"))
	require.True(t, dicStr.Pertenece("L"))
	require.True(t, dicStr.Pertenece("s"))

	require.EqualValues(t, "w", dicStr.Obtener("L"))
	require.EqualValues(t, "w", dicStr.Borrar("L"))
	require.False(t, dicStr.Pertenece("L"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicStr.Obtener("L")
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicStr.Borrar("L")
	})
	require.EqualValues(t, 2, dicStr.Cantidad())

	require.EqualValues(t, "B", dicStr.Obtener("A"))
	dicStr.Guardar("A", "H")

	require.True(t, dicStr.Pertenece("A"))
	require.EqualValues(t, "H", dicStr.Obtener("A"))
	require.EqualValues(t, 2, dicStr.Cantidad())
	require.EqualValues(t, "H", dicStr.Borrar("A"))

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicStr.Obtener("A")
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
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
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicStr.Obtener("y")
	})
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		dicStr.Borrar("s")
	})
	
}

func TestVolumenDiccionarioOrdenado(t *testing.T) {
	const N = 1000
	dicInt := TDADiccionarioOrdenado.CrearABB[int,int](compararEnteros)

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
	dicInt := TDADiccionarioOrdenado.CrearABB[int,int](compararEnteros)
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
    dic := TDADiccionarioOrdenado.CrearABB[int,int](compararEnteros)
    
    dic.Guardar(15,150)
    dic.Guardar(10,100)
    dic.Guardar(20,200)
    dic.Guardar(5,50)
    dic.Guardar(12,120)

    
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
	dic.Guardar(20, "b") // hijo derecho

	require.EqualValues(t, "a", dic.Borrar(10))
	require.True(t, dic.Pertenece(20))
	require.False(t, dic.Pertenece(10))
}

func TestBorrarNodoUnHijoIzquierdo(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, string](compararEnteros)
	dic.Guardar(20, "a")
	dic.Guardar(10, "b") // hijo izquierdo

	require.EqualValues(t, "a", dic.Borrar(20))
	require.True(t, dic.Pertenece(10))
	require.False(t, dic.Pertenece(20))
}


func TestStressDiccionario(t *testing.T) {
	const N = 10000
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)

	rand.Seed(time.Now().UnixNano())
	claves := rand.Perm(N) // genera una permutación aleatoria de 0 a N-1

	// Guardar claves con valor = clave * 10
	for _, clave := range claves {
		dic.Guardar(clave, clave*10)
	}

	require.EqualValues(t, N, dic.Cantidad())

	// Verificar que todas las claves estén
	for _, clave := range claves {
		require.True(t, dic.Pertenece(clave))
		require.EqualValues(t, clave*10, dic.Obtener(clave))
	}

	// Borrar todas las claves
	for _, clave := range claves {
		require.EqualValues(t, clave*10, dic.Borrar(clave))
		require.False(t, dic.Pertenece(clave))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
			dic.Obtener(clave)
		})
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() {
			dic.Borrar(clave)
		})
	}

	require.EqualValues(t, 0, dic.Cantidad())
}

func BenchmarkGuardar(b *testing.B) {
	dic := TDADiccionarioOrdenado.CrearABB[int, int](compararEnteros)
	for i := 0; i < b.N; i++ {
		dic.Guardar(i, i*10)
	}
}



//Test para el iterador que se me ocurrio 
/*func TestClavesStringsLargas(t *testing.T){
	dicStr := TDADiccionarioOrdenado.CrearABB(compararStrings)

	claves := []string{"a","aa","aaa","aaaa","aaaaa"}
	for _,clave := range claves{
		dic.Guardar(clave, len(clave))
	}

	var resultado []string
	dic.Iterar(clave string, dato int) bool{
		resultado = append(resultado, clave)
		return true
	}
	require.Equal(t, claves, resultado)
}*/
