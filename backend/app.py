from flask import Flask, render_template, redirect, abort, jsonify, request
from flask_talisman import Talisman
from pymongo import MongoClient

myclient = MongoClient('mongodb://localhost:27017/')
mydb = myclient['urlshortener']
mycol = mydb['urls']

app = Flask(__name__, static_folder='../dist/static', template_folder='../dist')
Talisman(app, content_security_policy=None)

@app.route('/')
@app.route('/index.html')
def index():
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

@app.route('/api/shorten', methods=['GET', 'POST'])
def shorten():
    response_object = {'status': 'success'}
    if request.method == 'POST':
        post_data = request.get_json()
        if post_data.get('operation') == 'shorten':
            id = post_data.get('id')
            url = post_data.get('url')
            mydict = {'id': id, 'url': url, 'clicks': 0}
            mycol.insert_one(mydict)
            response_object['shorturl'] = 'onebounce.me/{}'.format(id)
    return jsonify(response_object)

app.run(debug=True, host='10.128.0.3', port=443, ssl_context=('onebounce_me.crt', 'onebounce_me.key'))
