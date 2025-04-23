package lista

type NodoLista[T any] struct {
	proximo *NodoLista[T]
	dato    T
}

type listaEnlazada[T any] struct { // Creo que acá debería ir con minúscula para que no sea expor.
	primero *NodoLista[T]
	ultimo  *NodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] { // Esto debe devolver la interfaz
	return &ListaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
		largo:   0,
	}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {

}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {

}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

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
	act := lista.primero
	for act != nil {
		if !visitar(act.dato) {
			return
		}
		act = act.proximo
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {

}
