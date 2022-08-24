import routers.fish.routes as routes

def register(app):
    '''
    register register all api route
    '''
    app.register_blueprint(routes.fish, url_prefix='/fish')