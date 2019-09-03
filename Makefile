all: test

install_ictclas:
	sudo cp ictclas/include/*.h /usr/include
	if [ -e /usr/lib64 ]; then sudo cp ictclas/lib/linux64/*.so /usr/lib64/; else sudo cp ictclas/lib/linux64/*.so /usr/lib/; fi

install_sent:
	sudo cp sent/include/*.h /usr/include
	if [ -e /usr/lib64 ]; then sudo cp sent/lib/linux64/*.so /usr/lib64/; else sudo cp sent/lib/linux64/*.so /usr/lib/; fi

test:
	go test -v
