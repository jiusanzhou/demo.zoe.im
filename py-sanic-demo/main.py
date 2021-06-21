from sanic import Sanic
from sanic import response

app = Sanic(__name__)


@app.route('/')
def handle_request(request):
    return response.empty(
        headers={'Location': 'https://www.baidu.com'},
        status=302
    )

def main():
    app.run(port=8991, debug=True)

if __name__ == '__main__': main()
