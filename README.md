# tinygo-xiao

TinyGo demos and examples on Seeedstudio XIAO-ESP32C3 and XIAO-ESP32S3.

## blinky

Blinks an LED. The "Hello, World" of things.

### xiao-esp32c3

```
tinygo flash -target xiao-esp32c3 -size short ./blinky
```

### xiao-esp32s3

```
tinygo flash -target xiao-esp32s3 -size short ./blinky
```

## button

Push a button, and the LED lights up.

### xiao-esp32c3

```
tinygo flash -target xiao-esp32c3 -size short ./button
```

### xiao-esp32s3

```
tinygo flash -target xiao-esp32s3 -size short ./button
```

## echo

Type into the console, and the Xiao will echo back what you typed.

### xiao-esp32c3

```
tinygo flash -target xiao-esp32c3 -size short -monitor ./echo
```

### xiao-esp32s3

```
tinygo flash -target xiao-esp32s3 -size short -monitor ./echo
```

## display

Shows the xiao controlling an OLED display with an I2C interface


### xiao-esp32c3

```
tinygo flash -target xiao-esp32c3 -size short ./display
```

### xiao-esp32s3

```
tinygo flash -target xiao-esp32s3 -size short ./display
```

## conway

Shows the xiao controlling an OLED display with an I2C interface playing Conway's Game of Life

### xiao-esp32c3

```
tinygo flash -target xiao-esp32c3 -size short ./life
```

### xiao-esp32s3

```
tinygo flash -target xiao-esp32s3 -size short ./life
```

## scanner

Scans for WiFi access points and displays them on the OLED display.


### xiao-esp32c3

```
tinygo flash -target xiao-esp32c3 -size short ./scanner
```

### xiao-esp32s3

```
tinygo flash -target xiao-esp32s3 -size short ./scanner
```
