CA=abs_server_ca
ABS=abs_test
new_abs:
	@go build -o new_abs new_abs.go lagRange.go define.go
abs:
	@go build -o ${ABS} abs.go lagRange.go define.go
ca:
	@go build -o ${CA} ca.go lagRange.go abs.go define.go

clean:
	@rm -f ${CA} nohup.out
	@rm -f new_abs nohup.out

all: ca