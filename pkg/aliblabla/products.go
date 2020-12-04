package aliblabla

type IProductRepository interface {
	Get(id int) (product, bool)
}

type inMemoryProductRepository struct {
	Products map[int]product
}

func (r *inMemoryProductRepository) Get(id int) (product, bool) {
	p, found := r.Products[id]
	return p, found
}

type product struct {
	ID    int
	Name  string
	Price float64
}
