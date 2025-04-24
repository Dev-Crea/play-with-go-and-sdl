package contracts

type Contract struct{}

type Meta struct {
	Total int32
	Page  int32
	Limit int32
}
type Data struct {
	// Data Contract
	Meta Meta
}
