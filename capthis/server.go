/**
 * capthis v0.1-dev
 *
 * (c) Ground Six
 *
 * @package capthis
 * @version 0.1-dev
 * 
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 * 
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package capthis

import (
    "log"
    "strconv"
    "net/http"
)

/**
 * @var string server port
 */
type ServerConfig struct {
    port string
}

/**
 * Server constructor that assigns the
 * port
 *
 * @param string server port
 *
 * @return ServerConfig
 */
func Server(port string) *ServerConfig {
    server := new(ServerConfig)
    server.port = ":" + port

    return server
}

/**
 * Starts the web server that will
 * accept image request data
 *
 * @return nil
 */
func (s ServerConfig) StartServer() {
    http.HandleFunc("/", requestHandler)
    http.ListenAndServe(s.port, nil)
}

/**
 * Makes call to then process image data
 * after http request has been made
 *
 * @param http.ResponseWriter response writer
 * @param http.Request request data
 *
 * @return nil
 */
func requestHandler(w http.ResponseWriter, r *http.Request) {
    populateCaptionData(r)
}

/**
 * Handles the http POST request and
 * sends the image data off to be
 * processed
 *
 * @param http.Request request data
 *
 * @return nil
 */
func populateCaptionData(r *http.Request) {
    caption := New()

    fontSize, err := strconv.ParseFloat(r.PostFormValue("font_size"), 64)
    if err != nil {
        log.Fatal(err)
    }

    caption.SetName(r.PostFormValue("image_name"))
    caption.SetText(r.PostFormValue("text"))
    caption.SetFont(r.PostFormValue("font"))
    caption.SetFontSize(fontSize)
    caption.SetFillColor(r.PostFormValue("fill_color"))
    caption.SetStrokeColor(r.PostFormValue("stroke_color"))
    caption.SetNewImageName(r.PostFormValue("output"))

    caption.ProcessImage()
}
