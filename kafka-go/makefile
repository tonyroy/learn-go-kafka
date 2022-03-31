bc:
	go build -o out/consumer consumer/main.go

bp:
	go build -o out/producer producer/main.go

rc:
	out/consumer

rp:
	out/producer

consumer: bc rc

producer: bp rp