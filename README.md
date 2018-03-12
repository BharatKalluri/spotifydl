 
<h1 align="center">Spotifydl</h1>
<h4 align="center">A Downloader for Spotify Playlist's</h4>

----

Spotifydl is a spotify playlist downloader.

It uses youtube as the audio source and Spotify API for playlist details.

## Installation
Install golang and
```bash
git clone https://github.com/BharatKalluri/spotifydl
cd spotifydl
go install
```
and make sure you have `$GOPATH/bin` in your path.


## Usage

```bash
# help
spotifydl -h

# playlist and username must be supplied
# Suppose we consider an awesome playlist like waking the demon
# URL is https://open.spotify.com/user/22ywbwcqu2ci2rukqyhx7z5ha/playlist/13rgfJS9aI8PwfuDCaGJp0
# username will be 22ywbwcqu2ci2rukqyhx7z5ha and playlistid will be 13rgfJS9aI8PwfuDCaGJp0
spotifydl -username 22ywbwcqu2ci2rukqyhx7z5ha -playlistid 13rgfJS9aI8PwfuDCaGJp0
# Voila!!
```

If there are any Improvements or corrections that can be made, feel free to open an issue. I am still new to golang and would love to improve.

Note: This project was only done for as a learning experience for academic purposes. Usage of this product is up to the user and no responsibility will be taken for the user's action using this software.