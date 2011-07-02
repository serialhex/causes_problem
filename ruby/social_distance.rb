#!/usr/bin/env ruby

=begin social distance & thoughts
  
  * from:
  http://jobs.github.com/positions/181fb26e-71cc-11e0-96a9-8752f87b91a0
  jobs@causes.com
	
	Two words are friends if they have a Levenshtein distance of 1. That is, you can add, remove, or substitute exactly one letter in word X to create word Y. A word’s social network consists of all of its friends, plus all of their friends, and all of their friends’ friends, and so on. Write a program to tell us how big the social network for the word “causes” is, using this word list

  this is my implementation of social distance for causes.com

  * helpers...
  http://en.wikipedia.org/wiki/Levenshtein_distance

  ---

  so, i figure i have a couple of options...
  
  1) i can either write an algo that computes the Levenshtein distance of two strings, and then just generate an infinite number of strings, and take only the strings that are found in the dictionary...
  - this is just wrong and stupid... horribly inefficient and just dumb (though brute force always works, i know i can do better!)
  
  2) i can iteratively generate strings from a primary string, only changing 1 feature, and seeing if the generated string is in the dictionary.
  - this is probably very search intensive, i mean, i'll be searching through the dict for every string i generate, instead of simply working on the dict itself...  this was the primary reason for thinking of having the dict be symbols instead of strings, but i may change that... 

  3) i can write an implementation of Levenshtein distance and then iterate through the dict looking for words where string.levenshtein_distance(str) == 1.  then iterate though *those* words and do the same and so on, and so forth...
  - this is better, though maybe i'll think of something even better!!

  - Justin Patera aka serialhex

=end

require './levenshtein'
require 'yaml'

# throwing them in a hash which should be faster to search than an array, and i can keep track of which words point to a given word, so i dont have infinite feh => meh => feh => meh crud
# maybe the above, i dunno... fer now an array
puts "creating dictionary..."
dict = []
File.open("word.list", 'r') do |f|
  f.each_line do |word| # because each line only has one word
    dict << word.chomp
  end
end

# set the network to it's root: 'causes'
# using a hash cause it should be faster than an array for searching & stuff
network = {'causes' => {}}

=begin algorithm
  
  so, after some refelction, i have figured that the best way to go about it is thus:

  1) get the Levenshtein distance of all the words in dict from 'causes'
  2) from the words that are friends with 'causes', find their friends from all words that are twice removed from 'causes', and from those words, and the words thrice removed from causes and so on and so forth... (i know, it's meta, and you couldn't use this to do the same kind of thing for people, but i'm not using people i'm using words!)
  3) PROFIT!

=end

there = Time.now
counter = 0
puts 'distancifyin yo!' 
distances = Hash.new{|hsh, key| hsh[key] = Array.new}
distances[0] = 'causes' # 'causes' has a 0 Levenshtein distance from 'causes' :D
dict.each do |word|
  counter += 1
  print '.' if (counter % 10 == 0)
  distances[word.edit_dist 'causes'] << word
end

puts
puts "finished in: #{Time.now - there}"

File.open('socials.yml', 'w') {|f| YAML.dump(distances, f)}