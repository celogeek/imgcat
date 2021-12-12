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
go build -o imgcat
```

## Usage

```
./imgcat -u "https://images.pexels.com/photos/2558605/pexels-photo-2558605.jpeg?cs=srgb&dl=pexels-anel-rossouw-2558605.jpg&fm=jpg" -r 25
```

<img width="978" alt="Capture d’écran 2021-12-12 à 19 01 37" src="https://user-images.githubusercontent.com/65178/145723965-1a299d51-5cfb-42d4-b2b2-6fc97a715391.png">


## Help

```
Usage:
  imgcat [OPTIONS]

Application Options:
  -u, --url=  url of the jpg
  -r, --rows= maximum number of rows

Help Options:
  -h, --help  Show this help message
```
