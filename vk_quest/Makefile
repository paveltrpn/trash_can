
CC = g++-10
CFLAGS = --std=c++20

SRCS := vk_quest.cpp
OBJS := $(SRCS:.cpp=.o)

.phony: all clean

all: main

vk_quest.o: vk_quest.cpp
	$(CC) $(CFLAGS) -c vk_quest.cpp

main: $(OBJS)
	$(CC) $(OBJS) -o $@ -lstdc++ -larmadillo

clean:
	rm -f *.o main