from airflow.sdk import DAG
from airflow.operators.bash import BashOperator
from datetime import datetime

with DAG(
    dag_id="petabencana_dag",
    schedule="@daily",
    start_date=datetime(2024, 1, 1),
    catchup=False,
) as dag:
    task1 = BashOperator(
        task_id="task1",
        bash_command="/opt/airflow/bin/ingest",
    )
