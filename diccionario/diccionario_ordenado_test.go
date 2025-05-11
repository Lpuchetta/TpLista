package diccionario

import(
	"testing"
	"github.com/stretchr/testify/require"
	TDADiccionarioOrdenado "tdas/diccionario"
	"strings"
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
	return strings.compare(a, b)
}

func TestDiccionarioVacio(t *testing.T){
	dicInt := TDADiccionarioOrdenado.CrearAbb(compararEnteros)

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
	dicStr := TDADiccionarioOrdenado.CrearABB[compararStrings]

	require.EqualValues(t, 0, dicStr.Cantidad())
	require.False(t, dicStr.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func(){
		disStr.Borrar("")
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

	require.True(t, dicStr.Pertence("A"))
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

func TestVolumen(t *testing.T) {
	const N = 1000
	dicInt := TDADiccionarioOrdenado.CrearABB(compararEnteros)

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
	dicInt := TDADiccionarioOrdenado.CrearABB(compararEnteros)
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
