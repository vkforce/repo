from flask import Flask, Response
import prometheus_client
from prometheus_client import Counter
import time

app = Flask(__name__)

# Create a metric to track time spent and requests made.
REQUEST_COUNTER = Counter('app_requests_total', 'Total number of requests.')

@app.route('/metrics')
def metrics():
    REQUEST_COUNTER.inc()
    res = []
    for name, data in REQUEST_COUNTER._metrics.items():
        res.append(prometheus_client.exposition.make_line(name, data))
    return Response("\n".join(res), mimetype="text/plain")

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
