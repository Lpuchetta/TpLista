package cola_prioridad

const (
	_CAPACIDAD_INICIAL  = 23
	_FACTOR_REDIMENSION = 2
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, _CAPACIDAD_INICIAL),
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heapify(arreglo, funcion_cmp)
	return &colaConPrioridad[T]{
		datos: arreglo,
		cant:  len(arreglo),
		cmp:   funcion_cmp,
	}
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == 0
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {
	if cola.cant == len(cola.datos) {
		cola.redimensionar(len(cola.datos) * _FACTOR_REDIMENSION)
	}
	cola.datos[cola.cant] = dato
	upHeap(cola.datos, cola.cant, cola.cmp)
	cola.cant++

}

func (cola *colaConPrioridad[T]) VerMax() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.datos[0]
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := cola.datos[0]
	cola.cant--
	cola.datos[0] = cola.datos[cola.cant]
	downHeap(cola.datos, cola.cant, 0, cola.cmp)

	if cola.cant > _CAPACIDAD_INICIAL && cola.cant < len(cola.datos)/4 {
		cola.redimensionar(len(cola.datos) / _FACTOR_REDIMENSION)
	}
	return max
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

func upHeap[T any](arr []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}

	posPadre := (pos - 1) / 2
	if cmp(arr[pos], arr[posPadre]) > 0 {
		arr[posPadre], arr[pos] = arr[pos], arr[posPadre]
		upHeap(arr, posPadre, cmp)
	}
}

func downHeap[T any](arr []T, tam int, pos int, cmp func(T, T) int) {
	if pos == tam {
		return
	}

	posIzq := 2*pos + 1
	posDer := 2*pos + 2
	posMayor := pos

	if posIzq < len(arr) && cmp(arr[posIzq], arr[posMayor]) > 0 {
		posMayor = posIzq
	}

	if posDer < len(arr) && cmp(arr[posDer], arr[posMayor]) > 0 {
		posMayor = posDer
	}

	if posMayor != pos {
		arr[pos], arr[posMayor] = arr[posMayor], arr[pos]
		downHeap(arr, tam, posMayor, cmp)
	}

}

func heapify[T any](arr []T, cmp func(T, T) int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		downHeap(arr, len(arr), i, cmp)
	}
}

func (cola *colaConPrioridad[T]) redimensionar(nuevoTam int) {
	nuevos := make([]T, nuevoTam)
	copy(nuevos, cola.datos[:cola.cant])
	cola.datos = nuevos
}

func posMinimo[T any](arr []T, cmp func(T, T) int, a, b, c int) int {
	pos := a
	if cmp(arr[b], arr[pos]) < 0 {
		pos = b
	}
	if cmp(arr[c], arr[pos]) < 0 {
		pos = c
	}
	return pos
}
