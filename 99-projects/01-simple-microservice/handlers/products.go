package handlers

import (
	"log"
	"net/http"

	"avabodha.in/SimpleMicroservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
// 	lp := data.GetProducts()

// 	// marshel will buffer output to memory
// 	// and then we are writing to response.
// 	d, err := json.Marshal(lp)

// 	if err != nil {
// 		http.Error(rw, "Unable to marshal product list", http.StatusInternalServerError)
// 	}
// 	rw.Write(d)
// }

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("In getProducts")
	lp := data.GetProducts()

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal product list", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("In addProducts")

	// create product
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

	p.l.Printf("Prod: %#v", prod)
}
