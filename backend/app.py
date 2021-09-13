from flask import Flask, render_template, redirect, abort, jsonify, request, make_response
from werkzeug.security import generate_password_hash,check_password_hash
import jwt
from flask_talisman import Talisman
from flask_compress import Compress
from pymongo import MongoClient
from functools import wraps
import os
import datetime
import uuid

myclient = MongoClient('mongodb://localhost:27017/')
mydb = myclient['urlshortener']
mycol = mydb['urls']

app = Flask(__name__, static_folder='../dist/static', template_folder='../dist')
app.config['SECRET_KEY'] = os.environ['SECRET_KEY']
Talisman(app, content_security_policy=None)
Compress(app)

def isauthenticated(request):
    try:
        token = request.cookies.get('JWT')
        data = jwt.decode(token, app.config['SECRET_KEY'], algorithms=["HS256"])
        myquery = {'public_id': data['public_id']}
        current_user = mydb['auth'].find_one(myquery)
        if current_user is not None:
            return True
    except:
        return False

@app.route('/')
@app.route('/index.html')
def index():
    return render_template('index.html')

@app.route('/login')
def login():
    if isauthenticated(request):
        return redirect('/panel')
    return render_template('index.html')

@app.route('/panel')
def panel():
    if not isauthenticated(request):
        return redirect('/login')
    return render_template('index.html')

@app.route('/<id>')
def geturl(id):
    myquery = {'id': id}
    newvalues = {'$inc': {'clicks': 1}}
    mydoc = mycol.find_one_and_update(myquery, newvalues)
    if mydoc is not None:
        redirurl = mydoc['url']
        return redirect(redirurl)
    return abort(404)

# @app.route('/api/shorten', methods=['GET', 'POST'])
# def shorten():
#     response_object = {'status': 'success'}
#     if request.method == 'POST':
#         post_data = request.get_json()
#         if post_data.get('operation') == 'shorten':
#             id = post_data.get('id')
#             url = post_data.get('url')
#             mydict = {'id': id, 'url': url, 'clicks': 0}
#             mycol.insert_one(mydict)
#             response_object['shorturl'] = 'onebounce.me/{}'.format(id)
#     return jsonify(response_object)

@app.route('/api/auth/register', methods=['POST'])
def register():
    data = request.get_json()
    hashed_password = generate_password_hash(data['password'], method='sha256')
    myquery = {'public_id': str(uuid.uuid4()), 'username': data['username'], 'password': hashed_password}
    mydoc = mydb['auth'].insert_one(myquery)
    return jsonify({'message': 'registered successfully'})

@app.route('/api/auth/logout')
def logout():
    response = jsonify({'message': 'success'})
    response.set_cookie('JWT', '', expires=0)
    return response

@app.route('/api/auth/login', methods=['POST'])
def authenticate():
    auth = request.authorization
    if not auth or not auth.username or not auth.password:
       return make_response('could not verify', 401, {'Authentication': 'login required"'})
    myquery = {'username': auth.username}
    mydoc = mydb['auth'].find_one(myquery)
    if mydoc is not None:
        if check_password_hash(mydoc['password'], auth.password):
            token = jwt.encode({'public_id' : mydoc['public_id'], 'exp' : datetime.datetime.utcnow() + datetime.timedelta(minutes=30)}, app.config['SECRET_KEY'], "HS256").decode('utf-8')
            out = jsonify({'message': 'success'})
            expire_date = datetime.datetime.now()
            expire_date = expire_date + datetime.timedelta(minutes=30)
            out.set_cookie('JWT', token, expires=expire_date)
            return out

    return make_response('could not verify',  401, {'Authentication': '"login required"'})

app.run(debug=True, host='10.128.0.3', port=443, ssl_context=('onebounce_me.crt', 'onebounce_me.key'))
