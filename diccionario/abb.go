package diccionario

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:     nil,
		cantidad: 0,
		cmp:      cmp,
	}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		izquierdo: nil,
		derecho:   nil,
		clave:     clave,
		dato:      dato,
	}
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	return ab.pertenece(ab.raiz, clave)
}

func (ab *abb[K, V]) Obtener(clave K) V {
	return ab.obtener(ab.raiz, clave)
}

func (ab *abb[K, V]) Borrar(clave K) V {

}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	ab.raiz = ab.guardar(ab.raiz, clave, dato)
}

func (ab *abb[K, V]) Iterar(visitar func(K, V) bool) {

}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (abb[K, V]) Iterador() iterDiccionario[K, V] {

}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

}

func (ab *abb[K, V]) obtener(nodo *nodoAbb[K, V], clave K) V {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	comparacion := ab.cmp(clave, nodo.clave)
	switch {
	case comparacion < 0:
		return ab.obtener(nodo.izquierdo, clave)
	case comparacion > 0:
		return ab.obtener(nodo.derecho, clave)
	default:
		return nodo.dato
	}
}

func (ab *abb[K, V]) guardar(nodo *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if nodo == nil {
		ab.cantidad++
		return crearNodo(clave, dato)
	}
	comparacion := ab.cmp(clave, nodo.clave)
	switch {
	case comparacion < 0:
		nodo.izquierdo = ab.guardar(nodo.izquierdo, clave, dato)
	case comparacion > 0:
		nodo.derecho = ab.guardar(nodo.derecho, clave, dato)
	default:
		nodo.dato = dato
	}
	return nodo
}

func (ab *abb[K, V]) pertenece(nodo *nodoAbb[K, V], clave K) bool {
	if nodo == nil {
		return false
	}
	comparacion := ab.cmp(clave, nodo.clave)
	switch {
	case comparacion < 0:
		return ab.pertenece(nodo.izquierdo, clave)
	case comparacion > 0:
		return ab.pertenece(nodo.derecho, clave)
	}
	return true
}
