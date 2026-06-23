# CWE-79: XSS via html_safe
class WelcomeController < ApplicationController
  def index
    @greeting = "Hello #{params[:name]}".html_safe # unescaped user input
  end
end
