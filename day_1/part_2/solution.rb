frequencies = []

File.open("../input.txt", "r").each_line do |line|
  frequencies << line.chomp.to_i
end

total = 0
series = {}

while true do
  frequencies.each do |reading|
    if series[total]
      puts total # found duplicate
      exit
    else
      series[total] = true
    end

    total += reading
  end
end

puts "None found"
