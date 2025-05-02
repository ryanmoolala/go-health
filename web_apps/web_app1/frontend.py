from flask import Flask, render_template  # type: ignore

app = Flask(__name__)

@app.route('/')
def hello():
    return render_template('index.html')

@app.route('/health/liveness', methods=['GET'])
def health():
    return {
        "name": "Web App Liveness",
        "method": "GET",
        "status": "Alive",
        "interval": 60,
        "timeout": 5,
        "message": "Liveness check passed"
    }, 200

@app.route('/health/readiness', methods=['GET'])
def readiness():
    return {
        "name": "Web App Readiness",
        "method": "GET",
        "status": "Ready",
        "interval": 60,
        "timeout": 5,
        "message": "Readiness check passed"
    }, 200

