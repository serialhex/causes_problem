
=begin info

  module Levenshtein
  calculates Levenshtein distance between 2 strings or arrays
  
  mostly copied from:
  http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#Ruby
  though it's been refactored to make it more object-like

=end

module Levenshtein
	def levenshtein other
	  case
	    when self.empty? then other.length
	    when other.empty? then self.length
	    else [(self[0] == other[0] ? 0 : 1) + self[1..-1].levenshtein(other[1..-1]),
	          1 + self[1..-1].levenshtein(other),
	          1 + self.levenshtein(other[1..-1])].min
	  end
	end
	alias :edit_distance :levenshtein
  alias :edit_dist :levenshtein
end

# including the module in string, to make life easier...
String.class_eval do
  include Levenshtein
end

# and even though i don't need it here i'll include it in array also, cause i can!
Array.class_eval do
  include Levenshtein
end
