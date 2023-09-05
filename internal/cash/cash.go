package cash

import "level0/internal/postgre"

var Cash map[string]string = make(map[string]string)

// Init cash.
func Init() {
	postgre.GetData(Cash)
}
