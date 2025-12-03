$guardDirection = {'>' => 'V', 'V'=>'<', '<'=>'^', '^'=>'>'}
$guardMove = {'>' => [1,0], 'V'=>[0,1], '<'=>[-1,0], '^'=>[0,-1]}

def printMap(map)
  for row in map
    print row.join(), "\n"
  end
end

def makeMove(map, x, y)
  newX, newY = $guardMove[map[y][x]][0] + x, $guardMove[map[y][x]][1] + y
  return map, -1, -1 if !(newX >= 0 && newX < map.length && newY >= 0 && newY < map[x].length)
  if map[newY][newX] == "#"
    newX, newY =  $guardMove[$guardDirection[map[y][x]]][0] + x, $guardMove[$guardDirection[map[y][x]]][1] + y
    return map, -1, -1 if !(newX >= 0 && newX < map.length && newY >= 0 && newY < map[x].length)
    map[newY][newX] = $guardDirection[map[y][x]]
  else
    map[newY][newX] = map[y][x]
  end

  map[y][x] = "X"

  return map, newX, newY
end

def countMoves(map, x, y, max)
  count = 0
  while x != -1 && y != -1 && count < max
    map, x, y = makeMove(map,x,y)
    if map == nil
      break
    end
    count += 1
  end
  return count
end

# CREATE ORIGINAL MAP AND LOCATION
map = []
x, y = 0,0

File.open('6.txt', 'r') do |file|
  line = []
  posX,posY = 0,0
  file.each_char do |char|
    if char == "\n"
      map.append(line)
      line = []
      posX = 0
      posY += 1
    else
      line.append(char)
      x, y = posX, posY if $guardDirection.key?(char)
      posX += 1
    end
  end
  map.append(line)
end

# print "guard at: ", x, ",", y, "\n"
# printMap(map)
maxMoves = 10000

totalSpots = 0
for row in 0...map.length
  for column in 0...map[row].length
    if map[row][column] != "#" && !$guardDirection.key?(map[row][column])
      newMap = map.map(&:dup)
      newMap[row][column] = "#"
      totalSpots += 1 if countMoves(newMap, x, y, maxMoves) >= maxMoves - 10
      # puts countMoves(newMap, x, y, maxMoves)
    end
  end
end
print "total: ", totalSpots, "\n"