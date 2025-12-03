def day2_1()
  safe = 0
  File.foreach('test') do |line|
    current = line.split(' ').map(&:to_i)
    a = safe(current, 0)
    puts a
    if a
      safe += 1
    end
  end
  return safe
end

def isSafe(current)
  return false if depth > 1
  unsafe = 0
  if current[0] > current[1] && current[0] <= current[1] + 3
    for i in 1..current.length - 2
      if !(current[i] > current[i + 1] && current[i] <= current[i + 1] + 3)
        return false
      end
    end
  elsif current[0] < current[1] && current[0] + 3 >= current[1]
    for i in 1..current.length - 2
      if !(current[i] < current[i + 1] && current[i] + 3 >= current[i + 1])
        return false
      end
    end
  else
    return false
  end
  return true
end

def safe(current, depth)
  return false if depth > 1
  unsafe = 0
  if current[0] > current[1] && current[0] <= current[1] + 3
    for i in 1..current.length - 2
      if !(current[i] > current[i + 1] && current[i] <= current[i + 1] + 3)
        print current, arrayPop(current, i), arrayPop(current, i + 1), "\n"
        return (safe(arrayPop(current, i), depth + 1) || safe(arrayPop(current, i + 1), depth + 1))
      end
    end
  elsif current[0] < current[1] && current[0] + 3 >= current[1]
    for i in 1..current.length - 2
      if !(current[i] < current[i + 1] && current[i] + 3 >= current[i + 1])
        print current, arrayPop(current, i), arrayPop(current, i + 1), "\n"
        return (safe(arrayPop(current, i), depth + 1) || safe(arrayPop(current, i + 1), depth + 1))
      end
    end
  else
    print current, arrayPop(current, 0), arrayPop(current, 1), "\n"
    return (safe(arrayPop(current, 0), depth + 1) || safe(arrayPop(current, 1), depth + 1))
  end
  return true
end

def arrayPop(array, index)
  newArr = array.dup
  newArr[0...index] + newArr[index + 1..-1]
end


# frozen_string_literal: true

module Day2
  module Part2
    def self.run('test', _)
      reports = []
      FileReader.for_each_line(path) do |line|
        reports << line.split.map(&:to_i)
      end
      safe_count = 0
      reports.each do |levels|
        new_reports = []
        levels.each_with_index do |_, idx|
          new_reports << levels.reject.with_index { |_, i| i == idx }
        end

        new_reports.each do |new_levels|
          if safe_levels(new_levels)
            safe_count += 1
            break
          end
        end
      end
      safe_count
    end

    def self.safe_levels(levels)
      return false if levels.uniq != levels
      return false unless levels.sort == levels || levels.sort.reverse == levels

      safe = true
      levels.sort.each_cons(2).to_a.each do |f, s|
        if s - f > 3
          safe = false
          break
        end
      end
      safe
    end
  end
end


puts day2_1()