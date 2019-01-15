#!/bin/env python3
#
import sys
#
#  Separate a list of numbers into primes and composites.
#  Primes are identified using sieve of Eratosthenes
# 
#  Usage:  ./composite 15 17 101 189 ... 
#        It will print: 
#          Composite: [15, 189]
#          Prime: [17, 101]
#
#-------------------------------------------
# 
# sieve(lo, hi) returns list of primes from lo to hi, computed by sieving.
#
# lo and hi are further proof of the dislike of full-word-typing: think low and high
#
def sieve(lo,hi):
#You will need to implement the Sieve of Eratosthenes 
#which finds all the primes in a range of numbers
# https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes explains it
# A good way to start is to build a list of all the numbers 
# tip: "range" is a tool in python/Google is an OK resource for the language!
#ANSWER: below and 
# FILL IN THIS SECTION!!!!!!!!!! of code.
   primes = []
   nums = range(lo, hi+1)
   for x in nums:
      prime = True
      for y in range(2, int(hi/2)):
         if(x % y == 0 and x!=y):
            prime = False
            break
      if(prime):
         primes.append(x)

   return primes

# The goal is to:
# separate a set of numbers into primes and composites
#
#  given (set-of-numbers)
#
#  return ( set-of-composites, set-of-primes )
#
def separate(numbs):
    primes = set(sieve(min(numbs), max(numbs)))
    return (numbs.difference(primes), numbs.intersection(primes))

#--------------------------------------------
#
# Main program:
#
#  Extract numbers from argument list.
#  Use lowest and highest ones as argument to seive
#  Must be at least 1 number in the arguments (arguments are passed at the command line!)
#

if len(sys.argv)>=2:
#Uhoh  What is that Strange 1: action?
#Answer: it is only taking the items from index 1 onwards, instead of capturing the file name in index 0 as well
    numbs = sys.argv[1:]
    # QUESTION FOR YOU: What sort of thing is "numbs"?
    #ANSWER: an array of the input arguments (as strings) from the command line
    numbs = [int(x) for x in numbs  if x.isdigit()]
    if len(numbs)>0:
        comp, prim = separate(set(numbs))
        #So why does this print work?
        #TODO: ANSWER: Because python helps print arrays nicely                     MORE HERE?
        print("Composite: ", sorted(comp))
        print("Prime: ", sorted(prim))
        sys.exit()
#  
# usage
#
print("Prints which numbers are prime and which are composite.")
print("Usage: composite n1 n2 n3 ...")
print("example: composite 3 5 7 9 11 13 15")