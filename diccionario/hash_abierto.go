package diccionario

import(
	TDALista	"tdas/lista"
)

const _CAPACIDAD_INICIAL = 16

type hashAbierto[K comparable, V any] struct{
	casillas 	[]TDALista.Lista[K]
	cantidad	int
}

