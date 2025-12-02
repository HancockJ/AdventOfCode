data = File.read('3.txt')

regex = /mul\(\d{1,3},\d{1,3}\)|don\'t\(\)|do\(\)/

matches = data.scan(regex)

totalSum = 0
enabled = true
for match in matches
  if match == "do()"
    enabled = true
  elsif match == "don't()"
    enabled = false
  else
    if enabled
      nums = match.scan(/\d+/).map(&:to_i)
      totalSum += nums[0] * nums[1]
    end
  end
end

puts totalSum