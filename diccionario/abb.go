package diccionario


type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K,V]
	derecho *nodoAbb[K,V]
	clave	K
	dato	V
}

type abb[K comparable, V any] struct {
	raiz	*nodoAbb[K,V]
	cantidad	int
	cmp	func(K,K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K,V]{
	return &abb[K,V]{
		raiz: nil,
		cantidad: 0,
		cmp: funcion_cmp,
	}
}


func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K,V]{
	return &nodoAbb[K,V]{
		izquierdo: nil,
		derecho: nil,
		clave: clave,
		dato: dato,
	}
}  

func buscarReferencia[K comparable, V any](nodo **nodoAbb[K,V], clave K, cmp func(K,K) int) **nodoAbb[K,V]{
	if *nodo == nil || cmp((*nodo).clave, clave) == 0{
		return nodo
	}
	if cmp((*nodo).clave,clave) > 0{
		return buscarReferencia(&(*nodo).izquierdo, clave, cmp)
	}
	return buscarReferencia(&(*nodo).derecho, clave, cmp)
}

func buscarMaximoIzq[K comparable, V any](ref **nodoAbb[K, V]) **nodoAbb[K,V]{
	for (*ref).derecho != nil{
		ref = &((*ref).derecho)
	}
	return ref
}


func (ab *abb[K,V]) Pertenece(clave K) bool{
	return *(buscarReferencia(&(ab.raiz), clave, ab.cmp)) != nil
}


func (ab *abb[K,V]) Obtener(clave K) V{
	nodo := buscarReferencia(&(ab.raiz), clave, ab.cmp)
	if *nodo== nil{
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).dato
}


func (ab *abb[K,V]) Borrar(clave K) V{
	nodoPadre := buscarReferencia(&(ab.raiz), clave, ab.cmp)
	if *nodoPadre == nil{
		panic("La clave no pertenece al diccionario")
	} 

	nodoABorrar := *nodoPadre

	if nodoABorrar.izquierdo == nil && nodoABorrar.derecho == nil{
		*nodoPadre = nil
		ab.cantidad--
		return nodoABorrar.dato
	}

	if nodoABorrar.izquierdo == nil || nodoABorrar.derecho == nil{
		var hijo *nodoAbb[K,V]
		if nodoABorrar.izquierdo == nil{
			hijo = nodoABorrar.derecho
		}else{
			hijo = nodoABorrar.izquierdo
		}
		*nodoPadre = hijo
		ab.cantidad--
		return nodoABorrar.dato
	}
	
	nodoAnteriorRef := buscarMaximoIzq(&nodoABorrar.izquierdo)
	datoADevolver := nodoABorrar.dato
	nodoABorrar.clave = (*nodoAnteriorRef).clave
	nodoABorrar.dato = (*nodoAnteriorRef).dato

	if(*nodoAnteriorRef).izquierdo != nil{
		*nodoAnteriorRef = (*nodoAnteriorRef).izquierdo
	}else{
		*nodoAnteriorRef = nil
	}
	ab.cantidad--
	return datoADevolver
}

func (ab *abb[K,V]) Cantidad() int{
	return ab.cantidad
}

func (ab *abb[K,V]) Guardar(clave K, dato V){
	nodo := buscarReferencia(&(ab.raiz), clave, ab.cmp)
	if *nodo == nil{
		*nodo = crearNodo(clave, dato)
		ab.cantidad++
	}else{
		(*nodo).dato = dato
	}
	
}
