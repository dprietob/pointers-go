package main

import "fmt"

func main() {
	var v int = 4
	var p *int

	p = &v
	*p = 8

	incValor( v )
	incReferencia( &v )

	fmt.Printf( "El valor de v es %d \n", v )
	fmt.Printf( "La posición de memoria de v es %v \n", &v )

    fmt.Printf( "El valor de p es %d \n", *p )
    fmt.Printf( "La posición de memoria de p es %v \n", p )
}

func incValor( v int ) {
	v++
	fmt.Printf( "El valor incrementado por valor es %d \n", v )
}

func incReferencia( v *int  ) {
	*v++;
	fmt.Printf( "El valor incrementado por referencia es %d \n", *v  )
}
