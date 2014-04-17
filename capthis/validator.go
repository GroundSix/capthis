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
    "encoding/json"
    "net/http"
)

/**
 * Reports and returns any errors if they
 * were found while validating the post
 * fields
 *
 * @param http.Request request data
 *
 * @return string errors or empty string
 */
func ReportErrors(r *http.Request) string {
    errors := validatePostFields(r)
    if (len(errors) > 0) {
        jsonString, err := json.Marshal(errors)
        if (err != nil) {
            log.Fatal(err)
        }
        return string(jsonString[:])
    }
    return ""
}

/**
 * Validates each post field to see if they
 * exist individually
 *
 * @param http.Request request data
 *
 * @return []string slice of errors
 */
func validatePostFields(r *http.Request) []string {
    errors := []string{}
    if (r.PostFormValue("image_name") == "") {
        errors = append(errors, "image_name must be defined")
    }

    if (r.PostFormValue("text") == "") {
        errors = append(errors, "text must be defined")
    }

    if (r.PostFormValue("font") == "") {
        errors = append(errors, "font must be defined")
    }

    if (r.PostFormValue("font_size") == "") {
        errors = append(errors, "font_size must be defined")
    }

    if (r.PostFormValue("fill_color") == "") {
        errors = append(errors, "fill_color must be defined")
    }

    if (r.PostFormValue("stroke_color") == "") {
        errors = append(errors, "stroke_color must be defined")
    }

    if (r.PostFormValue("output") == "") {
        errors = append(errors, "output must be defined")
    }
    return errors
}
