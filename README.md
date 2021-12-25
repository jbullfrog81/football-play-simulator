# football-play-simulator
Program to simulate American Football plays

## Dependencies


### [Pixel](https://github.com/faiface/pixel)
[Pixel](https://github.com/faiface/pixel) is the backbone and has a backend called [PixelGL](https://godoc.org/github.com/faiface/pixel/pixelgl) backend which uses OpenGL to render graphics. So you'll need OpenGL development libraries. More details on Pixel [requirements](https://github.com/faiface/pixel#requirements) before you can compile unless you want to see the error message[1]. I was able to run `go get -u github.com/go-gl/glfw/v3.3/glfw` then `go mod vendor` to pull the dependency without manually building glfw with CMake.

[1]
```
# github.com/go-gl/glfw/v3.3/glfw
vendor/github.com/go-gl/glfw/v3.3/glfw/c_glfw.go:4:10: fatal error: 'glfw/src/context.c' file not found
#include "glfw/src/context.c"
         ^~~~~~~~~~~~~~~~~~~~
1 error generated.
```