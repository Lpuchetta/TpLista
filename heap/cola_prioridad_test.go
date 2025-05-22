package cola_prioridad_test

import (
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int { return a - b })

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())

	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarUnicoElemento(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	heap.Encolar(100)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 100, heap.VerMax())
	require.Equal(t, 100, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapifyArregloExistente(t *testing.T) {
	arr := []int{9, 3, 7, 1, 5}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })

	require.Equal(t, 9, heap.VerMax())

	expectedOrder := []int{9, 7, 5, 3, 1}
	for _, expected := range expectedOrder {
		require.Equal(t, expected, heap.Desencolar())
	}
}

func TestInsercionOrdenInverso(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	for i := 1; i <= 10; i++ {
		heap.Encolar(i)
	}

	for i := 10; i >= 1; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
}

func TestNumerosNegativos(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	heap.Encolar(-5)
	heap.Encolar(-10)
	heap.Encolar(-3)

	require.Equal(t, -3, heap.VerMax())
	require.Equal(t, -3, heap.Desencolar())
	require.Equal(t, -5, heap.Desencolar())
	require.Equal(t, -10, heap.Desencolar())
}

func TestElementosIguales(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	elementos := []int{5, 5, 5, 5}
	for _, e := range elementos {
		heap.Encolar(e)
	}

	for range elementos {
		require.Equal(t, 5, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5}
	TDAHeap.HeapSort(arr, func(a, b int) int { return a - b })

	expected := []int{1, 1, 3, 4, 5}
	require.Equal(t, expected, arr)
}


func TestHeapArrVacio(t *testing.T) {
	arr := []int{}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	
}

func TestHeapInverso(t *testing.T) {
	cmp := func(a,b int) int { return b - a }
	heap := TDAHeap.CrearHeap[int](cmp)
	heap.Encolar(3);heap.Encolar(1);heap.Encolar(2)
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
}

func TestHeapifyMinHeap(t *testing.T) {
	arr := []int{9, 3, 7, 1, 5}
	cmp := func(a, b int) int { return b - a }
	heap := TDAHeap.CrearHeapArr(arr, cmp)

	require.Equal(t, 1, heap.VerMax())
	expectedOrder := []int{1, 3, 5, 7, 9}
	for _, expected := range expectedOrder {
		require.Equal(t, expected, heap.Desencolar())
	}
}


func TestIntercalado(t *testing.T) {
	cmp := func(a,b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)
	heap.Encolar(5);heap.Encolar(2)
	require.Equal(t, 5, heap.VerMax())
	heap.Desencolar()
	heap.Encolar(7)
	require.Equal(t, 7, heap.VerMax())
}


type Persona struct { Nombre string; Edad int }
func TestHeapPersona(t *testing.T) {
    cmp := func(a, b Persona) int { return a.Edad - b.Edad }
    heap := TDAHeap.CrearHeap[Persona](cmp)
    personas := []Persona{{"Ana",30}, {"Luis",25}, {"Eva",35}}
    for _, p := range personas { heap.Encolar(p) }
    require.Equal(t, Persona{"Eva",35}, heap.VerMax())
}

func TestVolumen(t *testing.T) {
    const n = 10000
    cmp := func(a, b int) int { return a - b }
    heap := TDAHeap.CrearHeap[int](cmp)

    for i := 0; i < n; i++ {
        heap.Encolar(i)
        require.Equal(t, i, heap.VerMax())
        require.Equal(t, i+1, heap.Cantidad())
        require.False(t, heap.EstaVacia())
    }

    require.Equal(t, n, heap.Cantidad())

    for expected := n - 1; expected >= 0; expected-- {
        currMax := heap.VerMax()
        require.Equal(t, expected, currMax)
        
        require.Equal(t, expected+1, heap.Cantidad())

       
        got := heap.Desencolar()
        require.Equal(t, expected, got)

        
        require.Equal(t, expected, heap.Cantidad())

        
        if !heap.EstaVacia() {
            require.True(t, heap.VerMax() <= currMax)
        }
    }

    require.True(t, heap.EstaVacia())
}

func TestMuchosIgualesYUnDistinto(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	for i := 0; i < 1000; i++ {
		heap.Encolar(42)
	}
	heap.Encolar(999)

	require.Equal(t, 999, heap.VerMax())
	require.Equal(t, 1001, heap.Cantidad())

	require.Equal(t, 999, heap.Desencolar())
	for i := 0; i < 1000; i++ {
		require.Equal(t, 42, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestIntercaladoVolumen(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	for i := 0; i < 5000; i++ {
		heap.Encolar(i)
		if i%3 == 0 {
			heap.Desencolar()
		}
	}
	for !heap.EstaVacia() {
		heap.Desencolar()
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapRandom(t *testing.T) {
	cmp := func(a, b int) int { return a - b }
	heap := TDAHeap.CrearHeap[int](cmp)

	input := []int{23, 5, 76, 3, 9, 88, 1, 44}
	for _, v := range input {
		heap.Encolar(v)
	}

	prev := heap.Desencolar()
	for !heap.EstaVacia() {
		curr := heap.Desencolar()
		require.LessOrEqual(t, curr, prev)
		prev = curr
	}
}
