all:
	gcc -c ./gpio/gpio.c -o ./gpio/gpio.o
	ar rcs ./gpio/libgpio.a ./gpio/gpio.o
	ranlib ./gpio/libgpio.a

clean:
	rm ./gpio/gpio.o
	rm ./gpio/libgpio.a