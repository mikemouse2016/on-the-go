package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type Product struct {
	Id, Title string
}

type ProductResource struct {
	// DAO
}

func (p ProductResource) getOne(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	log.Println("getting product with id:" + id)
	resp.WriteEntity(Product{Id: id, Title: "test"})
}

func (p ProductResource) postOne(req *restful.Request, resp *restful.Response) {
	updatedProduct := new(Product)
	err := req.ReadEntity(updatedProduct)
	if err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	log.Println("updateing product with id:" + updatedProduct.Id)
}

func (p ProductResource) Register() {
	ws := new(restful.WebService)
	ws.Path("/products")
	ws.Consumes(restful.MIME_XML)
	ws.Produces(restful.MIME_XML)

	ws.Route(ws.GET("/{id}").To(p.getOne).
		Doc("get the product by its id").
		Param(ws.PathParameter("id", "identifier of the product").DataType("string")))

	ws.Route(ws.POST("").To(p.postOne).
		Doc("update or create a product").
		Param(ws.BodyParameter("Product", "a Product (XML)").DataType("main.Product")))

	restful.Add(ws)
}

func main() {
	ProductResource{}.Register()
	http.ListenAndServe(":8080", nil)
}
