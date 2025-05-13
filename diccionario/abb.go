package diccionario

import TDAPila "tdas/pila"

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

type iterABB[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   func(K, K) int
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
	ab.iterar(ab.raiz, visitar)
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	ab.iterarRango(ab.raiz, desde, hasta, visitar)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return ab.IteradorRango(nil, nil)
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterABB[K, V]{
		pila: TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
	}
	iter.apilarIzquierdosRango(ab.raiz, desde, hasta, ab.cmp)
	return iter
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

func (ab *abb[K, V]) iterar(nodo *nodoAbb[K, V], visitar func(K, V) bool) {
	ab.iterarRango(nodo, nil, nil, visitar)
}

func (ab *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(K, V) bool) {
	if nodo == nil {
		return
	}

	cmpDesde := -1
	if desde != nil {
		cmpDesde = ab.cmp(nodo.clave, *desde)
	}

	cmpHasta := 1
	if hasta != nil {
		cmpHasta = ab.cmp(nodo.clave, *hasta)
	}

	if desde == nil || cmpDesde >= 0 {
		ab.iterarRango(nodo.izquierdo, desde, hasta, visitar)
	}

	enRango := true
	if desde != nil && cmpDesde < 0 {
		enRango = false
	}

	if hasta != nil && cmpHasta > 0 {
		enRango = false
	}

	if enRango {
		if !visitar(nodo.clave, nodo.dato) {
			return
		}
	}

	if hasta == nil || cmpHasta <= 0 {
		ab.iterarRango(nodo.derecho, desde, hasta, visitar)
	}

}

func (it *iterABB[K, V]) HaySiguiente() bool {
	return !it.pila.EstaVacia()
}

func (it *iterABB[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	tope := it.pila.VerTope()
	return tope.clave, tope.dato
}

func (it *iterABB[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := it.pila.Desapilar()
	it.apilarIzquierdosRango(nodo.derecho, it.desde, it.hasta, it.cmp)
}

func (it *iterABB[K, V]) apilarIzquierdosRango(
	nodo *nodoAbb[K, V],
	desde *K,
	hasta *K,
	cmp func(K, K) int,
) {
	actual := nodo
	for actual != nil {
		cmpDesde := -1
		if desde != nil {
			cmpDesde = cmp(actual.clave, *desde)
		}

		cmpHasta := 1
		if hasta != nil {
			cmpHasta = cmp(actual.clave, *hasta)
		}

		if hasta != nil && cmpHasta > 0 {
			actual = actual.izquierdo
			continue
		}

		if desde != nil && cmpDesde < 0 {
			actual = actual.derecho
			continue
		}

		it.pila.Apilar(actual)
		actual = actual.izquierdo
	}
}
