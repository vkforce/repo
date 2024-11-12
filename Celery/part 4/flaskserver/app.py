from flask import Flask, request
from celery import Celery
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config['CELERY_BROKER_URL'] = 'redis://localhost:6379/0'
app.config['CELERY_RESULT_BACKEND'] = 'redis://localhost:6379/0'
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///messages.db'

celery = Celery(app.name, broker=app.config['CELERY_BROKER_URL'])
celery.conf.update(app.config)

db = SQLAlchemy(app)


class Message(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    content = db.Column(db.String(255))

    def __init__(self, content):
        self.content = content


@celery.task
def schedule_message(message_id):
    with app.app_context():
        message = Message.query.get(message_id)
        # Add your scheduling logic here
        print(f"Scheduled message: {message.content}")


@app.route('/message', methods=['POST'])
def add_message():
    content = request.form.get('content')
    message = Message(content)
    db.session.add(message)
    db.session.commit()
    schedule_message.delay(message.id)
    return 'Message added and scheduled successfully'


if __name__ == '__main__':
    db.create_all()
    app.run()
