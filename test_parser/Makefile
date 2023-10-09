
CC       := clang++
CFLAGS   := -Wall -std=c++17 -target x86_64-pc-windows-gnu
LFLAGS 	 := -lstdc++

SRCS := test_parser.cpp
OBJS := $(SRCS:.cpp=.o)

.PHONY: clean all

all: $(OBJS) main

main: test_parser.cpp
	$(CC) $(CFLAGS) test_parser.cpp -o main

clean:
	rm -f *.o main
 