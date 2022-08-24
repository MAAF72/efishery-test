from flask import Flask

import routers.router as router

app = Flask(__name__)

router.register(app)