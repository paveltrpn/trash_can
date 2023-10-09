
TSC      := tsc
TSCFLAGS := 

SRCS := brasenham.ts
JSS := $(SRCS:.ts=.js)

.PHONY: clean all

all: compile

compile: brasenham.js

brasenham.js: brasenham.ts
	$(TSC) $(TSCFLAGS) brasenham.ts

clean:
	rm -f *.js


