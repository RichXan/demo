proto_pay_file = pay/proto/pay.proto
proto_order_file = order/proto/order.proto

proto-pay:
	protoc --go_out=. --go-grpc_out=. $(proto_pay_file)

proto-order:
	protoc --go_out=. --go-grpc_out=. $(proto_order_file)