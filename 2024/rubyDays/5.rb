ruleMap = Hash.new
possibleUpdates = []
File.foreach('5.txt') do |line|
  if line.include?("|")
    rule = line.chomp().split("|")
    # print rule[0], rule[1], "\n"
    if ruleMap.key?(rule[0])
      ruleMap[rule[0]].append(rule[1])
    else
      ruleMap[rule[0]] = [rule[1]]
    end
  elsif line.include?(",")
    possibleUpdates.append(line.chomp().split(","))
  end
end

#print ruleMap, "\n", possibleUpdates, "\n"

updates = []
badUpdates = []

for update in possibleUpdates
  brokenRule = false
  for i in 0..update.length - 1
    if ruleMap.key?(update[i])
      # print "TESTING - ", update[i], ",", ruleMap[update[i]], "\n"
      for y in ruleMap[update[i]]
        if update[0...i].include?(y)
          brokenRule = true
          break
        end
      end
    end
    break if brokenRule
  end
  updates.append(update) if !brokenRule
  badUpdates.append(update) if brokenRule
end

# print updates, "\nbadupdates:"
print ruleMap, "\n"
# print updates, "\n"
print badUpdates, "\n"

sum = 0
for update in updates
  middle = update.length / 2
  sum += update[middle].to_i
end

print sum, "\n"

map = ruleMap.dup
sorted = []
for i in 1..ruleMap.length
  map.each do |key,value|
    if map[key].length == i
      sorted.append(key)
    end
  end
end

sorted.unshift(ruleMap[sorted[0]][0]).reverse!

print "SORTED", sorted.length, "\n"

goodUpdates = []
for update in badUpdates
  goodUpdate = []
  for i in sorted
    if update.include?(i)
      goodUpdate.append(i)
    end
  end
  goodUpdates.append(goodUpdate)
end

# print goodUpdates, "\n"

sum = 0
for update in goodUpdates
  middle = update.length / 2
  sum += update[middle].to_i
end

# print sum, "\n"

# # 5548 is too high
goodUpdates = []
for update in badUpdates
  brokenRule = true
  until brokenRule == false
    brokenRule = false
    for i in 0...update.length
      if ruleMap.key?(update[i])
        for y in ruleMap[update[i]]
          for x in 0...i
            if y == update[x]
              brokenRule = true
              element = update.delete_at(x)
              update.insert(i, element)
              break
            end
          end
        end
      end
      break if brokenRule
    end
  end
  goodUpdates.append(update)
end

print "goodupdates\n", goodUpdates, "\n"

sum = 0
for update in goodUpdates
  middle = update.length / 2
  sum += update[middle].to_i
end

puts sum