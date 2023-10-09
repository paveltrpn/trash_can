
CC = g++-10
CFLAGS = --std=c++20

SRCS := chainline.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: chainline

chainline.o: chainline.cpp
	$(CC) $(CFLAGS) -c chainline.cpp

main: chainline.o 
	$(CC) $(OBJS) -lstdc++ -lplplotcxx -o main 

clean:
	rm -f *.o main
