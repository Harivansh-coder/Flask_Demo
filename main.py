# In same main.py file
# import flask, flask_httpauth and g
from flask import Flask, g
from flask_httpauth import HTTPBasicAuth

# create a flask app
app = Flask(__name__)

# create a flask_httpauth object
auth = HTTPBasicAuth()

# Authentication route for user,admin
@auth.verify_password
def verify_password(username, password):
  if username == 'admin' and password == 'password':
      g.user = username
      return True
  return False

# This is an authecated route login is required by admin user
# Decorator to check if user is logged in
@app.route('/')
@auth.login_required
def index():
  return f'This is an authecated route logged in by user,{g.user}'

# Return current logged in user
@app.route('/current_user')
@auth.login_required
def get_current_user():
  return f"Current logged in user is {g.user}"

if __name__ == '__main__':
  app.run()