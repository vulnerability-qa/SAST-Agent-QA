# CWE-89: SQL Injection via string interpolation in ActiveRecord
class UsersController < ApplicationController
  def show
    @user = User.where("name = '#{params[:name]}'").first
  end
end
