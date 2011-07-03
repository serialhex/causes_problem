# trying my hand at a makefile... needs work
include $(GOROOT)/src/Make.inc

TARG=social_distance
GOFILES=\
	wordsort.go\
	social_distance.go\
	#editdist.go\

# only one of the following
#include $(GOROOT)/src/Make.pkg
include $(GOROOT)/src/Make.cmd