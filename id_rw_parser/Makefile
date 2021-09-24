
CC       := g++
FLAGS    := -Wall -std=c++17 
LIBS 	 := -lstdc++
RM       := rm -f

SRCS := main.cpp
OBJS := $(SRCS:.cpp=.o)

$(info COMMON MAKEFILE)

all: compile link 

compile:
	$(CC) $(FLAGS) -c $(SRCS)

link:
	$(CC) $(OBJS) -o main $(LIBS)

clean:
	$(RM) *.o
	$(RM) main
