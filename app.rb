Cuba.use		Rack::Session::Cookie
Cuba.plugin		Cuba::Render

Cuba.define do
	#generate lipsum
	def generate_lisum(count = 5)
		quotes = [
			'Hello, IT. Have you tried turning it off and on again? ',
			'Uh... okay, well, the button on the side, is it glowing? ',
			'Yeah, you need to turn it on... uh, the button turns it on. ',
			'Yeah, you do know how a button works don\'t you? No, not on clothes. ',
			'Hello, IT. Have you tried forcing an unexpected reboot? ',
			'No, no there you go, no there you go. I just heard it come on. ',
			'No, no, that\'s the music you heard when it come on. ',
			'No, that\'s the music you hear when... I\'m sorry are you from the past?',
			'See the driver hooks a function by patching the system call table, so its not safe to unload it unless another thread\'s about to jump in there and do its stuff, and you don\'t want to end up in the middle of invalid memory!',
			'Oh really? Then why don\'t you come down and make me then.',
			'Huh, what you think I\'m afraid of you? I\'m not afraid of you.',
			'You can come down here any time and I\'ll be waiting for you! [slams down phone] That told her!',
		]
		
		out = ''
		
		(0..count).each do
			out += '<p>'
		
			(0..6).each do
				out += quotes.sample			
			end
			
			out += '</p>'
		end
		
		out
	end
	
	#home!
	on get do
		on root do
			res.write render("views/home.erb", lipsum: generate_lisum())
		end
	end
end