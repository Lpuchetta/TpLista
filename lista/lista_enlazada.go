package lista

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

