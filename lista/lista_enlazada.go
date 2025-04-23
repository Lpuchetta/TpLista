package lista

type NodoLista[T any] struct {
	proximo *NodoLista[T]
	dato    T
}

type listaEnlazada[T any] struct {
	primero *NodoLista[T]
	ultimo  *NodoLista[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct {
	actual   *NodoLista[T]
	anterior *NodoLista[T]
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
		largo:   0,
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{
		actual:   l.primero,
		anterior: nil,
		lista:    l,
	}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo := &NodoLista[T]{
		dato:    dato,
		proximo: lista.primero,
	}
	if lista.EstaVacia() {
		lista.ultimo = nuevo
	}
	lista.primero = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {

}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.proximo
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil; actual = actual.proximo {
		if !visitar(actual.dato) {
			break
		}
	}
}

func (it *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return it.actual != nil
}

func (it *iteradorListaEnlazada[T]) VerActual() T {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return it.actual.dato
}

func (it *iteradorListaEnlazada[T]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	it.anterior = it.actual
	it.actual = it.actual.proximo
}

func (it *iteradorListaEnlazada[T]) Insertar(T) {

}

func (it *iteradorListaEnlazada[T]) Borrar() T {

}
