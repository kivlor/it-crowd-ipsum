package main

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
	"Yes! Yesterday's jam. That is what we are to them!... Actually, that doesn't work, as a thing, because, you know, jam lasts for ages.",
	"From today, dialing 999 won't get you the Emergency Services, and that's not the only thing that's changing!",
	"Nicer ambulances, faster response times and better looking drivers mean they're not just the Emergency Services, they're your Emergency Services.",
	"So, remember the new number! 0118 999! 88199, 9119 725! ... 3!",
	"Hello? I've had a bit of a tumble.",
	"Well that's easy to remember. 0118 999 88199 9119 725! ... 3!",
	"I don't see how they couldn't just keep it as it was. How hard is it to remember 911?",
	"You mean 999. Yes, yes, I mean 999! Yeah, I know. That's the American one, you berk!",
	"I'll put this over here, with the rest of the fire.",
	"0115... no... 0118... no... 0118 999 ... 3. Hello? Is this the emergency services? Then which country am I speaking to? Hello? Hello?",
	"Dear Sir stroke Madam, I am writing to inform you of a fire which has broken out at the premises of...",
	"Dear Sir stroke Madam. Fire, exclamation mark. Fire, exclamation mark. Help me, exclamation mark. 123 Carrendon Road. Looking forward to hearing from you. All the best, Maurice Moss.",
	"I'm a 32 year old IT-man who works in a basement. Yes, I do the whole Lonely Hearts thing!",
	"Shut up, do what I tell you, I'm not interested; these are just some of the things you'll be hearing if you answer this ad. I'm an idiot and I dont care about anyone but myself. P.S. No dogs!",
	"I'm going to murder you... You bloody woman!",
	"Might want to play a bit hard to get.",
	"We don't need no education. Yes you do. You've just used a double negative.",
	"How can you two... Don't Google the question, Moss!",
	"If anyone was ever rude to me, I used to carry their food around in my trousers. Oh my God! Before you brought it to their table? No, after! Of course, before! Why would I do it after?",
	"While he was eating, did you hear anyone laughing? Like... in the kitchen area? Yes! Yes I did, actually, yes I did. That'd be trouser food!",
	"OK. Moss, what did you have for breakfast this morning? Smarties cereal.",
	"Oh my God. I didn't even know Smarties made a cereal. They don't. It's just Smarties in a bowl with milk.",
	"I am a man, he's a man, we're men! Ok, tell me how your feeling. I feel delicate... and annoyed, and... I think I'm ugly.",
	"I've got Aunt Irma visiting. Oh, do you not like Aunt Irma? I've got an aunt like that.",
	"It's my term for my time of the month. Oh. What time of the month? The weekend?",
	"You know, it's high tide. But we're not on the coast. I'm closed for maintenance! Closed for maintenance? I've fallen to the communists! Well, they do have some strong arguments.",
	"Carrie, Moss! First scene in Carrie! Oh. Okay",
	"A gay musical, called Gay. That's quite gay. Gay musical? Aren't all musicals gay? This must be, like, the gayest musical ever.",
	"A story of a young man trying to find his sexuality in the uncaring Thatcher years. Warning: Contains scenes of graphic homoeroticism.",
	"Graphic homoeroticism? Does that mean they're going to get them out?",
	"You're not comfortable with your sexuality? Oh, I'm very comfortable with my sexuality, I just don't want to be slapped in the face with their sexuality.",
	"He's had quite an evening. Someone stole his wheelchair. Did you see who it was? Red bearded man.",
	"How long have you been disabled? Ten years? Ten years, and how did it happen? If that's not a rude question. Acid?",
	"When I started Reynholm Industries, I had just two things in my possession: a dream and 6 million pounds.",
	"Today I have a business empire the like of which the world has never seen the like of which. I hope it doesn't sound arrogant when I say that I am the greatest man in the world!",
	"Unbelievable! Some idiot disabled his firewall, meaning all the computers on Seven are teeming with viruses, plus I've just had to walk all the way down the motherfudging stairs, because the lifts are broken AGAIN!",
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

// GenerateLipsum will create a number of paragraphs using randome phrases
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
