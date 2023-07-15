package models

type Map struct {
}

/*
Moves and Locations are indexed based on a hexagonal
coordinate system as follows:

				   (0,1,1)			  (0,0,1) \	     / (0,1,0)
				   -------					   \----/
		 (-1,0,1) /	      \ (1,1,0)			   /\  /\
				 /		   \	   (-1,0,0) __/__\/__\__ (1,0,0)
				 \		   /				  \	 /\  /
	    (-1,-1,0) \	      / (1,0,-1)	 	   \/  \/
		           -------					   /----\
		          (0,-1,-1)			 (0,-1,0) /		 \ (0,0,-1)
*/
type Location struct {
	x, y, z int8
}

type Move struct {
	x, y, z int8
}

func NorthMove() Move     { return Move{0, 1, 1} }
func NorthEastMove() Move { return Move{0, 1, 1} }

type Biome uint8

const (
	DesertBiome   Biome = 0
	WoodsBiome    Biome = 1
	PastureBiome  Biome = 2
	RockBiome     Biome = 3
	MountainBiome Biome = 4
	SeaBiome      Biome = 5
)

type Tile interface {
	Biome() Biome
	AdjacentTiles() []Tile
	Location() Location
}

type DesertTile struct {
	adjacentTiles []Tile
}

func (desert DesertTile) Biome() Biome {
	return DesertBiome
}

func (desert DesertTile) AdjacentTiles() []Tile {
	return desert.adjacentTiles
}

type WoodsTile struct {
	adjacentTiles []Tile
}

func (woods WoodsTile) Biome() Biome {
	return WoodsBiome
}

func (woods WoodsTile) AdjacentTiles() []Tile {
	return woods.adjacentTiles
}

type PastureTile struct {
	adjacentTiles []Tile
}

func (pasture PastureTile) Biome() Biome {
	return PastureBiome
}

func (pasture PastureTile) AdjacentTiles() []Tile {
	return pasture.adjacentTiles
}

type RockTile struct {
	adjacentTiles []Tile
}

func (rock RockTile) Biome() Biome {
	return RockBiome
}

func (rock RockTile) AdjacentTiles() []Tile {
	return rock.adjacentTiles
}

type MountainTile struct {
	adjacentTiles []Tile
}

func (mountain MountainTile) Biome() Biome {
	return MountainBiome
}

func (mountain MountainTile) AdjacentTiles() []Tile {
	return mountain.adjacentTiles
}

type SeaTile struct {
	adjacentTiles []Tile
}

func (sea SeaTile) Biome() Biome {
	return SeaBiome
}

func (sea SeaTile) AdjacentTiles() []Tile {
	return sea.adjacentTiles
}
