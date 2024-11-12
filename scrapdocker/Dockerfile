FROM ubuntu:latest as builder


# Install python3 and pip3 and flask
RUN apt-get update && apt-get install -y python3 python3-pip python3-venv
RUN python3 -m venv /opt/venv
ENV PATH="/opt/venv/bin:$PATH"

RUN pip3 install flask requests

# Copy the python file
COPY app.py /app.py

FROM scratch

COPY --from=builder /usr/bin/python3 /usr/bin/python3
COPY --from=builder /usr/bin/pip3 /usr/bin/pip3
COPY --from=builder /usr/local/lib/python3.*/dist-packages /usr/local/lib/python3.*/dist-packages
