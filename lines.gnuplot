set xlabel 'iteration'

set ylabel 'time (micorseconds)'

plot "results.dat" using 1:2 with lines,\
"results.dat" using 1:3 with lines