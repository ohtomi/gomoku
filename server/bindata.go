package server

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets0e8e7c3521b97c4280efaeea9a1876bed479a018 = "#!/usr/bin/env python3\n\nimport datetime\n\nprint('HTTP request received at {}'.format(datetime.datetime.now()))\n"
var _Assetsc5cc2dadf26cb5b4a53904786db07d3095062cf7 = "<pre>\n{{ .CommandResult.Stdout }}\n</pre>\n"
var _Assets8521fe6f6be9deadc17b1ddb21551a774be89c0c = "#!/usr/bin/env python3\n\nimport datetime\n\nf = open('baz.txt', 'w')\nf.write('HTTP request received at {}'.format(datetime.datetime.now()))\nf.close()\n"
var _Assetsc467833fd49a3e903d8b62f3baf59795e9b30a70 = "#!/usr/bin/env python3\n\nimport sys\n\nprint('hello, gomoku. url is {}'.format(sys.argv[1]))\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{}, map[string]*assets.File{
	"bar.py": &assets.File{
		Path:     "bar.py",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511619427, 1511619427000000000),
		Data:     []byte(_Assets0e8e7c3521b97c4280efaeea9a1876bed479a018),
	}, "bar.tmpl": &assets.File{
		Path:     "bar.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511619226, 1511619226000000000),
		Data:     []byte(_Assetsc5cc2dadf26cb5b4a53904786db07d3095062cf7),
	}, "baz.py": &assets.File{
		Path:     "baz.py",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511619506, 1511619506000000000),
		Data:     []byte(_Assets8521fe6f6be9deadc17b1ddb21551a774be89c0c),
	}, "foo.py": &assets.File{
		Path:     "foo.py",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1511619226, 1511619226000000000),
		Data:     []byte(_Assetsc467833fd49a3e903d8b62f3baf59795e9b30a70),
	}}, "")
