#!/usr/bin/env python
# -*- coding: utf-8 -*-
# author： angus time:2024/5/27
from datetime import datetime

import pendulum
import json
from airflow.models import DAG
from airflow.operators.bash import BashOperator

log_dir = '/data/logs/x-project2'
local_tz = pendulum.timezone("Asia/Shanghai")
start_date = datetime(2024, 1, 1, tzinfo=local_tz)
bash_command = """ cd /data/services/x-project2 && ./service {runner_file} """

default_args = {
    'owner': 'airflow',
    'email': [""],
    'email_on_failure': True,
    'max_active_runs': 1,
    'params': {
        'log_dir': log_dir,
    }
}

dag_holding_daily_gain_loss = DAG(
    dag_id='dag_holding_daily_gain_loss',
    default_args=default_args,
    start_date=start_date,
    schedule_interval='0 22 * * MON-FRI',
    catchup=False,
    tags=['x-project2', 'holding daily gain loss task']
)

dag_holding_daily_gain_loss_task = BashOperator(
    task_id='dag_holding_daily_gain_loss_task',
    dag=dag_holding_daily_gain_loss,
    bash_command=bash_command.format(runner_file="--job=holding_daily_gain_loss_task")
)