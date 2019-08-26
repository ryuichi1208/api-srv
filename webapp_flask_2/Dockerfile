FROM python:3.7.4-alpine3.9
WORKDIR /home/app
COPY . .
RUN pip install -U pip && pip install -r requirements.txt
EXPOSE 80

ENTRYPOINT ["python", "main.py"]
CMD [""]
