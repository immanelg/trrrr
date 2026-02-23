Simple translation tool

# Install
## Build from source
Build with Go and install to something like ~/.go/bin/trrr:
```sh
CGO_ENABLED=0 go install -trimpath -ldflags="-s -w" github.com/immanelg/trrr@latest
```

## Download a binary
Install from [Github releases](https://github.com/immanelg/trrr/releases). From there, you can figure out how to unzip a file.

# Usage
```sh
trrr [OPTIONS]...
```
Options:
```
    -s SRC: source language (default: auto)
    -t TARGET: target language
    -b BACKEND: backend (google or lingva)
	-S STRING: use this string instead of stdin
```
Reads text from either stdin or `-S` option.

# Backends 
- Google Translate (unofficial)
- Lingva Translate

# Examples
Translate from English to Spanish. Type text to stdin via cat:
```sh
$ cat | trrr -s en -t es
Things!
<Ctrl-D>
¡Cosas!
```

Auto-detect source language, translate to Hungarian and read text from the second argument:
```sh
$ trrr -t hu 'things!'
A dolgok!
```

Translate text to Russian from X11 primary clipboard (selection) and show the result in a notification:
```sh
xclip -o | trrr -t ru | xargs -0 -I '{}' notify-send -- "trrr" '{}'
```

Same as above, but also copy the result to clipboard:
```sh
xclip -o | trrr -t ru | tee >(xclip -selection clipboard) | xargs -0 -I '{}' notify-send -- "trrr" '{}'
```

Prompt for text with rofi and display the translation with rofi:
```sh
rofi -dmenu -p 'translate' -l 0 | trrr -t ru | xargs -0 -I\{\} rofi -p 'translation' -e \{\}
```

Example for sxhkd config (~/.config/sxhkd/sxhkdrc):
```conf
super + {w,W}
    xclip -o | trrr -t {ru,en} | xargs -0 -I '\{\}' notify-send -- "trrr" '\{\}'

super + alt + {w,W}
    rofi -dmenu -p 'translate' -l 0 | trrr -t {ru,en} | xargs -0 -I\{\} rofi -p 'translation' -e \{\}
```



# License 
0BSD

