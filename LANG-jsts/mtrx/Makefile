
TSC      := tsc
TSCFLAGS := 

SRCS := mtrx4.ts
JSS := $(SRCS:.ts=.js)

.PHONY: clean all

all: compile

compile: mtrx4.js

mtrx4.js: mtrx4.ts
	$(TSC) $(TSCFLAGS) mtrx4.ts

clean:
	rm -f *.js


