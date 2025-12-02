equations = Hash.new
$combo_history = Hash.new

File.foreach('7.txt') do |line|
  a = line.chomp.split(":")
  equations[a[0].to_i] = a[1].split(" ").map(&:to_i)
end

def combos(i)
  return $combo_history[i] if $combo_history.key?(i)

  # Base case: when i == 1, return ["+", "*"]
  return [["+"], ["*"], ["||"]] if i == 1

  # Recursive case: build combinations for length i
  # Get all combinations for i - 1 and append "+" or "*" to each combination
  smaller_combos = combos(i - 1)
  a = smaller_combos.map { |inner| inner + ['+'] }
  b = smaller_combos.map { |inner| inner + ['*'] }
  c = smaller_combos.map { |inner| inner + ['||'] }

  # Cache and return the combinations in correct order
  $combo_history[i] = a + b + c
  return $combo_history[i]
end

full_sum = 0
equations.each do |sum, numbers|
  combinations = combos(numbers.length)
  for combination in combinations
    tempSum = numbers[0]
    for i in 1...numbers.length
      if combination[i - 1] == "+"
        tempSum = numbers[i] + tempSum
      elsif combination[i - 1] == "*"
        tempSum = numbers[i] * tempSum
      elsif combination[i-1] == "||"
        tempSum = (tempSum.to_s + numbers[i].to_s).to_i
      end
    end

    if tempSum == sum
      full_sum += tempSum
      break
    end
  end
end


print full_sum, "\n"
