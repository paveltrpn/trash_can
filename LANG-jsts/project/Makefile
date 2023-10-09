CC       := tsc
CFLAGS   := -t ES2016 -m 'es2015' --removeComments

SRCS := src/main.ts \
		src/utils.ts \
		src/common.ts
		
OBJS := main.js \
		utils.js \
		common.js
		
#OBJS := $(SRCS:.ts=.js)

.PHONY: clean all

all: $(OBJS)

common.js: src/common.ts
	$(CC) $(CFLAGS) --outDir . src/common.ts

utils.js: src/utils.ts
	$(CC) $(CFLAGS) --outDir . src/utils.ts

main.js: src/main.ts
	$(CC) $(CFLAGS) --outDir . src/main.ts

clean:
	rm -f *.js
