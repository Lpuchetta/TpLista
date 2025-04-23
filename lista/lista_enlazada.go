package lista

<<<<<<< HEAD
type nodoLista[T any] struct{
	proximo *nodoLista[T]
	dato T
}

type listaEnlazada[T any] struct{
	primero *nodoLista[T]
	ultimo *nodoLista[T]
	largo int
}

type iteradorListaEnlazada[T any] struct{
	actual *nodoLista[T]
	anterior *nodoLista[T] 
	lista *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() listaEnlazada[T]{
	return listaEnlazada[T]{
		primero : nil,
		ultimo : nil,
		largo : 0,
	}
}

func (l *listaEnlazada[T]) Iterador() iteradorListaEnlazada[T]{
	return iteradorListaEnlazada[T]{
		actual : l.primero,
		anterior : nil,
		lista : l,
	}
}

func (it *iteradorListaEnlazada[T]) HaySiguiente() bool{
	return it.actual != nil
}

func (it *iteradorListaEnlazada[T]) VerActual() T{
	if !it.HaySiguiente(){
		panic ("EL iterador termino de iterar")
	}
	return it.actual.dato	
}

func (it *iteradorListaEnlazada[T]) Siguiente(){
	if !it.HaySiguiente(){
		panic ("El iterador termino de iterar")
	}
	it.anterior = it.actual
	it.actual = it.actual.proximo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool){
	for actual := l.primero; actual != nil; actual = actual.proximo{
		if !visitar(actual.dato){
			break
		}
	}
}

=======
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
	return &listaEnlazada[T]{
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
>>>>>>> 2ad60d452e71246c64487a97c8f697fda4669cd2
