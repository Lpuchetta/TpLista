// Heap de mínimopackage diccionario
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

type iterAbb[K comparable, V any] struct {
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

func buscarReferencia[K comparable, V any](nodo **nodoAbb[K, V], clave K, cmp func(K, K) int) **nodoAbb[K, V] {
	if *nodo == nil || cmp((*nodo).clave, clave) == 0 {
		return nodo
	}
	if cmp((*nodo).clave, clave) > 0 {
		return buscarReferencia(&(*nodo).izquierdo, clave, cmp)
	}
	return buscarReferencia(&(*nodo).derecho, clave, cmp)
}

func buscarMaximoIzq[K comparable, V any](ref **nodoAbb[K, V]) **nodoAbb[K, V] {
	for (*ref).derecho != nil {
		ref = &((*ref).derecho)
	}
	return ref
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	return *(buscarReferencia(&(ab.raiz), clave, ab.cmp)) != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	nodo := buscarReferencia(&(ab.raiz), clave, ab.cmp)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).dato
}

func (ab *abb[K, V]) Borrar(clave K) V {
    ref := buscarReferencia(&ab.raiz, clave, ab.cmp)
    if *ref == nil {
        panic("La clave no pertenece al diccionario")
    }

    nodo := *ref
    datoOriginal := nodo.dato

    // 1) Caso hoja
    if nodo.izquierdo == nil && nodo.derecho == nil {
        *ref = nil
        ab.cantidad--
        return datoOriginal
    }

    // 2) Un único hijo
    if nodo.izquierdo == nil || nodo.derecho == nil {
        var hijo *nodoAbb[K, V]
        if nodo.izquierdo == nil {
            hijo = nodo.derecho
        } else {
            hijo = nodo.izquierdo
        }
        *ref = hijo
        ab.cantidad--
        return datoOriginal
    }

    // 3) Dos hijos: reemplazo por predecesor
    preRef := buscarMaximoIzq(&nodo.izquierdo)
    claveReemplazo := (*preRef).clave
    datoReemplazo := (*preRef).dato

    // Borro recursivamente el predecesor (ya decrementa ab.cantidad)
    ab.Borrar(claveReemplazo)

    // Sustituyo clave y dato en el nodo original
    nodo.clave = claveReemplazo
    nodo.dato = datoReemplazo

    // Devuelvo el dato del nodo “borrado”
    return datoOriginal
}


func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	nodo := buscarReferencia(&(ab.raiz), clave, ab.cmp)
	if *nodo == nil {
		*nodo = crearNodo(clave, dato)
		ab.cantidad++
	} else {
		(*nodo).dato = dato
	}

}

func (ab *abb[K, V]) Iterar(visitar func(K, V) bool) {
	ab.iterarRango(ab.raiz, nil, nil, visitar)
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(K, V) bool) {
	ab.iterarRango(ab.raiz, desde, hasta, visitar)
}

func (ab *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(K, V) bool) bool {
	if nodo == nil {
		return true
	}

	continuo := ab.iterarRango(nodo.izquierdo, desde, hasta, visitar)
	if !continuo {
		return false
	}

	estaEnRango := true
	if desde != nil && ab.cmp(nodo.clave, *desde) < 0 {
		estaEnRango = false
	}

	if hasta != nil && ab.cmp(nodo.clave, *hasta) > 0 {
		estaEnRango = false
	}

	if estaEnRango {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	return ab.iterarRango(nodo.derecho, desde, hasta, visitar)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return ab.IteradorRango(nil, nil)
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
		cmp:   ab.cmp,
	}

	iter.apilarIzquierdos(ab.raiz)
	iter.avanzarHastaRango()

	return iter
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pila.VerTope()
	return actual.clave, actual.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := iter.pila.Desapilar()
	iter.apilarIzquierdos(actual.derecho)
	iter.avanzarHastaRango()
}

func (iter *iterAbb[K, V]) apilarIzquierdos(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		iter.pila.Apilar(nodo)
		nodo = nodo.izquierdo
	}
}

func (iter *iterAbb[K, V]) avanzarHastaRango() {
	for !iter.pila.EstaVacia() {
		nodo := iter.pila.VerTope()

		if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
			iter.pila.Desapilar()
			iter.apilarIzquierdos(nodo.derecho)
			continue
		}

		if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
			iter.pila.Desapilar()
			continue
		}

		break
	}
}
