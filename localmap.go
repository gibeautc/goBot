package main

const tileSize = 100

type tile struct {
	data [][]byte
}

func newTile() *tile {
	var self *tile
	self.data = make([][]uint8, tileSize) // initialize a slice of dy slices
	for i := 0; i < len(self.data); i++ {
		self.data[i] = make([]uint8, len(self.data)) // initialize a slice of dx unit8 in each of dy slices
	}
	return self
}

//func (self *tile) toBmp() *image.Image{

//}
