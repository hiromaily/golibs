bld:
	go build -i -v -o ./cmdtool ./main.go

run1:
	./cmdtool -s aaaaa -n 100 -b true

run2:
	./cmdtool -s bbbb -n 999 -b false

run3:
	./cmdtool -s bbbb -n 999

all: bld run1 run2 run3