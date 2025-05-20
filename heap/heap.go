package cola_prioridad

const (
	_CAPACIDAD_INICIAL  = 5
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
	return nil
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == 0
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {
	if cola.cant == len(cola.datos) {
		cola.redimensionar(len(cola.datos) * _FACTOR_REDIMENSION)
	}
	cola.datos[cola.cant] = dato
	cola.upHeap(cola.cant)
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
	// TODO: Completar
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

func (cola *colaConPrioridad[T]) upHeap(pos int) {
	if pos == 0 {
		return
	}

	posPadre := posicionPadre(pos)
	padre := cola.datos[posPadre]
	hijo := cola.datos[pos]

	if cola.cmp(padre, hijo) > 0 {
		swap(cola.datos, posPadre, pos)
		cola.upHeap(posPadre)
	}
}

// TODO: Completar
// IDEA:
//  1. Se elimina al primero del arreglo y se considera al último como el primero.
//  2. Llamo a downHeap para ese elemento.
//     2.a) Calcular la pos de ambos hijos.
//     2.b) Se pregunta si se cumple la condición de heap. Si se cumple, termina; caso contrario,
//     se realiza el swap entre el padre e hijo mayor y se repite el paso 2.
func downHeap() {

}

func posicionPadre(posHijo int) int {
	return (posHijo - 1) / 2
}

func posicionHijoIzq(posPadre int) int {
	return 2*posPadre + 1
}

func posicioHijoDer(posPadre int) int {
	return 2*posPadre + 2
}

func (cola *colaConPrioridad[T]) redimensionar(nuevoTam int) {
	nuevos := make([]T, nuevoTam)
	copy(nuevos, cola.datos)
	cola.datos = nuevos
}

func swap[T any](arr []T, pos, otra int) {
	arr[pos], arr[otra] = arr[otra], arr[pos]
}
