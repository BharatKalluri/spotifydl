 
<h1 align="center">Spotifydl</h1>
<h4 align="center">A Music Downloader for Spotify</h4>

----

![Spotifydl Demo](spotifydl.gif)

Spotifydl is a spotify playlist downloader.

It uses youtube as the audio source and Spotify API for playlist details.

## Installation
Make sure you have golang, youtube-dl and ffmpeg installed.
```bash
go get github.com/BharatKalluri/spotifydl
```
Make sure you have python 3 with pip installed on your system. Instructions to install can be found [here](https://rg3.github.io/youtube-dl/download.html)
and make sure you have `$GOPATH/bin` in your path.


## Usage

```bash
# help

spotifydl -p 13rgfJS9aI8PwfuDCaGJp0
# Voila!!
```

If there are any Improvements or corrections that can be made, feel free to open an issue. I am still new to golang and would love to improve.

Note: This project was only done for as a learning experience for academic purposes. Usage of this product is up to the user and no responsibility will be taken for the user's action using this software.
