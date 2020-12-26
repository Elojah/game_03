### Space

## Operations

# Tile

ID
TS as redis score

Fetch tile first entities at max TS T
FetchManyTiles(X range, Y range)

world.TileSize
world.NTileX
world.NTileY

world:X:Y:tileID

TileSize = 10
[skill example]
center = tile[22, 35] XY[3, 6]
distance = 25
=> fetch all tiles in range
xmin = 3 - 25 = -22 => -22 / 10 => -2.3 => -3
     = centerTileX + (((centerX - distance) / tileSize) - 1)
     = 22 - 3 = 19
xmax = (25 - (10 - 3)) / 10 + 1 = (1.8) = +2
     = centerTileX + (((centerX + distance) / tileSize))
     = 22 + 2 = 24
ymin = ((6 - 25) / 10) - 1 = -2
     = 35 - 2 = 33
ymax = 35 + ((6 + 25) / 10)
     = 35 + 3 = 38

worldid:x:y:*score*
entityid

if prevposition.section !=  newposition.section {
     sectionID:TS:entityID:in
     sectionID:TS:entityID:out
}

fetch sectionID:EntityID:in/out

sectionID:out:entityID:*score*
sectionID:in:entityID:*score*

sectionID:entityID:*score* => add a field state in/out

fetch sectionID:* at score max X
filter state out
