Cuba.use		Rack::Session::Cookie
Cuba.plugin		Cuba::Render

Cuba.define do

	# generate lipsum
	def generate_lipsum(count = 5)
		# get the quotes from the lipsum files
		quotes = []
		
		lipsum = File.new('./lipsum.txt', 'r')
		lipsum.each do |quote|
			quotes << quote.strip
		end
		lipsum.close
		
		# create the out string
		out = []
		
		(0..count).each do
			paragraph = ''
		
			(0..6).each do
				paragraph << quotes.sample << ' '
			end
			
			out << paragraph
		end
		
		# done!
		out
	end
	
	# only one place to be, home!
	on get do
		on root do
			res.write render("views/home.erb", lipsum: generate_lipsum())
		end
	end
	
end