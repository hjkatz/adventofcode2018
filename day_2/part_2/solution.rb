require "awesome_print"

input = []
File.open("../input.txt", "r").each_line do |line|
  input << line.chomp
end

def string_differ(a, b)
  differ = 0
  common = ""

  (0..a.length - 1).each do |i|
    if a[i] == b[i]
      common << a[i]
    else
      differ += 1
    end
  end

  return differ, common
end

def find_common_id(input)
  input.combination(2) do |pair|
    amount, common = string_differ(*pair)
    return common if amount == 1
  end
end

puts find_common_id(input)
