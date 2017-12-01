package server

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets0e8e7c3521b97c4280efaeea9a1876bed479a018 = "#!/usr/bin/env python3\n\nimport datetime\n\nprint('HTTP request received at {}'.format(datetime.datetime.now()))\n"
var _Assetsc5cc2dadf26cb5b4a53904786db07d3095062cf7 = "<pre>\n{{ .CommandResult.Stdout }}\n</pre>\n"
var _Assetsc467833fd49a3e903d8b62f3baf59795e9b30a70 = "#!/usr/bin/env python3\n\nimport os\nimport sys\n\nprint('hello, {}. method is {}, url is {}'.format(os.environ['GOMOKU'], os.environ['METHOD'], sys.argv[1]))\n"
var _Assetsdd025a990ab8cb71a1541c932934fee949e86328 = "<!DOCTYPE html>\n<html>\n<head>\n    <title>gomoku</title>\n</head>\n<body>\n    <button id=\"greeting\">click!</button>\n    <script src=\"../js/page.js\"></script>\n</body>\n</html>"
var _Assets006710acf9897239b9f1e219506e50d19b8ebdce = "var button = document.getElementById('greeting');\nbutton.addEventListener('click', function() {\n    alert('hello, gomoku!');\n});"
var _Assets8521fe6f6be9deadc17b1ddb21551a774be89c0c = "#!/usr/bin/env python3\n\nimport datetime\n\nf = open('baz.txt', 'w')\nf.write('HTTP request received at {}'.format(datetime.datetime.now()))\nf.close()\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"static/html": []string{"index.html"}, "static/js": []string{"page.js"}, "static": []string{}}, map[string]*assets.File{
	"baz.py": &assets.File{
		Path:     "baz.py",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511619506, 1511619506000000000),
		Data:     []byte(_Assets8521fe6f6be9deadc17b1ddb21551a774be89c0c),
	}, "static": &assets.File{
		Path:     "static",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1512132932, 1512132932000000000),
		Data:     nil,
	}, "static/html": &assets.File{
		Path:     "static/html",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1512133255, 1512133255000000000),
		Data:     nil,
	}, "static/html/index.html": &assets.File{
		Path:     "static/html/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1512133255, 1512133255000000000),
		Data:     []byte(_Assetsdd025a990ab8cb71a1541c932934fee949e86328),
	}, "static/js": &assets.File{
		Path:     "static/js",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1512133675, 1512133675000000000),
		Data:     nil,
	}, "static/js/page.js": &assets.File{
		Path:     "static/js/page.js",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1512133675, 1512133675000000000),
		Data:     []byte(_Assets006710acf9897239b9f1e219506e50d19b8ebdce),
	}, "bar.py": &assets.File{
		Path:     "bar.py",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511619427, 1511619427000000000),
		Data:     []byte(_Assets0e8e7c3521b97c4280efaeea9a1876bed479a018),
	}, "bar.tmpl": &assets.File{
		Path:     "bar.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511619226, 1511619226000000000),
		Data:     []byte(_Assetsc5cc2dadf26cb5b4a53904786db07d3095062cf7),
	}, "foo.py": &assets.File{
		Path:     "foo.py",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511796965, 1511796965000000000),
		Data:     []byte(_Assetsc467833fd49a3e903d8b62f3baf59795e9b30a70),
	}}, "")
