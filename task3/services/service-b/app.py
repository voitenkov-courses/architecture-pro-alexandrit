from flask import Flask
import random
import time

from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.http.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
from opentelemetry.instrumentation.flask import FlaskInstrumentor
from opentelemetry.instrumentation.requests import RequestsInstrumentor

trace.set_tracer_provider(
   TracerProvider(
       resource=Resource.create({SERVICE_NAME: "calculate"})
   )
)
tracer = trace.get_tracer(__name__)
otlp_exporter = OTLPSpanExporter(endpoint="http://simplest-agent:4318/v1/traces")
trace.get_tracer_provider().add_span_processor(BatchSpanProcessor(otlp_exporter))

app = Flask(__name__)
FlaskInstrumentor().instrument_app(app)
RequestsInstrumentor().instrument()

@app.route("/order")
def order():
    with tracer.start_as_current_span("order"):    
        random_order_number = random.randint(1, 500)
        time.sleep(random_order_number*0.001) #задержка
        return str(random_order_number)