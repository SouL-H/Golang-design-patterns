package main


import "fmt"

type Notebook interface {
	insertLightingPort()
}

type Client struct{}

func (c *Client) ConnectByLighting(n Notebook) {
	fmt.Println("Start to insert to lighting port")
	n.insertLightingPort()
}

type Mac struct{}

func (m *Mac) insertLightingPort() {
	fmt.Println("Start connecting to Mac")
}

type Windows struct{}

func (w *Windows) insertTypeCPort() {
	fmt.Println("Start connecting to Windows")
}

type TypeCAdapterForWindows struct {
	w *Windows
}

func (a *TypeCAdapterForWindows) insertLightingPort() {
	fmt.Println("Transform lighting to type c")
	a.w.insertTypeCPort()
}

func main() {
	// mac
	c := &Client{}
	m := &Mac{}
	c.ConnectByLighting(m)

	// windows
	w := &Windows{}
	wa := &TypeCAdapterForWindows{w}
	c.ConnectByLighting(wa)
}