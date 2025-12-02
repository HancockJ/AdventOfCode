def day1_1()
  a,b = get_lists()
  total_diff = 0
  for i in 0..a.length
    total_diff += (a[i].to_i - b[i].to_i).abs
  end
  total_diff
end

def day1_2()
  a,b = get_lists()
  total_similarity = 0
  for x in a
    multiplier = 0
    for y in b
      if x == y
        multiplier += 1
      end
    end
    total_similarity += x * multiplier
  end
  total_similarity
end

def get_lists()
  a, b = [], []
  File.foreach('day1_input') do |line|
    tmp = line.chomp.split("   ")
    a << tmp[0].to_i
    b << tmp[1].to_i
  end
  a,b = a.sort, b.sort
end


print "Day 1 part 1", "\n"
puts day1_1()
print "Day 1 part 2", "\n"
puts day1_2()
