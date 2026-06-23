# CWE-915: Mass Assignment without strong parameters
class UsersController < ApplicationController
  def update
    @user = User.find(params[:id])
    @user.update_attributes(params[:user]) # admin flag settable
  end
end
