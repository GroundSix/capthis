CapThis
=======

CapThis allows you to give it an image, some text along with
various other options and then it overlays that text onto
the image, like a caption.

It runs with a HTTP interface, a server is started and
image / text data is then sent to allow CapThis to
process everything and output the final result.

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
$ git clone URL
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

MIT