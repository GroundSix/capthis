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

package main

import (
	"./capthis"
)

/**
 * Starts the CapThis http server
 * to handle all requests and process
 * all images
 *
 * @return nil
 */
func main() {
	capthis.StartServer()
}
