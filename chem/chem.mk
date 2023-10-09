
SRCS	:= test.cpp simple.cpp polynomial.cpp
OBJS	:= $(SRCS:.cpp=.o)
OBJS	+=  $(ALGEBRA_CPP)/mtrx234.o\
			$(ALGEBRA_CPP)/vec234.o\
			$(ALGEBRA_CPP)/qtnn.o\
			$(ALGEBRA_CPP)/plane.o

all: compile build

compile: simple.o polynomial.o test.o

simple.o: simple.cpp
	$(CC) $(CFLAGS) -I$(ALGEBRA_CPP)/ -c simple.cpp 

polynomial.o: polynomial.cpp
	$(CC) $(CFLAGS) -I$(ALGEBRA_CPP)/ -c polynomial.cpp 

test.o: test.cpp
	$(CC) $(CFLAGS) -I$(ALGEBRA_CPP)/ -c test.cpp

build:
	$(CC) $(OBJS) -o main -lstdc++