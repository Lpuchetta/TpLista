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
