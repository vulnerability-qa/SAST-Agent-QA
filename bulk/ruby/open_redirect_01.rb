# CWE-601: Open Redirect
class SessionsController < ApplicationController
  def create
    # ... authenticate user ...
    redirect_to params[:return_to] || root_path # no host check
  end
end
