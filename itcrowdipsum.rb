Cuba.use		Rack::Session::Cookie
Cuba.plugin		Cuba::Render

Cuba.define do
	on get do
		on root do
			res.write render("views/layout.erb") { render("views/home.erb") }
		end
	end
end