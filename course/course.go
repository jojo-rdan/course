package course

import "fmt"

// Encapsulamiento son las herramientas que provee un lenguaje
// para protejer las propiedades o métodos del usuario
// que está trabajando con la clase
// a diferencia de Java, los identificadores se dividen
// en expotados y no exportados

// un identificador exportado es donde
// su primer caracter sea en mayúscula
type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: isFree,
	}
}

func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

func (c *course) ChangePrice(price float64) {
	c.Price = price
}
