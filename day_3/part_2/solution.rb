require "awesome_print"
require "ostruct"

RE_CLAIM = /#(\d+) @ (\d+),(\d+): (\d+)x(\d+)/

input = []

# strings = [
#   "#1 @ 1,3: 4x4",
#   "#2 @ 3,1: 4x4",
#   "#3 @ 5,5: 2x2"
# ]
# strings.each do |line|
File.open("../input.txt", "r").each_line do |line|
  line.match(RE_CLAIM) do |m|
    input << OpenStruct.new(
      id: m[1],
      x: m[2].to_i,
      y: m[3].to_i,
      width: m[4].to_i,
      height: m[5].to_i
    )
  end
end

# 1000 x 1000 array filled with empty hashes
FABRIC = Array.new(1000) { Array.new(1000) { Hash.new() } }

# mark a section of FABRIC according to the claim
def mark(claim)
  FABRIC.slice(claim.y, claim.height).each do |f_y|
    f_y.slice(claim.x, claim.width).each do |inch|
      inch[claim.id] = true
    end
  end
end

input.each { |claim| mark(claim) }

def find_unique_claim
  invalidated_ids = {}
  possible_ids = {}

  FABRIC.each do |f_y|
    f_y.each do |inch|
      inch.keys.each do |id|
        # all ids in this inch are invalid, because they overlap
        if inch.keys.length > 1
            invalidated_ids[id] = true
            possible_ids.delete(id)
        else
          if invalidated_ids.has_key? id
            possible_ids.delete(id)
          else
            possible_ids[id] = true
          end
        end
      end
    end
  end

  return possible_ids.keys.first # there can be only one!
end

ap find_unique_claim
