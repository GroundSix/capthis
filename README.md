CapThis
=======

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

### Prerequisites

You will need:

  - Git
  - Go
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

**NOTE**: The above process will be changing soon

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

### Tests

To be written...

### License

[MIT](https://github.com/GroundSix/capthis/blob/master/LICENSE)