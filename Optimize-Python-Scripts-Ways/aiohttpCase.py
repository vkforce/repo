from aiohttp import web

async def home(req):
    return web.Response(text="home page")

app = web.Application()
app.add_routes([web.get('/', home)])

web.run_app(app)