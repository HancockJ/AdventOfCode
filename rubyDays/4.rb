def p1
  board = []

  File.foreach('4.txt') do |line|
    board.append(line.chomp().chars)
  end

  window = 4
  count = 0
  
  # horizontal
  for row in board
    for x in 0..row.length-window
      tmpStr = row[x..x+window-1].join()
      if tmpStr == "XMAS" || tmpStr == "SAMX"
        count += 1
      end
    end
  end

  # vertical
  for x in 0..board[0].length-1
    for y in 0..board[0].length-window
      tmpStr = board[y][x] + board[y + 1][x] + board[y + 2][x] + board[y + 3][x]
      if tmpStr == "XMAS" || tmpStr == "SAMX"
        count += 1
      end
    end
  end

  # diagnol 
  for x in 0..board[0].length-window
    for y in 0..board[0].length-window
      tmpStr = board[y][x] + board[y + 1][x + 1] + board[y + 2][x + 2] + board[y + 3][x + 3]
      if tmpStr == "XMAS" || tmpStr == "SAMX"
        count += 1
      end
    end
  end

  # diagnol - reverse
  (0..board[0].length - window).each do |x|
    (3..board.length - 1).each do |y|
      tmpStr = board[y][x] + board[y - 1][x + 1] + board[y - 2][x + 2] + board[y - 3][x + 3]
      if tmpStr == "XMAS" || tmpStr == "SAMX"
        count += 1
      end
    end
  end
end

def p2
  board = []

  File.foreach('4.txt') do |line|
    board.append(line.chomp().chars)
  end

  window = 3
  count = 0

  (0..board[0].length - window).each do |x|
    (0..board.length - window).each do |y|
      tmpA = board[y][x] + board[y + 1][x + 1] + board[y + 2][x + 2]
      tmpB = board[y + 2][x] + board[y + 1][x + 1] + board[y][x + 2]
      print tmpA, ",", tmpB, "\n"
      if (tmpA == "MAS" || tmpA == "SAM") && (tmpB == "MAS" || tmpB == "SAM")
        count += 1
      end
    end
  end
  puts count
end

p1()
p2()