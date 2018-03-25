FROM python:3.6-stretch

WORKDIR /app
RUN pip install pipenv
COPY Pipfile Pipfile.lock ./

RUN pipenv install -d --system

COPY . .

CMD ["python", "member.py"]