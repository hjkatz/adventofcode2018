total = 0
File.open("../input.txt", "r").each_line do |line|
  total += line.chomp.to_i
end

puts total
