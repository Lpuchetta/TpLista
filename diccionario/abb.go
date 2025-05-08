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

func avanzarHacia[K comparable, V any](nodo **nodoAbb[K,V], clave K, cmp func(K, K) int) **nodoAbb[K,V]{
	if cmp(clave, (*nodo).clave) < 0 {
		return &((*nodo).izquierdo)
	}else if cmp(clave, (*nodo).clave) > 0{
		return &((*nodo).derecho)
	}
	return nodo 
}

func buscarReferencia[K comparable, V any](nodo **nodoAbb[K,V], clave K, cmp func(K,K) int) **nodoAbb[K,V]{
	for *nodo != nil{
		if cmp(clave, (*nodo).clave) == 0{
			return nodo
		}
		nodo = avanzarHacia(nodo, clave, cmp)
	}
	return nodo
}

func buscarMaximoIzq[K comparable, V any](nodo **nodoAbb[K, V]) **nodoAbb[K,V]{
	for (*nodo).derecho != nil{
		nodo = (&(*nodo).derecho)
	}
	return nodo
}


func (ab *abb[K,V]) Obtener(clave K) V{
	nodoPadre := buscarReferencia(&(ab.raiz), clave, ab.cmp)
	if *nodoPadre == nil{
		panic("La clave no pertenece al diccionario")
	}
	return (*nodoPadre).dato
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
	
	nodoAnterior := buscarMaximoIzq(&(*nodoABorrar).izquierdo)
	datoADevolver := (*nodoABorrar).dato
	(*nodoABorrar).clave = (*nodoAnterior).clave
	(*nodoABorrar).dato = (*nodoAnterior).dato

	if(*nodoAnterior).izquierdo != nil{
		*nodoPadre = (*nodoAnterior).izquierdo
	}else{
		*nodoPadre = nil
	}
	
	ab.cantidad--
	return datoADevolver
}

func (ab *abb[K,V]) Cantidad() int{
	return ab.cantidad
}

func (ab *abb[K,V]) Guardar(clave K, dato V){
	
	return
}

