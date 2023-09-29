# Python program to convert the currency
# of one country to that of another country

# Import the modules needed
import requests
import os



from flask import Flask, jsonify, request

class Currency_convertor:
	# empty dict to store the conversion rates
	rates = {}
	def __init__(self, url):
		data = requests.get(url).json()

		# Extracting only the rates from the json data
		self.rates = data["rates"]

	# function to do a simple cross multiplication between
	# the amount and the conversion rates
	def convert(self, from_currency, to_currency, amount):
		initial_amount = amount
		if from_currency != 'EUR' :
			amount = amount / self.rates[from_currency]

		# limiting the precision to 2 decimal places
		amount = round(amount * self.rates[to_currency], 0)
		return int(amount)


# creating a Flask app
app = Flask(__name__)
YOUR_ACCESS_KEY = os.environ['KEY']

# on the terminal type: curl http://127.0.0.1:5000/
# returns hello world when we use GET.
# returns the data that we send when we use POST.
@app.route('/', methods = ['GET'])
def home():
    if(request.method == 'GET'):

        data = "API is healthy"
        return jsonify({'data': data})

@app.route('/convert/<int:num>', methods = ['GET'])
def disp(num):
    url = str.__add__('http://data.fixer.io/api/latest?access_key=', YOUR_ACCESS_KEY)
    c = Currency_convertor(url)
    from_country = 'USD'
    to_country = 'INR'
    amount = num

    converted_amount = c.convert(from_country, to_country, amount)
    
    return jsonify({'Amount_INR': converted_amount})

# Driver code
if __name__ == "__main__":

    app.run(host='0.0.0.0',debug = True)
	
