input = []
File.open("../input.txt", "r").each_line do |line|
  input << line.chomp
end

def count_chars(str)
  found = {}
  str.chars.each do |c|
    found[c] = 0 unless found.has_key? c
    found[c] += 1
  end

  counts = {}
  found.values.each do |count|
    counts[count] = true
  end

  return counts
end

def checksum(input)
  totals = {
    2 => 0,
    3 => 0,
  }

  input.each do |id|
    counts = count_chars(id)

    totals[2] += 1 if counts[2]
    totals[3] += 1 if counts[3]
  end

  return totals[2] * totals[3]
end

puts checksum(input)
