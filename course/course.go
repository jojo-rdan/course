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
	name    string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}

// Estos son setters en GO
func (c *course) SetName(name string) {
	c.name = name
}
func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}
func (c *course) SetUserIDs(userIDs []uint) {
	c.userIDs = userIDs
}
func (c *course) SetPrice(price float64) {
	c.price = price
}

// Este es un gettter con GO
func (c *course) Name() string {
	return c.name
}
func (c *course) Price() float64 {
	return c.price
}
func (c *course) IsFree() bool {
	return c.isFree
}
func (c *course) UserIDs() []uint {
	return c.userIDs
}

// Esto es un constructor para cursos
// con otras validaciones en GO
func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
