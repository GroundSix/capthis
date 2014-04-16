CapThis
=======

[![Build Status](https://travis-ci.org/GroundSix/capthis.svg?branch=master)](https://travis-ci.org/GroundSix/capthis)

CapThis allows you to give it an image, some text along with
various other options and then it overlays that text onto
the image, like a caption.

It runs with a HTTP interface, a server is started and
image / text data is then sent to allow CapThis to
process everything and output the final result.

This project is still in development and while it may
work currently, there is still a lot of work to be done
that will soon allow you to configure the way it works
in a lot more detail.

Using an image of a [sunset](https://www.flickr.com/photos/47515486@N05/9100133053/in/photolist-eS9yvK-8WrzYk-fUoFsz-ci8XH-9KeCSN-B3TY2-5RtKw-9qgjB2-g47ReA-g4zyk-2fRFt-fB3Q3T-9MLCEm-5saEMh-8t7YjD-7hDXrE-zmUFy-4zySqx-75xGxj-6dQHRg-4d3L2q-BidzN-5rVtQn-7a49ax-f3WSf7-4tnesA-5qSQgJ-4pR5GY-fh3ZNb-X6BnQ-kwUZd-7qJxop-7zjGac-9nLt95-7Yu691-7Ly4hp-kssmtv-7TmKcw-9isz3v-7ondJK-56Fgwg-JodDf-6kvZAh-7UMuQg-5pj8xf-5ggqda-5GRYNP-5CYFmQ-6Dhbv7-j8MvF/)
and a [cartoon styled font](http://www.1001freefonts.com/from_cartoon_blocks.font) CapThis generated this:

![An example](https://raw.githubusercontent.com/GroundSix/capthis/master/screenshots/example.png)

### Prerequisites

You will need:

  - Git
  - Go (versions 1.1 and 1.2)
  - Make
  - Imagick

#### Installing Imagick

##### OS X

Homebrew

```bash
$ brew install imagemagick
```

MacPorts

```bash
$ sudo port install ImageMagick
```

##### Ubuntu / Debian

```bash
$ sudo apt-get install libmagickwand-dev
```

### Installing

```bash
$ git clone https://github.com/GroundSix/capthis.git
$ cd capthis
$ make
$ sudo make install
```

### Running

You can start the server like so:

```bash
$ capthis
```

By default CapThis will run on port `8080`. You can change this by specifying:

```bash
$ capthis 2020
```

#### Http API

At the moment CapThis by default runs on port 8080.
Below are the post fields you must populate:

  - `image_name` - The full path to the original image you need to process
  - `text` - The text you are going to overlay the image with
  - `font` - The full path to a ttf font file of your choice
  - `font_size` - The font size (int) in pixels
  - `fill_color` - The fill colour of your text (e.g. 'black')
  - `stroke_color` - The stroke color of your text (e.g. 'black')
  - `output` - The full path of the output (e.g. 'my_image.png')

An example request:

Headers

  - **URL:** http://localhost:8080/
  - **Accept:** */*
  - **Accept-Encoding:** gzip, deflate, compress
  - **Content-Length:** 179
  - **Content-Type:** application/x-www-form-urlencoded

Paramaters

  - **fill_color:** black
  - **font:** /path/to/my/font.ttf
  - **font_size:** 32
  - **image_name:** /path/to/my/image.png
  - **output:** /path/to/output.png
  - **stroke_color:** black
  - **text:** Hello, World!

```
image_name=%2Fpath%2Fto%2Fmy%2Fimage.png&text=Hello%2C+World%21&fill_color=black&font_size=32&font=%2Fpath%2Fto%2Fmy%2Ffont.ttf&output=%2Fpath%2Fto%2Foutput.png&stroke_color=black
```

### Tests

```bash
$ make test
```

### License

[MIT](https://github.com/GroundSix/capthis/blob/master/LICENSE)