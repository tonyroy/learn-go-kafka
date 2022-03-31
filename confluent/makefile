bc:
	go build -o out/consumer consumer/main.go

bp:
	go build -o out/producer producer/main.go 

rc:
	out/consumer $(config)

rp:
	out/producer $(config)

consumer: bc rc 

producer: bp rp 

# e.g. make consumer config=./config_local.properties