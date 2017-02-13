# pbas
Config file finder.

## Usage
Get the library via

```
$ go get -u github.com/tkmru/pbas
```

### Finding Config File

Examples:

```
pbas.SetConfigName("pbas.conf")
pbas.AddConfigPath("/etc/pbas/")
pbas.AddConfigPath("conf/")
configPath := pbas.FindConfigPath()
fmt.Print(configPath)
```

## License
MIT License

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

Copyright (c) @tkmru
