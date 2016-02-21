package itcrowdipsum

import (
    "html/template"
    "math/rand"
    "net/http"
)

// these are the phrases we pick from when generating lipsum
var phrases = []string{
    "Hello, IT. Have you tried turning it off and on again?",
    "Uh... okay, well, the button on the side, is it glowing?",
    "Yeah, you need to turn it on... uh, the button turns it on.",
    "Yeah, you do know how a button works don't you? No, not on clothes.",
    "Hello, IT. Have you tried forcing an unexpected reboot?",
    "No, no there you go, no there you go. I just heard it come on.",
    "No, no, that's the music you heard when it come on.",
    "No, that's the music you hear when... I'm sorry are you from the past?",
    "See the driver hooks a function by patching the system call table, so its not safe to unload it unless another thread's about to jump in there and do its stuff, and you don't want to end up in the middle of invalid memory!",
    "Oh really? Then why don't you come down and make me then.",
    "Huh, what you think I'm afraid of you? I'm not afraid of you.",
    "You can come down here any time and I'll be waiting for you! [slams down phone] That told her!",
    "I mean, they have no respect for us up there! No respect whatsoever! We're all just drudgeons to them!",
    "Yes! If there were such a thing as a drudgeon, that is what we'd be to them.",
    "It's like they're pally-wally with us when there's a problem with their printer, but once it's fixed...",
    "They just toss us away like yesterday's jam.",
}

// index is the base html string for... index
var index = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>IT Crowd Ipsum</title>
        <link href="http://itcrowdisum.com/css/styles.css" rel="stylesheet" media="all">
        <link href="http://itcrowdisum.com/css/icons.css" rel="stylesheet" media="all">
        <script src="http://code.jquery.com/jquery.min.js"></script>
    </head>
    <body>
        <div id="wrap">
            <header>
                <h1><span class="icon-roy"></span> <span class="icon-jen"></span> <span class="icon-moss"></span></h1>
                <h2>Placeholder text taken from <em>The IT Crowd</em></h2>
            </header>
            <section>
                {{range .Paragraphs}}<p>{{ . }}</p>{{ end }}
                
                <menu>
                <textarea id="text">
{{range .Paragraphs}}{{ . }}

{{ end }}
</textarea>
                    <button type="button" id="copy">Copy?</button>                    
                    <span id="popup">Now press CMD + C / CTRL + C</span>
                </menu>
            </section>
            <footer>
                <p>
                    Inspired by <a href="http://bluthipsum.com">Bluth Ipsum</a>. Made by <a href="http://kivlor.com">Kivlor</a>
                </p>
            </footer>
        </div>
        <script type="text/javascript">
            jQuery(function($){
                // copy?
                $('#copy').on('click', function(){
                    // select the text
                    $('#text').select();

                    // show the popup
                    $('#popup').fadeIn(200).delay(2000).fadeOut(200);
                });
            });
        </script>
    </body>
</html>
`

// right up main street
func main() {
    http.HandleFunc("/", root)
    http.ListenAndServe(":3000", nil)
}

// root is the handler for requests to "/"
func root(w http.ResponseWriter, r *http.Request) {
    // allocate a new html template
    tmpl, err := template.New("home").Parse(index)
    if err != nil {
        panic("unable to parse index")
    }

    // build the template data
    data := struct {
        Paragraphs []string
    }{
        Paragraphs: GenerateLipsum(5),
    }

    // execute the template data
    tmpl.Execute(w, data)
}

// generateLipsum will create a number of paragraphs using randome phrases
func GenerateLipsum(count int) []string {
    var lipsum []string
    var paragraph string

    // loop the paragraph count
    for i := 0; i < count; i++ {
        paragraph = ""
        // about 6 phrases makes a goo paragrpah
        for j := 0; j < 6; j++ {
            paragraph += phrases[rand.Intn(len(phrases))]
        }

        // append our paragraph to lipsum
        lipsum = append(lipsum, paragraph)
    }

    // return lipsum
    return lipsum
}
