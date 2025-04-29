package diccionario

import(
	TDALista	"tdas/lista"
	"fmt"
	"hash/fnv" //Una funcion de hash que prove go, si queres fijate en esta pagina. https://pkg.go.dev/hash/fnv#New64a
)

const(
	_FACTOR_CARGA = 0.7
	_CAPACIDAD_INICIAL = 7
)

type parClaveValor[K comparable, V any] struct{
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct{
	casillas 	[]TDALista.Lista[parClaveValor[K, V]]
	cantidad	int
}

//Lo podemos crear asi al hash

/*func CrearHash[K comparable, V any]() Diccionario[K, V]{
	return &hashAbierto[K, V]{
		casillas: nil,
		cantidad: 0,
	}
}*/

func CrearHash[K comparable, V any]() Diccionario[K, V]{
	casillas := make([]TDALista.Lista[parClaveValor[K,V]],_CAPACIDAD_INICIAL)
	for i := range casillas{
		casillas[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return &hashAbierto[K,V]{
		casillas: casillas,
		cantidad: 0,
	}
}
//Aca tengo la duda de si esta bien o no crear el hash vacio. Entiendo que no tiene mucho sentido porque en la interfaz en ningun lugar aclara que
//si se quiere hacer algo sobre el diccionario tire un panic en caso de que este vacio. Revisar.
func (h *hashAbierto[K,V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K,V]) Pertenece(clave K) bool{
	if h.Cantidad() == 0{
		return false
	}
	indice := h.indexDe(clave)
	lista := h.casillas[indice]
	if lista.EstaVacia(){
		return false
	}
	it := lista.Iterador()
	for it.HaySiguiente(){	
		if it.VerActual().clave == clave{
			return true
		}
		it.Siguiente()
	}	
	return false
}

func (h *hashAbierto[K,V]) Guardar(clave K, dato V){
	
}

func (h *hashAbierto[K,V]) Borrar(clave K) V{
	if !h.Pertenece(clave){
		panic("La clave no pertenece al diccionario")
	}
	return 2
}

func (h *hashAbierto[K,V]) Obtener(clave K) V{
	if !h.Pertenece(clave){
		panic("La clave no pertenece al diccionario")
	}
	indice := h.indexDe(clave)
	lista := h.casillas[indice]
	it := lista.Iterador()
	var valor V
	for it.HaySiguiente(){
		actual := it.VerActual()
		if actual.clave == clave{
			valor = actual.valor
			break
		}
		it.Siguiente()
	}
	return valor
}

// indexDe define en que casilla del arreglo de listas enlazadas debe caer el par clave-valor.
func (h *hashAbierto[K,V]) indexDe(clave K) int{
	hv := h.hashClave(clave)
	return int(hv % uint64(len(h.casillas)))
}

// hashClave usa FNV-1a para generar un hash de la clave.
func (h *hashAbierto[K,V]) hashClave(clave K) uint64 {
	hf := fnv.New64a()
	hf.Write(convertirABytes(clave))
	return hf.Sum64()
}

// convertirABytes convierte una clave generica a []byte.
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
