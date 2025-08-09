all: test

test:
	cd integration_test && bash tests.sh
