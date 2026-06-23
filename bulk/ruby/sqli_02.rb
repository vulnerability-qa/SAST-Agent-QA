# CWE-89: SQL Injection via find_by_sql
class OrdersController < ApplicationController
  def index
    @orders = Order.find_by_sql("SELECT * FROM orders WHERE status = '" + params[:status] + "'")
  end
end
