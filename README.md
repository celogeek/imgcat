# imgcat

Display a remote image into the terminal.

## Build

Install Go Lang from: https://go.dev/

Then build

With make
```
make
```

Or directly
```
go build ./cmd/imgcat
```

## Usage

```
./imgcat -url "https://images.pexels.com/photos/2558605/pexels-photo-2558605.jpeg?cs=srgb&dl=pexels-anel-rossouw-2558605.jpg&fm=jpg" -height 25
```

<img width="1069" alt="Capture d’écran 2021-12-11 à 17 40 55" src="https://user-images.githubusercontent.com/65178/145684510-39617c22-5818-4573-896c-7e0e6915db91.png">

## Help

```
# imgcat -h
Usage of imgcat:
  -height int
    	maximum height in lines
  -url string
    	url of the jpg
```
